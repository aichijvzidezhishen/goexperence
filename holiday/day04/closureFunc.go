package main

import "fmt"

/*
闭包函数:返回函数的函数
闭包就像一层壳
内核各有各的状态，都在自己的壳里
调一次闭包就得到一个内核
*/
var heros = []string{"刘", "关", "张"}

// var index = 0

func mainclose() {
	f := NextHero()
	for i := 0; i < 10; i++ {
		fmt.Printf("hero%s\n", f())
	}

}

func NextHero() func() string {
	var index int
	f := func() string {
		hero := heros[index]
		index++
		if index > 2 {
			index = 0
		}
		return hero
	}
	return f
}

// 指针方案
func getNext(indexptr *int) {

}
