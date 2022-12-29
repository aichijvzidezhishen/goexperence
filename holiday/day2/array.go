package main

import "fmt"

// 切片：长度可以动态扩张的数组
// 数组：长度固定的容器
//array[start:end] 从数组上截取片段，形成切片
//cap(slice)获得切片容量
//创建之初，容量等于长度
//扩张时，一旦容量无法满足需要，就一翻倍的策略扩容
//切片扩容的时候，所有元素的地址重新拷贝到一个连续的地址中

func mainp2() {
	var (
		array = [5]int{1, 2, 3, 4, 5}
	)
	slice := array[0:3]
	fmt.Printf("slice的类型是%T,值是%v\n", slice, slice)
	fmt.Printf("array的类型是%T,值是%v\n", array, array)
	slice = array[:]
	fmt.Printf("slice的类型是%T,值是%v\n", slice, slice)
}
func maina2() {
	var slice = []int{0, 11, 22, 33}
	fmt.Printf("slice的长度是%d   容量是%d", len(slice), cap(slice))
	slice = append(slice, 11)
	fmt.Printf("slice的长度是%d   容量是%d", len(slice), cap(slice))
}
func maina3() {
	var (
		// arr    = make([]int, 1, 1)
		slice  = []int{1, 2, 3}
		slice2 = []int{4, 5, 6, 7, 8, 9}
	)
	slice = append(slice, slice2...)
	fmt.Println("slice1", slice)

	fmt.Println("slice2", slice2)
	fmt.Printf("slice %d\nslice1 %d\n", cap(slice), cap(slice2))
	slice = append(slice, slice2...)
	fmt.Printf("slice %d\nslice1 %d", cap(slice), cap(slice2))

	// fmt.Printf("arr的长度是%d   容量是%d\n", len(arr), cap(arr))
	// for i := 1; i <= 10; i++ {
	// 	arr = append(arr, i)
	// 	fmt.Printf("追加第%d   arr的长度是%d   容量是%d\n", i, len(arr), cap(arr))
	// }
	fmt.Println()

}
func maina4() {
	var (
		arr    = [5]int{1, 2, 3, 4, 5}
		slice1 = arr[:]
		slice2 = slice1[:]
	)
	// arr,slice1,slice2具有完全相同的初始地址
	fmt.Printf("%p %p %p\n", &arr, &slice1, &slice2)
	fmt.Printf("%p %p %p\n", &arr[0], &slice1[0], &slice2[0])
	fmt.Printf("%p %p %p\n", &arr[1], &slice1[1], &slice2[1])
	//牵一发动全身
	arr[1] = 11
	slice1[2] = 22
	slice2[3] = 33
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
	// slice2扩容后，所有地址发生迁移（拷贝了一份）
	slice2 = append(slice2, 5)
	slice2[0] = 1000
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)

}
func maina5() {
	slice1 := make([]int, 0)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Println(slice1, len(slice1), cap(slice1))
	slcie2 := make([]int, 0, 4)
	fmt.Println(slcie2, len(slcie2), cap(slcie2))

}
