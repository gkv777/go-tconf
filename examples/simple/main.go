package main

import (
	"fmt"
	"log"
	"os"

	gotconf "github.com/gkv777/go-tconf"
)

func main() {
	url := os.Getenv("TC_URL")
	client_id := os.Getenv("TC_CLIENT")
	client_secret := os.Getenv("TC_SECRET")

	cfg := gotconf.Config{
		ServerURL:    url,
		ClientId:     client_id,
		ClientSecret: client_secret,
	}

	tClient := gotconf.NewClient(cfg)
	if err := tClient.Login(); err != nil {
		log.Fatalln(err)
	}
	log.Println(tClient.GetLoginInfo())

	users, err := tClient.GetUsers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
