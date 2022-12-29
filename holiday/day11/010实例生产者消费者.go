package main

import (
	"fmt"
	"strconv"
	"time"
)

// 完全异步的生产者消费者模型
func main101() {
	// 商店管道，生产者和消费者传递的纽带
	chanstop := make(chan string, 100)
	go func() {
		for {
			product := strconv.Itoa(time.Now().Nanosecond())
			chanstop <- "产品" + product
			fmt.Println("生产了", product)
		}

	}()
	go func() {
		for {
			x := <-chanstop
			fmt.Println("消费了", x)
		}

	}()
	for {
		time.Sleep(time.Second)
	}

}
func main102() {
	// 商店管道，生产者和消费者传递的纽带
	chanstop := make(chan string, 100)
	chanstroe := make(chan string, 100)
	//生产者和消费者
	go producer(chanstroe)
	go logistics(chanstroe, chanstop)
	go customer(chanstop)
	for {
		time.Sleep(time.Second)
	}
}
func logistics(chanstroe, chanstop chan string) {
	ps := make([]string, 0)
	for v := range chanstroe {
		ps = append(ps, v)
		fmt.Println("物流装车", v)
	}
	fmt.Println("装车完毕")
	for _, v := range ps {
		chanstop <- v
		fmt.Println("货物抵达商店", v)
	}
}
func producer(chanstop chan string) {
	for i := 0; i < 10; i++ {
		product := strconv.Itoa(time.Now().Nanosecond())
		chanstop <- "产品" + product
		fmt.Println("生产了", product)
		time.Sleep(time.Second)
	}
	close(chanstop)
}

func customer(chanstop chan string) {
	for {
		x := <-chanstop
		fmt.Println("消费了", x)
		time.Sleep(time.Second)
	}
}
