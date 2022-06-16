package main

import (
	"encoding/json"
	"fmt"
	"instagram-manager/config"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"instagram-manager/infrastructure/client"
	"instagram-manager/infrastructure/persistance"
	"instagram-manager/presentation"
	"net/http"
)

func main() {
	config.Init()
	c := &http.Client{}
	conf := config.NewConfig()
	instagramClient := client.NewInstagramClient(c, conf)
	userRepository := persistance.NewUserRepository(conf.MongoClient)
	instagramService := presentation.NewInstagramService(instagramClient)
	userService := presentation.NewUserService(userRepository)
	followingFriends := instagramService.GetFollowing()
	saveAllFollowings(followingFriends, instagramService, userService)
}

func saveAllFollowings(followingFriends *[]instagram.Follow, instagramService *presentation.InstagramService, userService *presentation.UserService) {
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
