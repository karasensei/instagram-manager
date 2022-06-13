package main

import (
	"instagram-manager/config"
	"instagram-manager/infrastructure/client/instagram"
	"net/http"
)

func main() {
	config.Init()
	c := http.Client{}
	conf := config.NewConfig(&c)
	instagram.NewInstagramClient(conf)
}
