package redis

import (
	"common-server/startup"
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	Cli redis.UniversalClient
	ctx = context.Background()
)

func InitRedis() error {
	timeout := time.Duration(startup.GlobalViper.GetInt("redis.timeout")) * time.Second
	Cli = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        startup.GlobalViper.GetStringSlice("redis.addr"),
		PoolSize:     startup.GlobalViper.GetInt("redis.poolSize"),
		MasterName:   startup.GlobalViper.GetString("redis.masterName"),
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	})
	err := Cli.Ping(ctx).Err()
	if err != nil {
		return err
	}
	log.Info("InitRedis success")
	return nil
}
