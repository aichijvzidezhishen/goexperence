package main

import "fmt"

/*桶排序*/
func sortInBucket(bucket []int) {
	if len(bucket) == 1 {
		return
	}
	for i := 1; i < len(bucket); i++ {
		temp := bucket[i]
		j := i - 1
		for j >= 0 && temp < bucket[j] {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = temp
	}
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

func Sort(arr []int) []int {
	num := len(arr)
	max := max(arr)
	buckets := make([][]int, num)
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}
	tem := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(arr[tem:], buckets[i])

			tem += bucketLen
		}
	}
	return arr

}

func main() {
	var theArray = []int{5, 54, 12, 32, 45, 21, 23, 65, 1, 5, 1, 2, 5}
	fmt.Println("排序前：", theArray)
	Sort(theArray)
	fmt.Println("排序后：", theArray)
}
