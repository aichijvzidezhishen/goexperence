package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// runtime.GOMAXPROCS(2)
	// wg := sync.WaitGroup{}
	// wg.Add(20)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}()
	// }
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("i: ", i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// out := make(chan int)
	// wg := sync.WaitGroup{}
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 5; i++ {
	// 		out <- rand.Intn(5)
	// 	}
	// 	close(out)
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := range out {
	// 		fmt.Println(i)
	// 	}
	// }()
	// wg.Wait()

	random := make(chan int)
	done := make(chan bool)

	go func() {
		for {

			num, ok := <-random
			if ok {
				fmt.Println(num)
			} else {
				done <- true
			}
		}
	}()

	go func() {
		defer close(random)
		for i := 0; i < 5; i++ {
			random <- rand.Intn(5)
		}
	}()
	<-done
	close(done)
}
