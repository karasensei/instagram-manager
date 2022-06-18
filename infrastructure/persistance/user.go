package persistance

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) *UserRepository {
	return &UserRepository{mongoClient: mongoClient}
}
func (u *UserRepository) Save(user bson.M) error {
	collection := u.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) ExistsUserById(id int) bool {
	collection := u.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.M{"InstagramId": id})
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return false
	}
	return true
}
