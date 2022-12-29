package main

import "fmt"

func main() {
	n1 := 10
	test1(n1)
	fmt.Println("main", n1)
	num := 20
	test3(&num)
	fmt.Println("main", num)
	test4 := getSum
	fmt.Printf("test4的类型%T，getSum的类型是%T", test4, getSum)
	res := getSum(10, 40)
	fmt.Println("res", res)
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)
	//将匿名函数给一个变量（函数变量），在通过变量来调用匿名函数
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(10, 20)
	fmt.Println("res2=", res2)
	res3 := a(30, 40)
	fmt.Println("res3=", res3)
	//全局匿名函数
	var (
		func1 = func(n1 int, n2 int) int {
			return n1 * n2
		}
	)
	res5 := func1(4, 9)
	fmt.Println(res5)

}

func test1(n1 int) {
	n1 = n1 + 1
	fmt.Println("test", n1)
}
func test3(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test3 n1 = ", *n1)
}
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
