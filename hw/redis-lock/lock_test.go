package redislock

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test11(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			fmt.Println(fmt.Errorf("print stack %v", err).Error(), stack)
		}
	}()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 如果有密码
		DB:       0,  // 默认数据库
	})
	err := redisClient.Ping(context.TODO()).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接测试", redisClient.Ping(context.TODO()))
	lock := NewRedisLock(redisClient, "my_lock", 10*time.Second) // 锁的有效时间为10秒

	if lock.AcquireLock() {
		// defer lock.ReleaseLock()
		// 在获取到锁之后执行需要保护的代码
		fmt.Println("Lock acquired. Executing protected code...")
		time.Sleep(5 * time.Second) // 模拟执行一些操作
		fmt.Println("Protected code executed.")
	} else {
		fmt.Println("Failed to acquire lock.")
	}
}
