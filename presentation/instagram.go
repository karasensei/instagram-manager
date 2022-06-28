package presentation

import (
	"instagram-manager/domain/instagram"
	"log"
)

type instagramClient interface {
	GetFollowers(count int, nextToken string, linkType string) (*instagram.Friendships, error)
	GetFollowings(count int, nextToken string, userId int) (*instagram.Friendships, error)
	GetProfileInfo(userId int) (*instagram.ProfileInfo, error)
}

type InstagramService struct {
	instagramClient instagramClient
}

func NewInstagramService(instagramClient instagramClient) *InstagramService {
	return &InstagramService{
		instagramClient: instagramClient,
	}
}

func (f *InstagramService) GetAllFollower() *[]instagram.Follow {
	var fw []instagram.Follow
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

func (f *InstagramService) GetAllFollowing(userId int) *[]instagram.Follow {
	var fw []instagram.Follow
	nextToken := ""
	for {
		friendships, err := f.instagramClient.GetFollowings(12, nextToken, userId)
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

func (f *InstagramService) GetProfileInfo(userId int) *instagram.ProfileInfo {
	profileInfo, _ := f.instagramClient.GetProfileInfo(userId)
	return profileInfo
}
