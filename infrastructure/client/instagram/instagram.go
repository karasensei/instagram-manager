package instagram

import (
	"encoding/json"
	"instagram-manager/config"
	"instagram-manager/infrastructure/client/instagram/model"
	"io"
	"net/http"
	"strconv"
)

type InstagramClient interface {
	GetFollowers(count int, nextToken string, linkType string) (*model.Friendships, error)
	GetFollowings(count int, nextToken string) (*model.Friendships, error)
}

type instagramClient struct {
	c *config.Config
}

func NewInstagramClient(c *config.Config) *instagramClient {
	return &instagramClient{c: c}
}

func (i *instagramClient) GetFollowers(count int, nextToken string, linkType string) (*model.Friendships, error) {
	url := "https://i.instagram.com/api/v1/friendships/3154886759/followers/?count=" + strconv.Itoa(count)
	if nextToken != "" {
		url = url + "&max_id=" + nextToken
	}
	if linkType != "" {
		url = url + "&search_surface=" + linkType
	}
	req, _ := http.NewRequest("GET", url, nil)
	addHeaders(req, i.c.InstagramHeader)
	resp, _ := i.c.Client.Do(req)
	if resp == nil {
		return nil, nil
	}
	bytes, _ := io.ReadAll(resp.Body)
	friendships := model.Friendships{}
	json.Unmarshal(bytes, &friendships)
	return &friendships, nil
}

func (i *instagramClient) GetFollowings(count int, nextToken string) (*model.Friendships, error) {
	url := "https://i.instagram.com/api/v1/friendships/3154886759/following/?count=" + strconv.Itoa(count)
	if nextToken != "" {
		url = url + "&max_id=" + nextToken
	}
	req, _ := http.NewRequest("GET", url, nil)
	addHeaders(req, i.c.InstagramHeader)
	resp, _ := i.c.Client.Do(req)
	if resp == nil {
		return nil, nil
	}
	bytes, _ := io.ReadAll(resp.Body)
	friendships := model.Friendships{}
	json.Unmarshal(bytes, &friendships)
	return &friendships, nil
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
