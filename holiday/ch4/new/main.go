package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("mainres=")
	// for {
	// 	fmt.Println("main下面的代码")
	// 	time.Sleep(time.Second)
	// }

}
func test() {
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("err=", err)
	// 	}
	// }()
	// n1 := 10
	// n2 := 0
	// res := n1 / n2
	// fmt.Println("res=", res)
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test继续执行")
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误..")
	}
}
