package http

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// TestQPSCounter 测试QPS计数器的基本功能
func TestQPSCounter(t *testing.T) {
	// 创建QPS计数器
	counter := NewQPSCounter()
	defer counter.Close()

	// 创建一个测试用的HTTP处理器
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 包装处理器
	wrappedHandler := withQPSCounter(counter, testHandler)

	// 测试单请求计数
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)

	if atomic.LoadUint64(&counter.requestCounter) != 1 {
		t.Errorf("单请求计数错误，预期1，实际%d", counter.requestCounter)
	}

	// 测试多请求计数
	const reqCount = 100
	var wg sync.WaitGroup
	wg.Add(reqCount)

	for i := 0; i < reqCount; i++ {
		go func() {
			defer wg.Done()
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			wrappedHandler.ServeHTTP(w, req)
		}()
	}

	wg.Wait()

	if atomic.LoadUint64(&counter.requestCounter) != reqCount+1 { // +1是因为前面有一个单请求
		t.Errorf("多请求计数错误，预期%d，实际%d", reqCount+1, counter.requestCounter)
	}
}

// TestQPSCalculation 测试QPS计算的准确性
func TestQPSCalculation(t *testing.T) {
	// 创建QPS计数器
	counter := NewQPSCounter()
	defer counter.Close()

	// 创建一个测试用的HTTP处理器
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// 包装处理器
	wrappedHandler := withQPSCounter(counter, testHandler)

	// 记录开始时间
	startTime := time.Now()

	// 在1秒内发送尽可能多的请求
	var wg sync.WaitGroup
	stopChan := make(chan struct{})

	// 启动多个goroutine发送请求
	const goroutineCount = 10
	wg.Add(goroutineCount)

	for i := 0; i < goroutineCount; i++ {

		go func() {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					req := httptest.NewRequest("GET", "/", nil)
					w := httptest.NewRecorder()
					wrappedHandler.ServeHTTP(w, req)
				}
			}
		}()
	}

	// 运行1秒后停止
	time.Sleep(1 * time.Second)
	close(stopChan)
	wg.Wait()

	// 等待QPS计算完成
	time.Sleep(100 * time.Millisecond)

	// 获取计算出的QPS
	calculatedQPS := counter.GetQPS()

	// 计算实际请求数
	totalRequests := atomic.LoadUint64(&counter.requestCounter)
	elapsedSeconds := time.Since(startTime).Seconds()
	actualQPS := uint64(float64(totalRequests) / elapsedSeconds)

	// 允许一定的误差范围（10%）
	tolerance := actualQPS / 10
	if calculatedQPS < actualQPS-tolerance || calculatedQPS > actualQPS+tolerance {
		t.Errorf("QPS计算不准确，预期约%d，实际%d", actualQPS, calculatedQPS)
	}

	t.Logf("测试完成 - 实际请求数: %d, 实际QPS: %d, 计算QPS: %d",
		totalRequests, actualQPS, calculatedQPS)
}

// 世界变得
