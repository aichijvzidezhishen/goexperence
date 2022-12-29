package main

import "fmt"

func main() {
	//for循环的其他用法
	// j := 1
	// for j <= 10 {
	// 	fmt.Println("8888")
	// 	j++
	// }
	// k := 1
	// for {
	// 	if k < 10 {
	// 		fmt.Println("----")
	// 	} else {
	// 		break
	// 	}
	// 	k++
	// }
	a := [3]int{1, 2, 3}
	for _, v := range a { //复制一份a遍历[1, 2, 3]
		v += 100 //v是复制对象中的值，不会改变a数组元素的值
	}
	fmt.Println(a)        //1 2 3
	for i, v := range a { //i,v从a复制的对象里提取出
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a) //输出[1 200 300]
		}
		a[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
	}
	fmt.Println(a) //输出[101, 102, 103]
	// fmt.Println()
	type AStruct struct {
		bar string
	}
	list := []AStruct{
		{"1"},
		{"2"},
		{"3"},
	}
	copyed := make([]*AStruct, len(list))
	for i, v := range list {
		copyed[i] = &v
	}
	fmt.Println(list[0], list[1], list[2])
	fmt.Println(copyed[0], copyed[1], copyed[2])
}
