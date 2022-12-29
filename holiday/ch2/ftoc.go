package main

import "fmt"

func a() {
	const freezingF, boilingF = 32.0, 212.0
	a, b := 100, 120

	fmt.Printf("%d\n", gcd(a, b))
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

//Ftoc打印两次华氏度到摄氏度的转换。
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

//最大公约数 gcd
func gcd(int, int) int {
	var x, y int
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
