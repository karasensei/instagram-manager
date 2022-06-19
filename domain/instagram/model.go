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
	User ProfileInfoUser `json:"user"`
}

type ProfileInfoUser struct {
	FollowingCount       int            `json:"following_count"`
	FollowersCount       int            `json:"follower_count"`
	ProfilePicUrlHd      ProfilePicInfo `json:"hd_profile_pic_url_info"`
	MutualFollowersCount int            `json:"mutual_followers_count"`
}

type ProfilePicInfo struct {
	Url string `json:"url"`
}
