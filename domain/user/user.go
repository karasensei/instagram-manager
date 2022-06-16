package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"instagram-manager/domain/model"
)

type UserType string

const (
	UserType_MY        UserType = "MY"
	UserType_FOLLOWERS          = "FOLLOWERS"
	UserType_FOLLOWING          = "FOLLOWING"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	InstagramId     int                `bson:"InstagramId"`
	ProfilePicUrlHd string             `bson:"ProfilePicUrlHd"`
	UserName        string             `bson:"UserName"`
	FullName        string             `bson:"FullName"`
	IsPrivate       bool               `bson:"IsPrivate"`
	Followers       int                `bson:"Followers"`
	Following       int                `bson:"Following"`
	Type            UserType           `bson:"Type"`
}

func Convert(f model.Follow, p model.ProfileInfo, userType UserType) bson.M {
	return bson.M{
		"InstagramId":     f.Pk,
		"ProfilePicUrlHd": p.Data.User.ProfilePicUrlHd,
		"UserName":        f.UserName,
		"FullName":        f.FullName,
		"IsPrivate":       f.IsPrivate,
		"Followers":       p.Data.User.Followers.Count,
		"Following":       p.Data.User.Following.Count,
		"Type":            userType,
	}
}
