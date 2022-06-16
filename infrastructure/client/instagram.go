package client

import (
	"encoding/json"
	"instagram-manager/config"
	"instagram-manager/domain/instagram"
	"io"
	"net/http"
	"strconv"
)

type InstagramClient struct {
	c    *http.Client
	conf *config.Config
}

func NewInstagramClient(c *http.Client, conf *config.Config) *InstagramClient {
	return &InstagramClient{
		c:    c,
		conf: conf,
	}
}

func (i *InstagramClient) GetFollowers(count int, nextToken string, linkType string) (*instagram.Friendships, error) {
	url := "https://i.instagram.com/api/v1/friendships/3154886759/followers/?count=" + strconv.Itoa(count)
	if nextToken != "" {
		url = url + "&max_id=" + nextToken
	}
	if linkType != "" {
		url = url + "&search_surface=" + linkType
	}
	req, _ := http.NewRequest("GET", url, nil)
	addHeaders(req, i.conf.InstagramHeader)
	resp, _ := i.c.Do(req)
	if resp == nil {
		return nil, nil
	}
	bytes, _ := io.ReadAll(resp.Body)
	friendships := instagram.Friendships{}
	json.Unmarshal(bytes, &friendships)
	return &friendships, nil
}

func (i *InstagramClient) GetFollowings(count int, nextToken string) (*instagram.Friendships, error) {
	url := "https://i.instagram.com/api/v1/friendships/3154886759/following/?count=" + strconv.Itoa(count)
	if nextToken != "" {
		url = url + "&max_id=" + nextToken
	}
	req, _ := http.NewRequest("GET", url, nil)
	addHeaders(req, i.conf.InstagramHeader)
	resp, _ := i.c.Do(req)
	if resp == nil {
		return nil, nil
	}
	bytes, _ := io.ReadAll(resp.Body)
	friendships := instagram.Friendships{}
	json.Unmarshal(bytes, &friendships)
	return &friendships, nil
}

func (i *InstagramClient) GetProfileInfo(userName string) (*instagram.ProfileInfo, error) {
	url := "https://i.instagram.com/api/v1/users/web_profile_info/?username=" + userName
	req, _ := http.NewRequest("GET", url, nil)
	addHeaders(req, i.conf.InstagramHeader)
	resp, _ := i.c.Do(req)
	if resp == nil {
		return nil, nil
	}
	bytes, _ := io.ReadAll(resp.Body)
	profileInfo := instagram.ProfileInfo{}
	json.Unmarshal(bytes, &profileInfo)
	return &profileInfo, nil
}

func addHeaders(req *http.Request, header *config.InstagramHeader) {
	req.Header.Add("authority", header.Authority)
	req.Header.Add("accept", header.Accept)
	req.Header.Add("accept-language", header.AcceptLanguage)
	req.Header.Add("cookie", header.Cookie)
	req.Header.Add("origin", header.Origin)
	req.Header.Add("referer", header.Referer)
	req.Header.Add("sec-ch-ua", header.SecChUa)
	req.Header.Add("sec-ch-ua-mobile", header.SecChUaMobile)
	req.Header.Add("sec-ch-ua-platform", header.SecChUaPlatform)
	req.Header.Add("sec-fetch-dest", header.SecFetchDest)
	req.Header.Add("sec-fetch-mode", header.SecFetchMode)
	req.Header.Add("sec-fetch-site", header.SecFetchSite)
	req.Header.Add("user-agent", header.UserAgent)
	req.Header.Add("x-asbd-id", header.XAsbdId)
	req.Header.Add("x-csrftoken", header.XCsrftoken)
	req.Header.Add("x-ig-app-id", header.XIgAppId)
	req.Header.Add("x-ig-www-claim", header.XIgWwwClaim)
}
