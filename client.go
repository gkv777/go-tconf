package gotconf

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type gotconf struct {
	mx      sync.RWMutex
	client  *resty.Client
	cfg     Config
	isLogin bool
	token   string
	exp_in  time.Time
}

type Config struct {
	ServerURL    string
	ClientId     string
	ClientSecret string
}

func NewClient(cfg Config) GoTConf {
	tlsCfg := &tls.Config{InsecureSkipVerify: true}
	client := resty.New()
	client.SetTLSClientConfig(tlsCfg)

	return &gotconf{
		cfg:     cfg,
		mx:      sync.RWMutex{},
		client:  client,
		isLogin: false,
	}
}

func (c *gotconf) Login() error {
	body := fmt.Sprintf(`{
		"grant_type": "client_credentials",
    	"client_id": "%s",
    	"client_secret": "%s"
	}`, c.cfg.ClientId, c.cfg.ClientSecret)

	var auth AuthSuccess
	_, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(body)).
		SetResult(&auth).
		Post(fmt.Sprintf("%s/oauth2/v1/token", c.cfg.ServerURL))

	if err != nil {
		return err
	}

	c.mx.Lock()
	defer c.mx.Unlock()

	c.token = auth.AccessToken
	c.exp_in = time.Now().Add(time.Second * time.Duration(auth.ExpiresIn))
	
	c.isLogin = true

	return nil
}

func (c *gotconf) IsLogin() bool {
	c.mx.RLock()
	defer c.mx.RUnlock()

	now := time.Now()
	if c.isLogin && now.Before(c.exp_in) {
		return true
	}
	return false
}

func (c *gotconf) GetLoginInfo() string {
	if !c.IsLogin() {
		return "the client is not connected to the TrueConf server"
	}
	return fmt.Sprintf("the client is connected to TrueConf server, the token will expire in %s", c.exp_in.Format("2006-01-02 15:04:05"))
}

func (c *gotconf) GetUsers() ([]*User, error) {
	var res usersListResp

	_, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&res).
		SetQueryParam("access_token", c.token).
		Get(fmt.Sprintf("%s/api/%s/users", c.cfg.ServerURL, API_VER))

	
	if err != nil {
		return nil, err
	}

	return res.Users, nil

}
