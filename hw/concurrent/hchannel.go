package concurrent

import (
	"fmt"
	"sync"
	"time"
)

// 空读写阻塞，写关闭异常，读关闭空零
// send to nil channel
func SendToNil() {
	ch := make(chan int)
	// close(ch)
	ch <- 1
	//panic: send on closed channel
}

// recv from nil channel
func RecvFromNil() {
	ch := make(chan int)
	<-ch
	//
}

// send to close nil channel

// recv from close nil channel

// 无缓冲的channel是同步的，而有缓冲的channel是非同步的

// 消息传递

// 同步控制

// 控制并发数量
func work(id int, jobs <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d start job %d\n", id, j)
		// time.Sleep(time.Second)
		res <- j * 2
		fmt.Printf("worker %d end job %d\n", id, j)
	}
	close(res)
}

func ImplWorker() {
	jobs := make(chan int, 5)
	res := make(chan int, 5)

	var wg sync.WaitGroup

	//
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go work(i, jobs, res, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	wg.Wait()
}

// 实现超时控制
func doWork(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "work done !"
}

func HandleTimeout() {
	ch := make(chan string)
	go doWork(ch)
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}

// 发布订阅模式
func publisher(pub chan string) {
	msgs := []string{"1", "2", "3"}
	for _, msg := range msgs {
		pub <- msg
		time.Sleep(time.Second)
	}
	close(pub)
}

func subscriber(id int, sub <-chan string) {
	for msg := range sub {
		fmt.Printf("sub %d get %s\n", id, msg)
	}
}

func ImplPAndSMod() {
	pub := make(chan string)
	go publisher(pub)

	// open two sub
	for i := 1; i <= 2; i++ {
		go subscriber(i, pub)

	}
	time.Sleep(4 * time.Second)
}
