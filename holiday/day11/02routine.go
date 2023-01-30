package main

import (
	. "fmt"
	"runtime"
	"strconv"
	"time"
)

/*
	出让协程资源
*/
func main02() {

	for i := 0; i < 10; i++ {
		go func(index int) {
			task("子协程" + strconv.Itoa(index))
		}(i)
	}
	time.Sleep(5 * time.Second)

}
func task(name string) {
	for i := 0; i < 10000; i++ {
		if name == "子协程5" {
			//优先级排到最后但并不是一点机会都没有
			runtime.Gosched()
		}
		Println(name, i)
	}
}
