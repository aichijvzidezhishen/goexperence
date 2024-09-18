package main

import "fmt"

func main() {
	NumberOf1(5)
}
func NumberOf1(n int) int {
	ans := 0
	for n != 0 {
		n = n & int(int32(n)-1)
		fmt.Println("n--", n)
		ans++
	}
	return ans
}
