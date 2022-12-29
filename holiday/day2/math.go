package main

import (
	"math"
)

const pi = 3.14

// func main() {
// 	fmt.Println(2 + 3)
// 	result := math.Pow(2, 3)
// 	fmt.Println("result", result)
// 	result = math.Sqrt(9)
// 	fmt.Println("result", result)
// 	//绝对值
// 	result = math.Abs(-9)
// 	fmt.Println("result绝对值", result)
// 	//纯舍 扔掉小数部分
// 	result = math.Floor(3.5)
// 	fmt.Println("result纯舍", result)
// 	//纯进
// 	result = math.Ceil(3.1)
// 	fmt.Println("result纯进", result)
// 	//正弦
// 	result = math.Sin(30.0 * math.Pi / 180.0)
// 	fmt.Println("resultSin", result)
// 	//余弦
// 	result = math.Cos(30.0 * math.Pi / 180.0)
// 	fmt.Println("result余弦", result)
// 	//正切
// 	result = math.Tan(30.0 * math.Pi / 180.0)
// 	fmt.Println("result正切", result)
// 	a := round(2.6)
// 	fmt.Println("a", a)
// }

//四舍五入
func round(x float64) int {
	return int(math.Floor(x + 0.5))
}
