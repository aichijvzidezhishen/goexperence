package main

import (
	"fmt"
	"math"
)

/*
处理恐慌
*/

func getArea(radius float64) float64 {
	//挂起一个延时执行的恐慌处理程序
	// 即使程序因为恐慌而死，也能在此复活并得到致死的error
	defer func() {
		// 如果没有恐慌，error为空

		if err := recover(); err != nil {
			fmt.Println("err=\n", err)
		}
	}()
	if radius < 0 {
		panic("半径不能为负数！")
	}
	area := math.Pi * math.Pow(radius, 2)
	return area
}

func main084() {
	area := getArea(-1)
	fmt.Println("area", area)
}
