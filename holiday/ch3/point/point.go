package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i address", &i)
	var ptr *int = &i
	fmt.Printf("ptr=%v", ptr)
	fmt.Printf("ptr address=%v", &ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
}
