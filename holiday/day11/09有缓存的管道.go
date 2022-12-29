package main

import (
	"fmt"
	"time"
)

// 写入阻塞
func main091() {
	ch := make(chan int, 3)
	go func() {
		ch <- 1
		fmt.Println("write", 1)
		ch <- 2
		fmt.Println("write", 2)
		ch <- 3
		fmt.Println("write", 3)
		// 缓存能力已经耗尽了，没人读就再也写不进去了（阻塞）
		ch <- 4
		fmt.Println("write", 4)
	}()
	time.Sleep(time.Second)
}

// 读出阻塞
func main092() {
	ch := make(chan int, 3)
	go func() {
		x := <-ch
		fmt.Println("read", x)
		x = <-ch
		fmt.Println("read", x)
		x = <-ch
		fmt.Println("read", x)
		x = <-ch
		fmt.Println("read", x)
	}()
	time.Sleep(time.Second)
}
func main093() {
	ch := make(chan int, 3)
	fmt.Println("元素个数", len(ch), "缓存能力", cap(ch))
	ch <- 123
	fmt.Println("元素个数", len(ch), "缓存能力", cap(ch))
	ch <- 1234
	fmt.Println("元素个数", len(ch), "缓存能力", cap(ch))
	ch <- 1235
	fmt.Println("元素个数", len(ch), "缓存能力", cap(ch))
	/*
		主协程阻塞永远造成死锁
		fatal error: all goroutines are asleep - deadlock!

	*/
	ch <- 12352
	fmt.Println("元素个数", len(ch), "缓存能力", cap(ch))
}
