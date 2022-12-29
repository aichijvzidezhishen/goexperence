package main

import (
	"fmt"
	"sync"
	"time"
)

func main041() {
	// 创建任务等待组
	var wait sync.WaitGroup
	// 向等待组中添加任务
	wait.Add(1)
	wait.Add(1)
	wait.Add(1)
	// 从等待组中抹掉任务
	wait.Done()
	wait.Done()
	wait.Done()
}
func main042() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			<-time.After(1 * time.Second)
			fmt.Println(i)
		}
		wait.Done()
	}()
	wait.Add(1)
	go func() {
		timer := time.NewTicker(1 * time.Second)
		var i int
		for {
			<-timer.C
			i++
			fmt.Println("秒表", i)
			if i > 9 {
				break
			}
		}
		wait.Done()
	}()
	wait.Wait()
}
