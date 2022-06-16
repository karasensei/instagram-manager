package model

type Follow struct {
	Pk        int    `json:"pk"`
	UserName  string `json:"username"`
	FullName  string `json:"full_name"`
	IsPrivate bool   `json:"is_private"`
}
