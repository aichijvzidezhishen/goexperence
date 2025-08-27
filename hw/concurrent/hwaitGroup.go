package concurrent

import (
	"fmt"
	"sync"
)

/*
避免重复调用wait
传递函数需要传递指针，内部函数需要修改其状态
*/
func proessMutiItems() {
	items := []int{1, 2, 3, 4}
	var wg sync.WaitGroup
	wg.Add(len(items))
	for _, item := range items {
		go func(item int) {
			defer wg.Done()
			// do something
			fmt.Println("process item:", item)
		}(item)
	}

	//
	wg.Wait()
	fmt.Println("all done")
}

func NestingWaitGroup() {
	var wg sync.WaitGroup
	go childTask(&wg)
	wg.Wait()
	fmt.Println("all done")
}

func childTask(wg *sync.WaitGroup) {
	defer wg.Done()
	var childWg sync.WaitGroup
	defer childWg.Done()

	childWg.Add(2)

	go func() {
		defer childWg.Done()
		// do something
		fmt.Println("child task 1")
	}()

	go func() {
		defer childWg.Done()
		// do something
		fmt.Println("child task 2")
	}()

	childWg.Wait()
	fmt.Println("child task done")
}
