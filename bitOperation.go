package main

import (
	"fmt"
	"strconv"
)

var A int32

// 设置二进制的某一位为1
func SetBitValue(position int32) {
	A |= (1 << (position))
	fmt.Println("a", A)
}

//十进制转换成二进制
func DectoBin(d int64) {
	bin := strconv.FormatInt(d, 2)
	fmt.Println("bin", bin)
}

//判断二进制的某一位是否是1
func IsOne(x int, n int) bool {
	return (x>>(n-1))&1 == 1
}

//获取二进制中1的个数
func BitCount1(n int) int {
	num := n

	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
