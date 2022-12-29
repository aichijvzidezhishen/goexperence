package main

import (
	"fmt"
	"time"
)

func main021() {

	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		s := <-ticker.C
		i++
		fmt.Println(i, s)
		if i == 5 {
			//停掉秒表会导致ticker永远无法读出数据，执着要读出会导致死锁
			ticker.Stop()
			break
		}
	}
	fmt.Println("over")
}
