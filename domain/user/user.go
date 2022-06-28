package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"instagram-manager/domain/instagram"
)

type UserType string
type FollowType string

const (
	UserType_MY        UserType = "MY"
	UserType_FOLLOWERS          = "FOLLOWERS"
	UserType_FOLLOWING          = "FOLLOWING"
)

const (
	FollowType_FOLLOWERS FollowType = "FOLLOWERS"
	FollowType_FOLLOWING            = "FOLLOWING"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id"`
	InstagramId          int                `bson:"InstagramId"`
	ProfilePicUrlHd      string             `bson:"ProfilePicUrlHd"`
	UserName             string             `bson:"UserName"`
	FullName             string             `bson:"FullName"`
	IsPrivate            bool               `bson:"IsPrivate"`
	FollowersCount       int                `bson:"FollowersCount"`
	FollowingCount       int                `bson:"FollowingCount"`
	Type                 UserType           `bson:"Type"`
	MutualFollowersCount int                `bson:"MutualFollowersCount"`
	IsImport             bool               `bson:"IsImport"`
}

func ConvertBson(f instagram.Follow, p instagram.ProfileInfo, userType UserType) bson.M {
	return bson.M{
		"InstagramId":          f.Pk,
		"ProfilePicUrlHd":      p.User.ProfilePicUrlHd.Url,
		"UserName":             f.UserName,
		"FullName":             f.FullName,
		"IsPrivate":            f.IsPrivate,
		"Followers":            p.User.FollowersCount,
		"Following":            p.User.FollowingCount,
		"Type":                 userType,
		"MutualFollowersCount": p.User.MutualFollowersCount,
		"IsImport":             false,
	}
}

func GetBson(u User) bson.M {
	return bson.M{
		"InstagramId":          u.InstagramId,
		"ProfilePicUrlHd":      u.ProfilePicUrlHd,
		"UserName":             u.UserName,
		"FullName":             u.FullName,
		"IsPrivate":            u.IsPrivate,
		"FollowersCount":       u.FollowersCount,
		"FollowingCount":       u.FollowingCount,
		"Type":                 u.Type,
		"MutualFollowersCount": u.MutualFollowersCount,
		"IsImport":             u.IsImport,
	}
}
