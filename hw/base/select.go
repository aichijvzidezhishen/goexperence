package base

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 多通道选择
func MultiChanChoose() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello"
	}()

	select {
	case value := <-ch1:
		fmt.Println("Received from ch1:", value)
	case value := <-ch2:
		fmt.Println("Received from ch2:", value)
	}
}

func TimeoutContro() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 42
	}()

	select {
	case value := <-ch:
		fmt.Println("Received value:", value)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}

// 信号推出机制
func OsExitNotify() {
	exitSignChan := make(chan os.Signal, 1)
	signal.Notify(exitSignChan, syscall.SIGINT, syscall.SIGALRM)
	otherChan := make(chan bool)
	go func() {
		//模拟其他进程
		otherChan <- true
	}()
	select {
	case <-exitSignChan:
		//
		os.Exit(0)
	case <-otherChan:
		//
		fmt.Println("recv otherChan")
	}
	// 主goroutine继续执行其他逻辑
	// ...
}

//服务器优雅启停
