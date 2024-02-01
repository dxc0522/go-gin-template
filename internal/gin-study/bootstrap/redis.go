package bootstrap

import (
	"fmt"
	"github.com/go-gin-template/internal/gin-study/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.App.Config.Redis.Host, global.App.Config.Redis.Port),
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
