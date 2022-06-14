package main

import (
	"instagram-manager/config"
	"instagram-manager/infrastructure/client"
	"instagram-manager/infrastructure/service"
	"net/http"
)

func main() {
	config.Init()
	c := &http.Client{}
	conf := config.NewConfig()
	instagramClient := client.NewInstagramClient(c, conf)
	service.NewFriendshipsService(instagramClient)
}
