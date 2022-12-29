package main

import "fmt"

/*
	类型断言（type assertion）
	判断一个接口实例的具体类型
*/
type peoples struct {
	name string
	age  int
	no   int
}
type stu struct {
	peoples
	sno int
}

func main61() {
	i := "sda"
	baseType(i)
	a := 33
	baseType(a)
	s := stu{peoples{"kika", 33, 1}, 1}
	structType(s)

	p := peoples{"lisi", 22, 123}
	structType(p)
}

// 基本类型
func baseType(i interface{}) {
	v, ok := i.(string)
	if ok {
		fmt.Println("ok,string is :", v)
	}
	switch i.(type) {
	case string:
		fmt.Println(i, "is string")
	case int:
		fmt.Println(i, "is int")

	}
}

//结构体类型
func structType(i interface{}) {
	v, ok := i.(stu)
	if ok {
		fmt.Printf("SNo is: %v\n", v.sno)
		fmt.Printf("%#v is Student, after convert is: \n%#v\n", i, v)
	}
	v2, ok2 := i.(peoples)
	if ok2 {
		// fmt.Printf("SNo is: %v\n", i.No)
		fmt.Printf("SNo is: %v\n", v2.no)
		fmt.Printf("%#v is People, after convert is: \n%#v\n", i, v2)
	}
}
