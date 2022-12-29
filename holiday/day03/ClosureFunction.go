package main

import (
	"fmt"
	"time"
)

//声明在函数外面的变量=全局变量 = 所有函数都能调用=生命周期与整个程序相同
var heroSlice = []string{"刘", "关", "张"}
var index = 0

func main() {
	for i := 0; i < 10; i++ {
		hero := GetNext()

		fmt.Printf("index%v", index)
		time.Sleep(time.Second)
		fmt.Printf("hero:%s\n", hero)
	}

}

func GetNext() string {
	//函数内变量 = 局部变量，出了函数就没人认识
	// var index = 0
	var hero string
	hero = heroSlice[index]
	index++
	if index == 3 {
		index = 0
	}
	return hero
}
