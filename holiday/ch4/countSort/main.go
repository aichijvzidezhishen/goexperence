package main

import "fmt"

func countSort(arr []int) []int {
	maxNum := max(arr)
	sortArr := make([]int, len(arr))
	counts := make([]int, maxNum+1)
	for _, v := range arr {
		counts[v]++
	}
	for i := 1; i <= maxNum; i++ {
		counts[i] += counts[i-1]
	}
	for _, v := range arr {
		sortArr[counts[v]-1] = v
		counts[v]--
	}
	return sortArr
}
func max(arr []int) int {
	a := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > a {
			a = arr[i]
		}
	}
	return a
}
func main() {
	var theArray = []int{5, 54, 12, 32, 45, 21, 23, 65, 1, 5, 1, 2, 5}
	fmt.Println("排序前：", theArray)
	fmt.Println("排序后：", countSort(theArray))
}
