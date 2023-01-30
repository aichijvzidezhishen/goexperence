// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
const (
	AbsoluteZeroC Celsius = ‐273.15 // 绝对零度
	FreezingC Celsius = 0 // 结冰点温度
	BoilingC Celsius = 100 // 沸水温度
) 
func CToF(c Celsius) Fahrenheit {
	 return Fahrenheit(c*9/5 + 32)
	 }
func FToC(f Fahrenheit) Celsius {
	 return Celsius((f - 32) * 5 / 9) 
}