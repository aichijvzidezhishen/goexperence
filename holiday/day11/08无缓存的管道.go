package main

import (
	"fmt"
	"time"
)

/*
	无缓存的管道没人读就永远写不进（阻塞）
	没人写就永远读不出


*/
func main081() {
	//创建没有缓存的通道
	ch := make(chan int)
	go func() {
		ch <- 132
		fmt.Println("数据已写入")
	}()
	go func() {
		x := <-ch
		fmt.Println("数据读出", x)
	}()
	time.Sleep(5 * time.Second)
}
