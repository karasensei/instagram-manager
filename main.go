package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"instagram-manager/application/app/controller"
	"instagram-manager/config"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"instagram-manager/infrastructure/client"
	"instagram-manager/infrastructure/persistance"
	"instagram-manager/presentation"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fc := fiber.Config{
		AppName: "instagram-manager",
	}
	app := fiber.New(fc)
	config.Init()
	c := &http.Client{}
	conf := config.NewConfig()
	instagramClient := client.NewInstagramClient(c, conf)
	userRepository := persistance.NewUserRepository(conf.MongoClient)
	userService := presentation.NewUserService(userRepository)
	instagramService := presentation.NewInstagramService(instagramClient)
	controller.NewInstagramUserController(userService, instagramService, app)
	log.Fatal(app.Listen("8080"))

	//1.Step Get all following my friends
	saveAllFollowings(instagramService, userService)
	//2.Step Withdrawal of friends of all followers whose number of friends is less than 1000
	saveAllWithdrawalOfFriends(instagramService, userService)
	//3.Step Send friend requests to recently pulled friends who have more than 10 mutual friends
	/* TODO: last version

	 */
}

func saveAllFollowings(instagramService *presentation.InstagramService, userService *presentation.UserService) {
	fmt.Println("Save all my following friends started.")
	followingFriends := instagramService.GetFollowing(3154886759)
	for _, it := range *followingFriends {
		saveUser(instagramService, userService, &it, user.UserType_MY)
	}
	fmt.Println("Save all my following friends finished.")
}

func saveAllWithdrawalOfFriends(instagramService *presentation.InstagramService, userService *presentation.UserService) {
	filter := user.Filter{}
	users := userService.GetAllUsers(filter)
	for _, element := range users {
		if element.IsImport == true {
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
		for _, it := range *userFollowingFriends {
			saveUser(instagramService, userService, &it, user.UserType_FOLLOWING)
		}
		element.IsImport = true
		err := userService.Update(&element)
		if err != nil {
			fmt.Println("User not updated. UserName: " + element.UserName)
			continue
		}
		fmt.Println("Save all " + element.UserName + " following friends finished.")
	}
}

func saveUser(instagramService *presentation.InstagramService, userService *presentation.UserService, f *instagram.Follow, userType user.UserType) {
	isExists := userService.IsExistsByInstagramId(f.Pk)
	if isExists {
		fmt.Println("Friend already saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		return
	}
	fmt.Println("Friend will be saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	profileInfo := instagramService.GetProfileInfo(f.Pk)
	u := user.ConvertBson(*f, *profileInfo, userType)
	err := userService.Save(u)
	if err != nil {
		fmt.Println("Friend not saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		return
	}
	fmt.Println("Friend saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
}
