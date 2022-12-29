package main

import (
	"fmt"
	"unsafe" //提供了跳过go语言类型安全限制的操作
)

func main() {
	var i int
	fmt.Println("i=", i)
	n1, n2, n3 := 100, "tom", 888
	//多变量声明
	var (
		n4 = 900
		n5 = "marry"
	)
	fmt.Println("n1", n1, "n2", n2, "n3", n3, "n4", n4, "n5", n5)
	var (
		c byte = 255
		b uint = 1
	)
	fmt.Println("c=", c, "b=", b)
	//查看某个变量的数据类型
	//fmt.Printf()可以用作格式化输出
	fmt.Printf("n5的类型 %T \n", n5)
	//查看某个变量的占用字节大小和数据类型
	fmt.Printf("n5的类型 %T n5占用的字节数%d \n", n5, unsafe.Sizeof(n5))
}
