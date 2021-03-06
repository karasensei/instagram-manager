package persistance

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"instagram-manager/domain/user"
	"log"
	"time"
)

type UserRepository struct {
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client) *UserRepository {
	return &UserRepository{mongoClient: mongoClient}
}
func (ur *UserRepository) Save(user bson.M) error {
	collection := ur.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ExistsByInstagramId(instagramId int) bool {
	collection := ur.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.M{"InstagramId": instagramId})
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return false
	}
	return true
}

func (ur *UserRepository) GetAllUsers(f user.Filter) []user.User {
	collection := ur.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	var users []user.User
	for cur.Next(context.TODO()) {
		var user user.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		users = append(users, user)
	}
	return users
}

func (ur *UserRepository) Update(u *user.User) error {
	collection := ur.mongoClient.Database("instagramManager").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()
	_, err := collection.UpdateByID(ctx, u.ID, user.GetBson(*u))
	return err
}
