package main

import (
	"fmt"
	"runtime"
	"time"
)

func main041() {
	go wildman()
	time.Sleep(5 * time.Second)
	//主协程自杀=>子协程失去约束
	runtime.Goexit()
}
func wildman() {
	for {
		fmt.Println("野人")
		time.Sleep(time.Second)
	}
}
