package main

import (
	"fmt"
	"os"
)

func maincommand() {
	fmt.Println(os.Args)
	if len(os.Args) != 2 {
		fmt.Println("input filename")
		return
	}
	fmt.Println(os.Args[1])
}
