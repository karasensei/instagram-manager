package main

import (
	"encoding/json"
	"fmt"
	"instagram-manager/config"
	"instagram-manager/domain/user"
	"instagram-manager/infrastructure/client"
	"instagram-manager/infrastructure/repository"
	"instagram-manager/infrastructure/service"
	"net/http"
)

func main() {
	config.Init()
	c := &http.Client{}
	conf := config.NewConfig()
	instagramClient := client.NewInstagramClient(c, conf)
	userRepository := repository.NewUserRepository(conf.MongoClient)
	instagramService := service.NewInstagramService(instagramClient)
	userService := service.NewUserService(userRepository)
	followingFriends := instagramService.GetFollowing()
	for _, it := range *followingFriends {
		profileInfo := instagramService.GetProfileInfo(it.UserName)
		u := user.Convert(it, *profileInfo, user.UserType_MY)
		err := userService.Save(u)
		out, _ := json.Marshal(it)
		if err != nil {
			fmt.Println("User not saving. User: " + string(out))
		}
	}
}
