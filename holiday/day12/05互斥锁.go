package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
	何为数据的并发安全
*/
func main051() {
	money := 2000
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100000; j++ {
				money += 1
			}
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("最终金额", money)
}
func main052() {
	money := 2000
	// 一只麦克
	var mt sync.Mutex
	wg.Add(1)
	go func() {
		fmt.Println("1号开始抢麦克")
		// 谁抢到麦才有资格修改money，其他抢麦的协程必须阻塞等待
		mt.Lock()
		fmt.Println("1号抢到麦克")
		money -= 100
		<-time.After(3 * time.Second)
		mt.Unlock()
		fmt.Println("1号弃麦")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("2号开始抢麦克")
		// 谁抢到麦才有资格修改money，其他抢麦的协程必须阻塞等待
		mt.Lock()
		fmt.Println("2号抢到麦克")
		money -= 150
		<-time.After(3 * time.Second)
		mt.Unlock()
		fmt.Println("2号弃麦")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("最终金额", money)
}
