package redislock

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test11(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 如果有密码
		DB:       0,  // 默认数据库
	})

	fmt.Println("连接测试", redisClient.Ping(context.TODO()))
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
