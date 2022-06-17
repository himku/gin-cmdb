package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/17 14:12
 **/

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(initConfig.Redis.Hostname+":%d", initConfig.Redis.Port),
		Password: initConfig.Redis.Password,
		DB:       initConfig.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil
	}
	return client
}
