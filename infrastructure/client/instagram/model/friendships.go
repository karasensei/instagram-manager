package model

import model "instagram-manager/domain/user"

type Friendships struct {
	Users     []model.User `json:"users"`
	PageSize  int          `json:"page_size"`
	NextMaxId string       `json:"next_max_id"`
	Status    string       `json:"status"`
}
