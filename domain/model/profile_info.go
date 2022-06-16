package model

type ProfileInfo struct {
	Data ProfileInfoData `json:"data"`
}

type ProfileInfoData struct {
	User ProfileInfoUser `json:"user"`
}

type ProfileInfoUser struct {
	Following       CountInfo `json:"edge_follow"`
	Followers       CountInfo `json:"edge_followed_by"`
	ProfilePicUrlHd string    `json:"profile_pic_url_hd"`
}

type CountInfo struct {
	Count int `json:"count"`
}
