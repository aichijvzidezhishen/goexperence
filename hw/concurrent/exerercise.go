package concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// 题目：启动两个 Goroutine，分别打印奇数和偶数，按顺序输出 1-10。

func OddAndEven() {
	var wg sync.WaitGroup
	wg.Add(2)
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-oddChan
			fmt.Println(i)
			evenChan <- struct{}{} // 发送一个信号，通知偶数 goroutine 可以执行了
		}
		<-oddChan // 等待最后一个奇数 goroutine 执行完毕
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-evenChan
			fmt.Println(i)
			if i != 10 {
				oddChan <- struct{}{} // 发送一个信号，通知奇数 goroutine 可以执行了
			}
		}

		wg.Done()
	}()

	// 启动第一个 goroutine
	oddChan <- struct{}{}
}

// 题目：使用并发的埃拉托斯特尼筛法，找出 100 以内的所有素数。
func GeneratePrimes() {
	gen := func(n int) <-chan int {
		out := make(chan int)
		go func() {
			for i := 2; i <= n; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	}
	filter := func(in <-chan int, prime int) <-chan int {
		out := make(chan int)
		go func() {
			for num := range in {
				// i := <-in
				if num%prime != 0 {
					out <- num
				}
			}
			close(out)
		}()
		return out
	}
	n := 100

	in := gen(n)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			prime, ok := <-in
			if !ok {
				break
			}
			fmt.Println(prime, " ")
			in = filter(in, prime)
		}
	}()
	wg.Wait()
}

//题目：使用 Goroutine 并发计算斐波那契数列的前 20 项。

// 定义一个函数Fibonacci，接收一个整数n作为参数
func Fibonacci(n int) {
	// 创建一个整数类型的通道ch
	ch := make(chan int)
	// 创建一个WaitGroup类型的变量wg，用于等待goroutine完成
	var wg sync.WaitGroup
	// 将wg的计数器加1
	wg.Add(1)

	// 启动一个goroutine，执行匿名函数
	go func(n int, ch chan<- int, wg *sync.WaitGroup) {
		// 在函数结束时，将wg的计数器减1
		defer wg.Done()
		// 初始化两个变量a和b，分别表示斐波那契数列的前两个数
		a, b := 0, 1
		// 循环n次，计算斐波那契数列
		for i := 0; i < n; i++ {
			// 将a的值发送到通道ch中
			ch <- a
			// 更新a和b的值
			a, b = b, a+b

		}

		// 关闭通道ch
		close(ch)
	}(n, ch, &wg)
	// 等待wg的计数器减为0，即所有goroutine都完成
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Wait()
	}()
}

// 题目：并发读取多个文件内容并统计行数。
func CountLinesInFiles(files []string) {
	//
	statistics := func(filename string, ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(filename, "data", string(data))
		lines := strings.Split(string(data), "\n")
		ch <- len(lines)
	}
	ch := make(chan int)
	var wg sync.WaitGroup
	for _, filename := range files {
		wg.Add(1)
		go statistics(filename, ch, &wg)
	}

	go func() {

		wg.Wait() // 等待所有 goroutine 完成
		close(ch)
	}()
	// 统计行数
	var total int
	for v := range ch {
		fmt.Println("file length:", v)
		total += v
	}
	fmt.Println("Total lines:", total)
}

// 题目：并发发送多个 HTTP 请求，返回最快的响应（超时控制）。
func FastestResponse() {
	singleFetch := func(url string, ch chan<- string) { // 单个请求
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		ch <- fmt.Sprintf("url:%s, status code:%d, length:%d", url, resp.StatusCode, len(body))
		// fmt.Println("url:", url, "status code:", resp.StatusCode)
	}
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.yahoo.com",
	}

	ch := make(chan string)
	for _, url := range urls {
		go singleFetch(url, ch)
	}

	// 等待	所有 goroutine 完成
	select {
	case result := <-ch:
		fmt.Println("result :", result)
	case <-time.After(10 * time.Second):
		fmt.Println("timeout")
	}
}

// 题目：使用扇出 / 扇入模式并发处理数据。
/* func FanOutFanIn() {
	producer := func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}
	worker := func(id int, in <-chan int, out chan<- int) {
		for v := range in {
			out <- v * id
		}
		close(out)
	}
	consumer := func(in <-chan int) {
		for v := range in {
			fmt.Println(v)
		}
	}
	ch := make(chan int)
	go producer(ch)
	var wg sync.WaitGroup
	wg.Add(2)

	// 扇入
} */

// 并发处理任务错误
func ConcurrentErrorHandling() {
	var (
		wg   sync.WaitGroup
		ch   = make(chan error, 10)
		errs []error
		mu   sync.Mutex
	)
	taskFunc := func(id int) error {
		if id%5 == 0 {
			return fmt.Errorf("task %d faild", id)
		}
		fmt.Println("Task completed", id)
		return nil
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := taskFunc(id); err != nil {
				ch <- err
			}
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for err := range ch {
		mu.Lock()
		errs = append(errs, err)
		mu.Unlock()
	}

	// handle
	if len(errs) > 0 {
		fmt.Println("Errors occurred:", errs)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}
