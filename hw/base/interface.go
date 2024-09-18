package base

import "fmt"

// 定义一个基础接口
type BasicInterface interface {
	BasicMethod()
}

// 定义一个扩展接口，嵌套了基础接口
type ExtendedInterface interface {
	BasicInterface // 嵌套基础接口
	ExtendedMethod()
}

// 实现基础接口
type MyStruct struct{}

func (m MyStruct) BasicMethod() {
	fmt.Println("Basic method")
}

// 实现扩展接口
func (m MyStruct) ExtendedMethod() {
	fmt.Println("Extended method")
}

func IfaceExtend() {
	// 创建一个实现了扩展接口的对象
	obj := MyStruct{}

	// 调用基础接口的方法
	obj.BasicMethod()

	// 调用扩展接口的方法
	obj.ExtendedMethod()
}
