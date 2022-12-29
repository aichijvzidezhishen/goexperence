package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

/*
	生产完了一次性消费
*/
func main() {
	chanshop := make(chan string, 100)
	chantel := make(chan int)
	go Producer(chanshop, chantel)
	// go logistics(chanstroe, chanshop)
	go Customer(chanshop, chantel)
	for {
		time.Sleep(time.Second)
	}
}
func Producer(chanshop chan string, chantel chan int) {
	// 强大的生产能力
	runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		product := strconv.Itoa(time.Now().Nanosecond())
		chanshop <- "产品" + product
		fmt.Println("生产了", product)
		time.Sleep(time.Second)
	}
	// 打电话
	chantel <- 125855445
	fmt.Println("呼出电话")
}
func Customer(chanshop chan string, chantel chan int) {
	runtime.GOMAXPROCS(1)
	// 阻塞等待电话
	a := <-chantel
	fmt.Println("收到来电", a)
	fmt.Println("")
	for {
		x := <-chanshop
		fmt.Println("消费了", x)
	}
}
