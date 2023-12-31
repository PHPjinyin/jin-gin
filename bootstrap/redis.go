package bootstrap

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"jin-gin/global"
)

func InitializeRedis() *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return Client
}
