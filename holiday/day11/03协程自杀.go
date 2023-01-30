package main

import (
	. "fmt"
	"runtime"
	"time"
)

func main031() {
	go func() {
		// 1
		Println("大梦谁先觉")
		task031()
		//协程被杀死
		Println("----诸葛亮")
	}()
	time.Sleep(time.Second)
}
func task031() {
	// 4
	defer Println("交党费")
	// 2
	Println("平生我自知")
	// 3
	Println("草堂秋睡足")
	// 杀死所在协程
	runtime.Goexit()

	Println("----窗外日迟迟")
}
