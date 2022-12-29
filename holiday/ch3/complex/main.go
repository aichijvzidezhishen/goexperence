package main

import (
	"fmt"
	// "math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2) //1+2i
	var y complex128 = complex(3, 4) //1+2i
	fmt.Println(x * y)
	fmt.Println(real(x * y)) //实部
	fmt.Println(imag(x * y)) //虚部
	// fmt.Println(cmplx.Sqrt(‐1)) // "(0+1i)" 
}
