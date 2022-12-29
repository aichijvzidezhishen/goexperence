package main

import (
	"fmt"
	"time"
)

func main031() {
	timer := time.NewTimer(1 * time.Second)
	fmt.Println(time.Now())

	time.Sleep(2 * time.Second)
	//趁定时器还没到时间，重置时间
	// 如果已经超时重置无效
	timer.Reset(3 * time.Second)
	fmt.Println("炸弹引包与", <-timer.C)
}
func main032() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())
	timer.Stop()
	// fatal error: all goroutines are asleep - deadlock!
	// 定时器已被叫停时间管道永远读不出数据 --死锁
	t := <-timer.C
	fmt.Println(t)
}
