package app

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"net/http"
	"strconv"
)

type userService interface {
	Save(user bson.M) error
	IsExistsById(id int) bool
}

type instagramService interface {
	GetFollowing() *[]instagram.Follow
	GetProfileInfo(userName string) *instagram.ProfileInfo
}

type UserController struct {
	userService      userService
	instagramService instagramService
}

func NewUserController(userService userService, instagramService instagramService) *UserController {
	uc := &UserController{
		userService:      userService,
		instagramService: instagramService,
	}
	http.HandleFunc("/", uc.saveAllFollowings)
	return uc
}

func (uc *UserController) saveAllFollowings(w http.ResponseWriter, req *http.Request) {
	followingFriends := uc.instagramService.GetFollowing()
	for _, it := range *followingFriends {
		isExists := uc.userService.IsExistsById(it.Pk)
		if isExists {
			fmt.Println("Friend already saved. Id: " + strconv.Itoa(it.Pk) + ", UserName: " + it.UserName)
			continue
		}
		fmt.Println("Friend will be saving. Id: " + strconv.Itoa(it.Pk) + ", UserName: " + it.UserName)
		profileInfo := uc.instagramService.GetProfileInfo(it.UserName)
		u := user.Convert(it, *profileInfo, user.UserType_MY)
		err := uc.userService.Save(u)
		if err != nil {
			fmt.Println("Friend not saving. Id: " + strconv.Itoa(it.Pk) + ", UserName: " + it.UserName)
		}
		fmt.Println("Friend saved. Id: " + strconv.Itoa(it.Pk) + ", UserName: " + it.UserName)
	}
}
