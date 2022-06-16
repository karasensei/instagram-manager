package service

import (
	"instagram-manager/domain/model"
	"log"
)

type instagramClient interface {
	GetFollowers(count int, nextToken string, linkType string) (*model.Friendships, error)
	GetFollowings(count int, nextToken string) (*model.Friendships, error)
	GetProfileInfo(userName string) (*model.ProfileInfo, error)
}

type InstagramService struct {
	instagramClient instagramClient
}

func NewInstagramService(instagramClient instagramClient) *InstagramService {
	return &InstagramService{
		instagramClient: instagramClient,
	}
}

func (f *InstagramService) GetFollowers() *[]model.Follow {
	var fw []model.Follow
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
		fw = append(fw, friendships.Users...)
		if friendships.NextMaxId == "" {
			break
		}
		nextToken = friendships.NextMaxId
	}
	return &fw
}

func (f *InstagramService) GetFollowing() *[]model.Follow {
	var fw []model.Follow
	nextToken := ""
	for {
		friendships, err := f.instagramClient.GetFollowings(12, nextToken)
		if err != nil {
			log.Fatal(err)
		}
		if friendships == nil {
			break
		}
		fw = append(fw, friendships.Users...)
		if friendships.NextMaxId == "" {
			break
		}
		nextToken = friendships.NextMaxId
	}
	return &fw
}

func (f *InstagramService) GetProfileInfo(userName string) *model.ProfileInfo {
	profileInfo, _ := f.instagramClient.GetProfileInfo(userName)
	return profileInfo
}
