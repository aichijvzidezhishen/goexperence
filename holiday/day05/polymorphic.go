package main

import (
	"fmt"
)

type Biology interface {
	Sayhi()
}

type Man struct {
	name string
	age  int
}

type Monster struct {
	name string
	age  int
}

func (this *Man) Sayhi() { // 实现抽象方法1
	fmt.Printf("Man[%s, %d] sayhi\n", this.name, this.age)
}

func (this *Monster) Sayhi() { // 实现抽象方法1
	fmt.Printf("Monster[%s, %d] sayhi\n", this.name, this.age)
}

func WhoSayHi(i Biology) {
	i.Sayhi()
}

func main() {
	man := &Man{"我是人", 100}
	monster := &Monster{"妖怪", 1000}
	WhoSayHi(man)
	WhoSayHi(monster)
}
