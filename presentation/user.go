package presentation

import (
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository interface {
	Save(user bson.M) error
	ExistsUserById(id int) bool
}

type UserService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Save(user bson.M) error {
	return u.userRepository.Save(user)
}

func (u *UserService) IsExistsById(id int) bool {
	return u.userRepository.ExistsUserById(id)
}
