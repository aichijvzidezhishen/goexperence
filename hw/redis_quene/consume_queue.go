package redisquene

import (
	"context"
	"log"
	"sync"
)

func consumeQueue(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := RedisClient.BLPop(context.Background(), 0, "msgQueue").Result()
			if err != nil {
				log.Printf("Error consuming message: %v", err)
				continue
			}
			log.Printf("recv message: %v", msg[1])
		}
	}
}
