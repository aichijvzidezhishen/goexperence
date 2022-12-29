package main

import (
	"fmt"
	"time"
)

//匿名函数使用场景
// 一次性调用
// 有必要进行封装
func mainanoy1() {
	sum := func(a, b int) int {
		return a + b
	}(3, 5)
	fmt.Println(sum)
}

//匿名函数场景1：延时一个任务
func mainanoy2() {
	defer func(count int) {
		for i := 0; i < count; i++ {
			fmt.Println(32)
		}
	}(5)
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

}

//匿名函数场景2 :并发一个一次性任务
func mainanoy3() {
	go func() {
		//每半秒
		for i := 2; i < 11; i += 2 {
			fmt.Println("并发协程", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("game over")
}
