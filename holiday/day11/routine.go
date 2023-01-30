package main

import(
	"fmt"
	"time"
	"strconv"
)

func main01(){
	// 在一条独立的协程中执行匿名函数
	// go func(){
	// 	for{
	// 		fmt.Println("child routine")
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	for i := 0; i < 1000000; i++ {
		go dosomething("小分队"+strconv.Itoa(i))
	}
	//主协程死，子协程也死
	
	for{
		fmt.Println("main routine")
		time.Sleep(time.Second)
	}
}
func dosomething(s string){
	for{
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}