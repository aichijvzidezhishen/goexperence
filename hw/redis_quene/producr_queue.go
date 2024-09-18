package redisquene

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // Redis服务器密码
		DB:       0,                // Redis数据库索引
	})
}

func producrMessages(msg string) error {
	for i := 0; i < 3; i++ {
		err := RedisClient.RPush(context.TODO(), "msgQueue", msg).Err()
		if err == nil {
			return nil
		}
		log.Printf("Failed to push message to queue: %v", err)
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("failed to push message to queue after multiple attempts")
}
