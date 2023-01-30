package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i int32 = 100
	// var n1 float32 = float32(i)
	// var n2 int8 = int8(i)
	// var n3 int64 = int64(i)
	// fmt.Printf("i=%v n1=%v n2=%v n3=%v \n", i, n1, n2, n3)
	// fmt.Printf("i type is %T\n", i)

	// // var b bool = true
	// // var char byte = 'a'

	// // %d  int
	// // %c byte
	// // %f float
	// // %t bool
	// var (
	// 	num1 int     = 99
	// 	num2 float64 = 23.456
	// 	str  string
	// 	str1 string
	// 	str2 string
	// 	str3 string
	// 	str4 string
	// 	b    bool
	// 	err  error
	// 	// a    int
	// )
	// str = fmt.Sprintf("%d", num1)
	// fmt.Printf("数据类型：%T str=%q", str, str)
	// str1 = strconv.FormatInt(int64(num1), 10)
	// fmt.Printf("数据类型：%T str=%q", str1, str1)
	// str2 = strconv.FormatFloat(num2, 'f', 10, 64)
	// fmt.Printf("数据类型：%T str=%q", str2, str2)
	// var num3 int64 = 4567
	// str3 = strconv.Itoa(int(num3))
	// fmt.Printf("str3数据类型：%T str=%q\n", str3, str3)
	// b, err = strconv.ParseBool(str4)
	// fmt.Println(err)
	// fmt.Printf("b数据类型：%T b=%v", b, b)

	var (
		// b   bool
		str string = "123145"
		n1  int32
		n2  float32
		// n3  byte
		n4 int64
		n5 int8
		n6 float64
	)
	n5 = int8(n1) + 127
	fmt.Println(n5)
	n2 = float32(n1)
	n4 = int64(n1)
	fmt.Printf("n1=%v n2=%v n4=%v", n1, n2, n4)
	//被转换的是变量存储的数值，变量本身的数据类型不变
	fmt.Printf("n1 tyoe is %T", n1)
	// radix:进制数
	str = strconv.FormatInt(int64(num), radix)
	// n6只能是float64，f表示格式，10表示保留小数的位数，32表示float32浮点数
	str = strconv.FormatFloat(n6, 'f', 10, 32)
	//b为布尔型
	str = strconv.FormatBool(b)
	//num4是int64   ()内为int型
	str = strconv.Itoa(int(num4))
	//n4为int64
	n4, _ = strconv.ParseInt(str, 10, 64)
	//n6为float64
	n6, _ = strconv.ParseFloat(str, 64)
	b, _ = strconv.ParseBool(str)

}
