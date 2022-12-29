package main

import (
	"fmt"
	"strings"
)

func mainos() {
	// 获取当前工作路径：默认当前工程根目录
	// dir, err := os.Getwd()
	// fmt.Println(dir, err)

	// env := os.Getenv("path")
	// fmt.Println(env)
	// strings
	fmt.Println(strings.Contains("Hello", "el"))      //true
	fmt.Println(strings.Contains("Hello", "eb"))      //false
	fmt.Println(strings.ContainsAny("Hello", "los"))  //true
	fmt.Println(strings.ContainsRune("he132lo", 'h')) //true
	fmt.Println(strings.Index("Hello", "el"))         //1
	// 0 if a==b, -1 if a < b, and +1 if a > b.
	fmt.Println(strings.Compare("ab", "cd")) //-1
	fmt.Println(strings.Compare("ab", "ab")) //0
	fmt.Println(strings.Compare("ef", "ab")) //1

}
