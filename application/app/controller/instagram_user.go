package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"instagram-manager/domain/instagram"
	"instagram-manager/domain/user"
	"strconv"
)

type userService interface {
	Save(user bson.M) error
	IsExistsByInstagramId(id int) bool
}

type instagramService interface {
	GetAllFollower(id int, followers user.FollowType) *[]instagram.Follow
	GetAllFollowing(userId int, following string) *[]instagram.Follow
	GetProfileInfo(userId int) *instagram.ProfileInfo
}

type InstagramUserController struct {
	userService      userService
	instagramService instagramService
}

func NewInstagramUserController(userService userService, instagramService instagramService, app *fiber.App) {
	uc := &InstagramUserController{
		userService:      userService,
		instagramService: instagramService,
	}
	app.Post("/instagram-users/:id/follow", func(c *fiber.Ctx) error {
		userId, _ := c.ParamsInt(c.Params("id"))
		followType := c.Params("type")
		return uc.saveAllFollowerByType(userId, followType)
	})
	app.Get("/instagram-users/:id", func(c *fiber.Ctx) error {
		userId, _ := c.ParamsInt(c.Params("id"))
		followType := c.Params("type")
		allMyUsers := uc.getAllMyUsers(userId, followType)
		return c.SendString(fmt.Sprintln(allMyUsers))
	})
}

func (uc *InstagramUserController) saveAllFollowerByType(userId int, followType string) error {
	var follows *[]instagram.Follow
	if followType == "FOLLOWERS" {
		follows = uc.instagramService.GetAllFollower(userId, user.FollowType_FOLLOWERS)
	} else if followType == "FOLLOWING" {
		follows = uc.instagramService.GetAllFollowing(userId, user.FollowType_FOLLOWING)
	}
	return saveAllInstagramUser(*follows, *uc, user.UserType_MY)
}

func (uc *InstagramUserController) getAllMyUsers(userId int, followType string) *[]instagram.Follow {
	if followType == "FOLLOWERS" {
		return uc.instagramService.GetAllFollower(userId, user.FollowType_FOLLOWERS)
	} else if followType == "FOLLOWING" {
		return uc.instagramService.GetAllFollowing(userId, user.FollowType_FOLLOWING)
	}
	return nil
}

func saveAllInstagramUser(followingFriends []instagram.Follow, uc InstagramUserController, userType user.UserType) error {
	for _, f := range followingFriends {
		isExists := uc.userService.IsExistsByInstagramId(f.Pk)
		if isExists {
			fmt.Println("Friend already saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
			return nil
		}
		fmt.Println("Friend will be saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		profileInfo := uc.instagramService.GetProfileInfo(f.Pk)
		u := user.ConvertBson(f, *profileInfo, userType)
		err := uc.userService.Save(u)
		if err != nil {
			fmt.Println("Friend not saving. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
		}
		fmt.Println("Friend saved. Id: " + strconv.Itoa(f.Pk) + ", UserName: " + f.UserName)
	}
	return nil
}
