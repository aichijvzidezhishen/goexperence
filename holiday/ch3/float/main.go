package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	for x := 0; x < 8; x++ {
		fmt.Printf("x= %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
