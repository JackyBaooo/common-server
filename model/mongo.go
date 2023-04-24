package model

import (
	"common-server/config"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	MongoCli *mongo.Client
	DB       *mongo.Database
	Timeout  time.Duration
)

func InitMongoDriver() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GlobalConfig.Url))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	log.Info("InitMongoDriver success")
	MongoCli = client
	DB = client.Database(config.GlobalConfig.DB)
	Timeout = time.Duration(config.GlobalConfig.Timeout) * time.Second
	return nil
}
