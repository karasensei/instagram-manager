package model

type Friendships struct {
	Users     []Follow `json:"users"`
	PageSize  int      `json:"page_size"`
	NextMaxId string   `json:"next_max_id"`
	Status    string   `json:"status"`
}
