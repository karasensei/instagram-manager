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
	"sync"
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
	wg := new(sync.WaitGroup)
	for _, it := range *followingFriends {
		wg.Add(len(*followingFriends))
		go saveUser(instagramService, userService, &it, user.UserType_MY, wg)
	}
	wg.Wait()
	fmt.Println("Save all my following friends finished.")
}

func saveAllWithdrawalOfFriends(instagramService *presentation.InstagramService, userService *presentation.UserService) {
	filter := user.UserFilter{}
	users := userService.GetAllUsers(filter)
	for _, element := range users {
		if element.Import == true {
			fmt.Println("Already imported. UserName: " + element.UserName)
			continue
		}
		if element.FollowingCount > 1500 {
			fmt.Println("User following count bigger than 1500. UserName: " + element.UserName)
			continue
		}
		fmt.Println("Save all " + element.UserName + " following friends started.")
		userFollowingFriends := instagramService.GetFollowing(element.InstagramId)
		if *userFollowingFriends == nil {
			fmt.Println("userFollowingFriends is nil. UserName: " + element.UserName)
			continue
		}
		wg := new(sync.WaitGroup)
		for _, it := range *userFollowingFriends {
			wg.Add(len(*userFollowingFriends))
			go saveUser(instagramService, userService, &it, user.UserType_FOLLOWING, wg)
		}
		wg.Wait()
		fmt.Println("Save all " + element.UserName + " following friends finished.")
	}
}

func saveUser(instagramService *presentation.InstagramService, userService *presentation.UserService, f *instagram.Follow, userType user.UserType, wg *sync.WaitGroup) {
	isExists := userService.IsExistsById(f.Pk)
	if isExists {
		fmt.Println("Friend already saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		wg.Done()
		return
	}
	fmt.Println("Friend will be saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	profileInfo := instagramService.GetProfileInfo(f.Pk)
	u := user.Convert(*f, *profileInfo, userType)
	err := userService.Save(u)
	if err != nil {
		fmt.Println("Friend not saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		wg.Done()
		return
	}
	wg.Done()
	fmt.Println("Friend saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
}
