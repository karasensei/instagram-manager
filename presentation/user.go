package presentation

import (
	"go.mongodb.org/mongo-driver/bson"
	"instagram-manager/domain/user"
)

type userRepository interface {
	Save(user bson.M) error
	ExistsUserById(id int) bool
	GetAllUsers(f user.Filter) []user.User
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

func (u *UserService) GetAllUsers(f user.Filter) []user.User {
	return u.userRepository.GetAllUsers(f)
}
