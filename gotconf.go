package gotconf

const API_VER = "v3.3"

type GoTConf interface {
	Login() error
	IsLogin() bool
	GetLoginInfo() string

	GetUsers()([]*User, error)
}
