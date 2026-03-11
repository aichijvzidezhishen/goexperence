package http

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

type QPSCounter struct {
	requestCounter uint64
	lastCount      uint64
	qps            uint64
	ticker         *time.Ticker
}

func NewQPSCounter() *QPSCounter {
	counter := &QPSCounter{
		ticker: time.NewTicker(1 * time.Second),
	}

	go counter.start()
	return counter
}

func (c *QPSCounter) start() {
	for range c.ticker.C {
		// 获取当前请求数
		current := atomic.LoadUint64(&c.requestCounter)

		c.qps = current - c.lastCount

		c.lastCount = current

		fmt.Printf("[%s] Current QPS: %d\n", time.Now().Format("2006-01-02 15:04:05"), c.qps)
	}
}

func (c *QPSCounter) Incr() {
	atomic.AddUint64(&c.requestCounter, 1)
}

func (c *QPSCounter) GetQPS() uint64 {
	return atomic.LoadUint64(&c.qps)
}

func (c *QPSCounter) Close() {
	c.ticker.Stop()
}

func withQPSCounter(counter *QPSCounter, hander http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter.Incr()

		//
		hander(w, r)
	}
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	// 模拟网关处理逻辑
	time.Sleep(time.Millisecond * time.Duration(10+randInt(0, 40))) // 随机延迟10-50ms
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func randInt(min, max int) int {
	return min + int(time.Now().UnixNano()%int64(max-min+1))
}
