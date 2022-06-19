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
	GetFollowing(userId int) *[]instagram.Follow
	GetProfileInfo(userId int) *instagram.ProfileInfo
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
	followingFriends := uc.instagramService.GetFollowing(3154886759)
	for _, f := range *followingFriends {
		isExists := uc.userService.IsExistsById(f.Pk)
		if isExists {
			fmt.Println("Friend already saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
			return
		}
		fmt.Println("Friend will be saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		profileInfo := uc.instagramService.GetProfileInfo(f.Pk)
		u := user.Convert(f, *profileInfo, user.UserType_MY)
		err := uc.userService.Save(u)
		if err != nil {
			fmt.Println("Friend not saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		}
		fmt.Println("Friend saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	}
}
