package app

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"net/http"
)

type userService interface {
	Save(user bson.M) error
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
		profileInfo := uc.instagramService.GetProfileInfo(it.UserName)
		u := user.Convert(it, *profileInfo, user.UserType_MY)
		err := uc.userService.Save(u)
		out, _ := json.Marshal(it)
		if err != nil {
			fmt.Println("User not saving. User: " + string(out))
		}
	}
}
