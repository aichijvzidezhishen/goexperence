package test

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestXxx(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	r, err := NewRedisLimiter(RedisLimiterConfig{
		Client: client,
		Key:    "test-rate",
		Burst:  1024,
		QPS:    64,
	})
	if err != nil {
		log.Fatalf("failed to create rate limiter: %v", err)
	}

	for i := 0; i < 5; i++ {
		err := r.Wait(context.Background())
		log.Printf("worker %d allowed: %v", i, err)
	}
}

func BenchmarkT(b *testing.B) {

	// rand.Seed(time.Now().Unix())
	b.ResetTimer()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r, err := NewRedisLimiter(RedisLimiterConfig{
		Client: client,
		Key:    "test-rate",
		Burst:  100000,
		QPS:    10000,
	})
	if err != nil {
		log.Fatalf("failed to create rate limiter: %v", err)
	}
	var wg sync.WaitGroup
	wg.Add(100)
	for g := 0; g < 100; g++ {
		go func() {
			defer wg.Done()

			for i := 0; i < b.N; i++ {
				err = r.Wait(context.Background())
			}
		}()
	}

	wg.Wait()

}

// goroutines ：10
// limit : 100000
// qps:10000

// goos: windows
// goarch: amd64
// pkg: goexperence/test
// cpu: 11th Gen Intel(R) Core(TM) i7-11700F @ 2.50GHz
// BenchmarkT-16               3542           1331650 ns/op
// PASS
