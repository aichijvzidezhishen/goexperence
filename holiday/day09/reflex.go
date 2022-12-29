package main

import (
	"fmt"
	"reflect"
)

/*
	反射：运行时动态的生成一个不知道具体类型是什么的实例或访问无法提前预知的属性或方法
*/

type Animal struct {
	Life int
}
type people struct {
	Name string "姓名"
	Age  int    "年龄"
	Animal
}

func main() {
	p := people{Name: "张三", Age: 20, Animal: Animal{Life: 100}}
	fmt.Println("***", int('4'-'2'))
	typeAPI(p)

}
func typeAPI(obj interface{}) {
	//类型
	oType := reflect.TypeOf(obj)
	fmt.Println(oType) //main.people
	// 原始类型
	oKind := oType.Kind()
	fmt.Println(oKind) //struct
	// 类型名称
	oName := oType.Name()
	fmt.Println(oName) //people
	// 属性的方法和个数
	fmt.Println(oType.NumField())
	fmt.Println(oType.NumMethod())

	fmt.Println("全部属性：")
	for i := 0; i < oType.NumField(); i++ {
		structFiled := oType.Field(i)
		fmt.Println(structFiled.Name, structFiled.Type)
	}
	fmt.Println("全部方法：")
	for i := 0; i < oType.NumMethod(); i++ {
		method := oType.Method(i)
		fmt.Println(method.Name, method.Type)
	}
	fmt.Println("获得父类属性：")
	fmt.Println(oType.FieldByIndex([]int{0, 0}).Name)

}
func valueAPI(p people) {
	oValue := reflect.ValueOf(p)
	fmt.Println("value的原始类型：", oValue.Kind())
	fmt.Println("全部属性的值")
	for i := 0; i < oValue.NumField(); i++ {
		fValue := oValue.Field(i)
		fmt.Println(fValue.Interface())
	}

}
