package main

import "fmt"

func main1() {
	var (
		x   = 123
		ptr *int
	)

	ptr = &x
	fmt.Println(*ptr, ptr)
}
func mainp1() {
	var (
		x        = 123
		ptr *int = &x
	)
	mptr := &ptr
	fmt.Printf("mptr的类型是%T\n", mptr)
	fmt.Println(*ptr, *(*mptr))
}
