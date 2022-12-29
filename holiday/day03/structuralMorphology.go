package main

import "fmt"

type Person struct {
	name   string
	age    int
	sex    bool
	rmb    float64
	habbit []string
}

func mains1() {

	// p := Person{}
	var p Person = Person{}
	fmt.Printf("p的类型是%T，p的值是%v\n", p, p)
	a := map[string]string{
		"c": "111",
		"d": "222",
	}

	fmt.Println(a)

	p.name = "sss"
	p.age = 20
	p.sex = true
	p.rmb = 500
	p.habbit = []string{"women", "stupid"}
	showPerson(p)
	// showPerson2(&p)
	fmt.Println("函数调用后name", p.name)
	c := make(chan int, 3)
	close(c)
	c <- 1

}

//参数p是PErson的实例——值传递
//值传递，传入的并非结构体本身，而是一个拷贝的副本！！！！
// 对副本的修改不影响主函数中结构体的值
func showPerson(p Person) {
	p.name = "无名氏"
	fmt.Println("showPerson")
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.sex)
	fmt.Println(p.rmb)
	fmt.Println(p.habbit)
}

//参数ptr是一个Person的实例指针——引用传递
//传入的是结构体本身
// 对传入的指针进行修改，会直接改变主函数中结构体本身的值
func showPerson2(ptr *Person) {
	fmt.Println("showPerson2")
	fmt.Println(ptr.age)
	fmt.Println(ptr.sex)
	fmt.Println(ptr.rmb)
	fmt.Println(ptr.habbit)
}
func showPersoninfo() {

}
