package main

import (
	"fmt"
	"holiday/ch4/jianzhi/solve"
)

func main() {
	var (
		// arr         = []int{6, 3, 2, 0, 0, 5, 2}
		arr = []int{6, 3, 2, 0, 2, 5, 0}
		// arr         = []int{}
		duplication = [1]int{}
	)
	var flag = solve.Duplicate(arr, &duplication)
	fmt.Println("flag = ", flag)
}
