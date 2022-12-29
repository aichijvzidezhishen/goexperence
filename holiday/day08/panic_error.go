package main

import "fmt"

/*
	系统级异常：
	给的地址无用；
	找不到文件；
	赋值给空映射中的条目
	map没有初始化就复制
	非法内存地址或者空指针
	下标越界
*/
// 系统报恐慌
func main081() {
	/* 	指针的默认值为nil
	   	var a *int
	   	fmt.Println(*a)
	   	var slice = []int{3, 1, 4}
	   	fmt.Println(slice[5])
	   	给空map写入键值对
	   	var maps map[string]interface{}
	   	maps["name"] = "xiaowang" */

}

//自己报恐慌，直接杀死程序
func main082() {
	jiehun(&human{name: "ximen"}, &pig{name: "小潘"})

}

type human struct {
	name   string
	yanzhi int
}
type pig struct {
	name   string
	yanzhi int
}

//自定义错误
type lowlookErr struct{}

// 让低颜值错误实现系统的error接口
func (l *lowlookErr) Error() string {
	return "颜值过低"
}

func jiehun(laogong *human, laopo interface{}) (ok bool, err error) {
	if h, ok := laopo.(*human); !ok {
		panic("不配得到爱情！")
	} else {
		//颜值过低，不能生活，返回【生活失败信息】
		if h.yanzhi < 60 {
			// 造一个实现了error接口的【低颜值错误】实例，并返回
			err := &lowlookErr{}
			return false, err
		}
	}

	fmt.Printf("%v和%v开始审核", laogong, laopo)
	return true, nil
}

func main083() {
	ok, err := jiehun(&human{name: "arong", yanzhi: 80}, &human{name: "sexy", yanzhi: 59})
	fmt.Println(ok, err)
}
