package main

import "fmt"

/*
	基数排序
*/
func radixSort(arr []int) []int {
	maxNum := max(arr)
	for bit := 1; maxNum/bit > 0; bit *= 10 {
		arr = Sort(arr, bit)
	}
	return arr
}
func Sort(arr []int, bit int) []int {
	n := len(arr)
	fmt.Println("n", n)

	//对应位相同的数存放在counts
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10
		counts[num]++
	}
	for i := 1; i < n; i++ {
		counts[i] += counts[i-1]
	}
	temp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		temp[counts[num]-1] = arr[i]
		counts[num]--
	}
	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}
	return arr
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
	radixSort(theArray)
	fmt.Println("排序后：", theArray)
}
