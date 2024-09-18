package redisquene

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Test_queueMessages(t *testing.T) {
	err := producrMessages("msg1")
	if err != nil {
		t.Error(err)
	}
}

func Test_consumeQueue(t *testing.T) {
	err := producrMessages("msg2")
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumeQueue(ctx, &wg)

	time.Sleep(1 * time.Second)
}
