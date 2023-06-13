package mongo

import (
	"common-server/startup"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	DB      *mongo.Database
	Timeout time.Duration
)

func InitMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(startup.GlobalConfig.Url))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	log.Info("InitMongo success")
	DB = client.Database(startup.GlobalConfig.DB)
	Timeout = time.Duration(startup.GlobalConfig.Timeout) * time.Second
	return nil
}
