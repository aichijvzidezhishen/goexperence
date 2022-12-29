package main

import "fmt"

func main() {
	// var sc = map[string]int{}
	// var  sc map[string]int = map[string]int{}
	// sc := map[string]int{}

	score := make(map[string]int, 5)
	score["a"] = 59
	score["b"] = 32
	score["v"] = 34
	score["d"] = 54
	fmt.Println(score["a"])
	//访问不存在的键
	fmt.Println(score["e"])
	// 待校验的访问值
	grade, ok := score["e"]
	fmt.Println(grade, ok)

}
