package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Config struct {
	InstagramHeader InstagramHeader
	MongoClient     *mongo.Client
}

func NewConfig() *Config {
	return &Config{
		InstagramHeader: fillInstagramHeader(),
		MongoClient:     initialMongoDB(),
	}
}

func Init() {
	active := os.Getenv("ACTIVE")
	if active == "" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName("config-" + active)
	}
	viper.SetConfigType("yml")
	viper.AddConfigPath("./resource")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

type InstagramHeader struct {
	Authority       string
	Accept          string
	AcceptLanguage  string
	Cookie          string
	Origin          string
	Referer         string
	SecChUa         string
	SecChUaMobile   string
	SecChUaPlatform string
	SecFetchDest    string
	SecFetchMode    string
	SecFetchSite    string
	UserAgent       string
	XAsbdId         string
	XCsrftoken      string
	XIgAppId        string
	XIgWwwClaim     string
}

func fillInstagramHeader() InstagramHeader {
	return InstagramHeader{
		viper.GetString("instagram.header.authority"),
		viper.GetString("instagram.header.accept"),
		viper.GetString("instagram.header.accept-language"),
		viper.GetString("instagram.header.cookie"),
		viper.GetString("instagram.header.origin"),
		viper.GetString("instagram.header.referer"),
		viper.GetString("instagram.header.sec-ch-ua"),
		viper.GetString("instagram.header.sec-ch-ua-mobile"),
		viper.GetString("instagram.header.sec-ch-ua-platform"),
		viper.GetString("instagram.header.sec-fetch-dest"),
		viper.GetString("instagram.header.sec-fetch-mode"),
		viper.GetString("instagram.header.sec-fetch-site"),
		viper.GetString("instagram.header.user-agent"),
		viper.GetString("instagram.header.x-asbd-id"),
		viper.GetString("instagram.header.x-csrftoken"),
		viper.GetString("instagram.header.x-ig-app-id"),
		viper.GetString("instagram.header.x-ig-www-claim"),
	}
}

func initialMongoDB() *mongo.Client {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.url")))
	if err != nil {
		panic(fmt.Errorf("Fatal error not new mongo client created: %w \n", err))
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("Fatal error mongo client not connected: %w \n", err))
	}
	return mongoClient
}
