package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var runes []rune
	for i, r := range "Hello, 世界" {
		runes = append(runes, r)
		fmt.Printf("%d\n%q", i, runes)
	}
	var x []int
	x = append(x, 1, 2, 3, 4, 5, 6)
	// [1 2 3 4 5 6 1 2 3 4 5 6]
	x = append(x, x...) //将x数组复制到x数组中
	fmt.Println(x)
	var b, a []int
	//
	for i := 0; i < 10; i++ {
		a = appendInt(b, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(a), a)
		b = a
	}
	s := []int{5, 6, 7, 8, 9}
	var (
		arr1 [5]int
		arr2 = [3]int{1, 2, 3}
		arr3 = [...]int{8, 9, 10}
		// //类型推导
		// strarr = [...]string{1: "tom", 0: "jack", 2: "mary"}
		// arr4 = [...]int{1: 800, 0: 900, 2: 999}
	)
	test(&arr2)
	fmt.Println(arr2)
	test1(arr3)
	fmt.Println("ar3", arr3)
	fmt.Println(remove(s, 2))
	fmt.Println(remove_optimzation(s, 2))
	var char [26]byte
	for i := 0; i < 26; i++ {
		char[i] = 'A' + byte(i)
	}
	len := len(arr1)
	//每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr1[i] = rand.Intn(100)
	}

	fmt.Println("before swap", arr1)
	for i := 0; i < len/2; i++ {
		temp := arr1[i]
		arr1[i] = arr1[len-1-i]
		arr1[len-1-i] = temp
	}
	fmt.Println("after swap", arr1)
	//slice从底层来说 其实就是一个数据结构。
	// type slcie struct {
	// 	ptr *[2]int
	// 	len int
	// 	cap int
	// }

	//
	//使用
}

func test(arr *[3]int) {
	(*arr)[0] = 5
}
func test1(arr [3]int) {
	arr[0] = 100
	fmt.Println("test1的arr", arr)
}
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built‐in function; see text
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func remove_optimzation(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
