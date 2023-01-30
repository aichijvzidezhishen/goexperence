package main

import (
	"fmt"
	"time"
)

/*
	固定时长定时器
*/
func main010() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())
	S := <-timer.C

	fmt.Println("", S)

}

func main012() {
	fmt.Println(time.Now())
	x := <-time.After(3 * time.Second)
	fmt.Println("", x)
}
