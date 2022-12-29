package main

import "fmt"

func mainfunc() {

	GetSum(1, 1, 1)
	b := GetSum2(1, 1, 1)
	c := GetSum3(1, 1, 2)
	sum, ok := GetSum4(1, 1, 2)
	fmt.Println("和为：", b)
	fmt.Println("f3de和为：", c)
	fmt.Printf("f3de和为：%d,是否奇偶%v", sum, ok)
}

//不定长参数
func GetSum(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("type is %T\n", a)
	fmt.Printf("value is %v", a)
}

//返回值
func GetSum2(a ...int) int {
	var sum = 0
	for _, v := range a {
		sum += v
	}
	return sum

}

//返回值
func GetSum3(a ...int) (sum int) {

	for _, v := range a {
		sum += v
	}
	return

}

//返回参数之和，和的奇偶性
func GetSum4(a ...int) (int, bool) {
	var sum int
	for _, v := range a {
		sum += v
	}

	var flag bool
	if sum%2 == 0 {
		flag = false
	} else {
		flag = true
	}
	return sum, flag

}
