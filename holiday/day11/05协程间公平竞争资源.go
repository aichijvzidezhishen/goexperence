package main

import (
	"fmt"
	"runtime"
	"time"
)

func main051() {
	go task051()

	go task052()

	time.Sleep(time.Second)
}
func task051() {
	for _, v := range "今生注定" {
		fmt.Printf("%c\n", v)

		runtime.Gosched()
		// time.Sleep(time.Nanosecond)
	}
}
func task052() {
	for _, v := range "FUCK" {
		fmt.Printf("%c\n", v)
		runtime.Gosched()
		// time.Sleep(time.Nanosecond)
	}
}
