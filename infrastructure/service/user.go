package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"instagram-manager/infrastructure/repository"
	"time"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Save(user bson.M) error {
	collection := u.userRepository.MongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
