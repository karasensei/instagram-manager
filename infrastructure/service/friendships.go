package service

import (
	"instagram-manager/domain/model"
	"instagram-manager/domain/user"
	"log"
)

type instagramClient interface {
	GetFollowers(count int, nextToken string, linkType string) (*model.Friendships, error)
	GetFollowings(count int, nextToken string) (*model.Friendships, error)
}

type FriendshipsService struct {
	instagramClient instagramClient
}

func NewFriendshipsService(instagramClient instagramClient) *FriendshipsService {
	return &FriendshipsService{
		instagramClient: instagramClient,
	}
}

func (f *FriendshipsService) GetFollowers() *[]user.User {
	var users []user.User
	nextToken := ""
	linkType := "follow_list_page"
	for {
		friendships, err := f.instagramClient.GetFollowers(12, nextToken, linkType)
		if err != nil {
			log.Fatal(err)
		}
		if friendships == nil {
			break
		}
		users = append(users, friendships.Users...)
		if friendships.NextMaxId == "" {
			break
		}
		nextToken = friendships.NextMaxId
	}
	return &users
}
