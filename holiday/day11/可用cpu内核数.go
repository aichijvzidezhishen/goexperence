package main

import (
	. "fmt"
	"runtime"
)

func main03() {
	Println("可用内核数", runtime.NumCPU())
	//将可用核数设置为1，并返回先前的设置
	Println("设置CPU可用内核数", runtime.GOMAXPROCS(1))
	Println("设置CPU可用内核数", runtime.GOMAXPROCS(40))
}
