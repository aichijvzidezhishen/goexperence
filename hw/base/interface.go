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
	// 创建一个实现了扩展接口的对象}
	obj := MyStruct{}
	// ExtendedInterface{}
	// 调用基础接口的方法
	obj.BasicMethod()

	// 调用扩展接口的方法
	obj.ExtendedMethod()
}

// 类型断言
type Stu struct {
	Name string
	Age  int
}

func Judge(v interface{}) {
	fmt.Printf("%p %v\n", &v, v)

	switch v := v.(type) {
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)

	case Stu:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Stu type[%T] %v\n", v, v)

	case *Stu:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Stu type[%T] %v\n", v, v)

	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unknow\n")
	}
}
