package main

import (
	"errors"
	"fmt"
	"math"
)

/*
	自定义异常：
		使用 New 函数创建自定义错误
		使用 Errorf 给错误添加更多信息
		使用结构体类型和字段提供错误的更多信息
		使用结构体类型的方法来提供错误的更多信息
*/
// 用New创建自定义错误
func clcArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed,radius is less than zero")
	}
	return math.Pi * math.Pow(radius, 2), nil
}

// 使用 Errorf 给错误添加更多信息
func clcArea01(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * math.Pow(radius, 2), nil
}

// 使用结构体类型和字段提供错误的更多信息
type areaErr struct {
	err    string
	radius float64
}

func (e *areaErr) Error() string {
	return fmt.Sprintf("radius %0.2f: %s\n", e.radius, e.err)
}
func clcArea02(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaErr{"radius is negative", radius}
	}
	return math.Pi * math.Pow(radius, 2), nil

}

func main085() {

	radius := -20.0
	area, err := clcArea02(radius)
	if err != nil {
		if err, ok := err.(*areaErr); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
		// clcArea01:Area calculation failed, radius -2.00 is less than zero
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}

// errors包中New函数实现
/* func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
} */
