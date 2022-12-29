package main

import (
	"fmt"
	"time"
)

func main071() {
	ch := make(chan int)
	go func() {
		x := <-ch
		fmt.Println("---", x)
	}()
	ch <- 123
	time.Sleep(time.Second)
	fmt.Println("over")

}

func main072() {
	/*
		不能关闭一个没有初始化的管道 panic!
	*/
	// var ch chan int
	// // panic: close of nil channel
	// close(ch)
	/*
		管道不能重复关闭 panic!
	*/
	// ch := make(chan int)
	// close(ch)
	// // panic: close of closed channel
	// close(ch)
	/*
		向已经关闭的管道写信息  panic!
	*/
	// ch := make(chan int)
	// close(ch)
	// // panic: send on closed channel
	// ch <- 123

	/*
		不能向已经关闭的管道写数据，可以从一个已经关闭的管道中读数据,已经读完，继续读取，读出零值
	*/
	ch := make(chan int, 3)
	ch <- 123
	close(ch)

	go func() {
		x := <-ch
		fmt.Println("read1", x)
		x = <-ch
		fmt.Println("read2", x)
		x, ok := <-ch
		fmt.Println("read3", x, ok)
	}()
	time.Sleep(time.Second)
	fmt.Println("over")
}

/*
	校验管道数据的有效性
*/
func main073() {
	ch := make(chan int, 3)
	ch <- 123
	ch <- 1235
	close(ch)

	go func() {
		for i := 0; i < 3; i++ {
			if x, ok := <-ch; ok {
				fmt.Println("读数据", x)
			} else {
				fmt.Println("数据已读尽！")
			}
		}

	}()
	time.Sleep(time.Second)
}

/*
遍历管道中的全部数据
*/
func main074() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	go func() {

	}()
}
func main075() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("写入", i)
			time.Sleep(time.Second)
		}
		close(ch)
		//关闭管道时会向所有的读取协程发送通知，通知其不要在遍历了
		fmt.Println("channel close，写入结束")
	}()
	go func() {
		// 管道close时结束遍历
		for v := range ch {
			fmt.Println(v)
		}
		fmt.Println("读取结束")
	}()
	time.Sleep(6 * time.Second)
}
