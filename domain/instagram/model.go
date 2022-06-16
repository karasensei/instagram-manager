package instagram

type Friendships struct {
	Users     []Follow `json:"users"`
	PageSize  int      `json:"page_size"`
	NextMaxId string   `json:"next_max_id"`
	Status    string   `json:"status"`
}

type Follow struct {
	Pk        int    `json:"pk"`
	UserName  string `json:"username"`
	FullName  string `json:"full_name"`
	IsPrivate bool   `json:"is_private"`
}

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
