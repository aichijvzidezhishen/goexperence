package main

import (
	"fmt"
	"unicode"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	// reverse_rewrite(&testArr)
	fmt.Println(testArr)

	a := rotate(testArr[:], 2)
	fmt.Println(a)
	str := "В июне в Харбине состоится VI Китайско-российское ЭКСПО"
	str = string([]rune(str)[0:15])
	// 结果为 "你好f"
	fmt.Println(str)
	n := []string{"tao", "taoshihan", "shi", "shi", "han"}
	repeat(n)
	d := []byte("abc bcd wer  sdsd  taoshihan     de")
	fmt.Println(emptyString2(d))
}

/*
4.3 重写reverse函数，使用数组指针代替slice    反转函数
*/
func reverse_rewrite(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}

/*
练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
*/
func rotate(s []int, r int) []int {
	lens := len(s)
	//create a empty slice without element
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}

/*
练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作
*/
func repeat(strings []string) []string {
	i := 0
	index := 0
	num := len(strings)
	for _, v := range strings {
		index = i + 1
		if index >= num {
			break
		}
		if v != strings[index] {
			strings[i] = v
			i++
		}
	}
	fmt.Println(strings[:i])
	return strings[:i]
}

/*
编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
*/

func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			//结合remove函数
			copy(s[i:], s[index:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
