package main

import (
	"fmt"
	"strings"
)

func main() {
	f := Addupper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	f1 := makeSuffix(".jpg")
	fmt.Println(f1("aa.jpg"))
	res := sum(10, 20)
	fmt.Println("res = ", res)

}
func sum(n1 int, n2 int) int {
	//当执行defer时，暂时不执行，会将defer后面的语句压入独立的栈（defer栈）
	//函数执行完毕后，按照先入后出的方式出栈
	defer fmt.Println("ok1 n1 = ", n1)
	defer fmt.Println("ok2 n2 = ", n2)
	res := n1 + n2
	fmt.Println("ok3 n3 = ", res)
	return res
}
func Addupper() func(int) int {
	var n int = 10
	var str string = "hello"
	return func(x int) int {
		n += x
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func test() {

}
