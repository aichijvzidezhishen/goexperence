package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	redisClient *redis.Client
	lockKey     string
	lockTimeout time.Duration
}

func NewRedisLock(redisClient *redis.Client, lockKey string, lockTimeout time.Duration) *RedisLock {
	return &RedisLock{
		redisClient: redisClient,
		lockKey:     lockKey,
		lockTimeout: lockTimeout,
	}
}

func (rl *RedisLock) AcquireLock() bool {
	ctx := context.TODO()
	timeout := time.After(5 * time.Second) // 设置超时时间为5秒

	for {
		select {
		case <-timeout:
			// 超时退出
			return false
		default:
			// 获取当前时间戳
			currentTime := time.Now().UnixNano() / int64(time.Millisecond)
			// 尝试在Redis实例上设置锁
			lockAcquired, err := rl.redisClient.SetNX(ctx, rl.lockKey, currentTime, rl.lockTimeout).Result()
			if err != nil {
				// 处理错误
				return false
			}
			if lockAcquired {
				return true
			} else {
				// 未能获取锁，等待一段时间后重试
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func (rl *RedisLock) ReleaseLock() bool {
	ctx := context.TODO()
	_, err := rl.redisClient.Del(ctx, rl.lockKey).Result()
	if err != nil {
		// 处理错误
		return false
	}
	return true
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果有密码
		DB:       0,  // 默认数据库
	})

	lock := NewRedisLock(redisClient, "my_lock", 10*time.Second) // 锁的有效时间为10秒

	if lock.AcquireLock() {
		defer lock.ReleaseLock()
		// 在获取到锁之后执行需要保护的代码
		fmt.Println("Lock acquired. Executing protected code...")
		time.Sleep(5 * time.Second) // 模拟执行一些操作
		fmt.Println("Protected code executed.")
	} else {
		fmt.Println("Failed to acquire lock.")
	}
}