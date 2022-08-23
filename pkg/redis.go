package pkg

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var RedisCtx = context.Background()
var Redis *redis.Client

func RedisSetup() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     RedisConfig.Addr,
		Password: RedisConfig.Password,
		DB:       RedisConfig.Db,
	})

	pong, err := Redis.Ping(RedisCtx).Result()

	logrus.WithFields(logrus.Fields{
		"pong": pong,
		"err":  err,
	}).Info("redis 连接情况")

}
