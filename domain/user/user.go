package user

type User struct {
	Pk        int    `json:"pk"`
	Username  string `json:"username"`
	Fullname  string `json:"full_name"`
	IsPrivate bool   `json:"is_private"`
}
