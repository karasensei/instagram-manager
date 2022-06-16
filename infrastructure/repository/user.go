package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	MongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) *UserRepository {
	return &UserRepository{MongoClient: mongoClient}
}
