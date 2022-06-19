package main

import (
	"fmt"
	"instagram-manager/config"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"instagram-manager/infrastructure/client"
	"instagram-manager/infrastructure/persistance"
	"instagram-manager/presentation"
	"net/http"
	"strconv"
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
	saveAllWithdrawalOfFriends(instagramService, userService)
	//3.Step Send friend requests to recently pulled friends who have more than 10 mutual friends

	/* TODO: last version
	app.NewUserController(userService, instagramService)
	http.ListenAndServe(":8090", nil)
	*/
}

func saveAllFollowings(instagramService *presentation.InstagramService, userService *presentation.UserService) {
	fmt.Println("Save all my following friends started.")
	followingFriends := instagramService.GetFollowing(3154886759)
	for _, it := range *followingFriends {
		saveUser(instagramService, userService, it, user.UserType_MY)
	}
	fmt.Println("Save all my following friends finished.")
}

func saveAllWithdrawalOfFriends(instagramService *presentation.InstagramService, userService *presentation.UserService) {
	filter := user.UserFilter{}
	users := userService.GetAllUsers(filter)
	for _, element := range users {
		fmt.Println("Save all " + element.UserName + " following friends started.")
		if element.FollowingCount > 1000 {
			continue
		}
		userFollowingFriends := instagramService.GetFollowing(element.InstagramId)
		if userFollowingFriends == nil {
			fmt.Println("userFollowingFriends is nil. UserId: " + strconv.Itoa(element.InstagramId))
		}
		for _, it := range *userFollowingFriends {
			saveUser(instagramService, userService, it, user.UserType_FOLLOWING)
		}
		fmt.Println("Save all " + element.UserName + " following friends finished.")
	}
}

func saveUser(instagramService *presentation.InstagramService, userService *presentation.UserService, f instagram.Follow, userType user.UserType) {
	isExists := userService.IsExistsById(f.Pk)
	if isExists {
		fmt.Println("Friend already saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		return
	}
	fmt.Println("Friend will be saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	profileInfo := instagramService.GetProfileInfo(f.Pk)
	u := user.Convert(f, *profileInfo, userType)
	err := userService.Save(u)
	if err != nil {
		fmt.Println("Friend not saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	}
	fmt.Println("Friend saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
}
