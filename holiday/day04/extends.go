package main

import "fmt"

/*
0.0	继承的目的：扩展和改造父类
*/
type person struct {
	Name string
	age  int //其它包不能直接访问..
	sal  float64
}
type coder struct {
	//持有一个父类声明
	person
}
type driver struct {
	person
	//拓展新的属性
	licence string
}

// 访问子类扩展的方法
func main042() {
	dptr := new(driver)
	dptr.Name = "golang cow"
	dptr.drive()
}

//子类默认拥有父类的全部属性和方法
func main() {
	d := new(driver)
	d.Name = "西门吹雪"
	d.drink()
}

func (d *driver) drink() {
	fmt.Printf("%v喝水", d.Name)
}

//
func (d *driver) drive() {
	fmt.Printf("%v开车", d.Name)
}
