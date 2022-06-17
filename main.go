package main

import (
	"encoding/json"
	"fmt"
	"instagram-manager/config"
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
	userService := presentation.NewUserService(userRepository)
	instagramService := presentation.NewInstagramService(instagramClient)
	//1.Step Get all following my friends
	saveAllFollowings(instagramService, userService)
	//2.Step Withdrawal of friends of all followers whose number of friends is less than 1000
	//3.Step Send friend requests to recently pulled friends who have more than 10 mutual friends
}

func saveAllFollowings(instagramService *presentation.InstagramService, userService *presentation.UserService) {
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
