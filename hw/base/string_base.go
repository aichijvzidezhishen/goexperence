package base

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
	ascii
*/

func StringBase() {
	var str string = "Hello,我是 World!"
	fmt.Println("len(str) =", len(str))
	//有两种方式可以获取字符串长度
	fmt.Println("RuneCount(str) : ", utf8.RuneCountInString(str))
	//通过rune类型获取unicode字符串长度
	fmt.Println("len([]byte(str)) : ", len([]rune(str)))
	//
}

// 全角字符转半角
func FullWidthToHalfWidth(r rune) rune {
	if r == 12288 {
		return 32
	}

	if r >= 65281 && r <= 65374 {
		return r - 65248
	}

	return r
}

func IsChineseChar(r rune) bool {
	// Unicode 中文范围 U+4E00 ～ U+9FFF
	return unicode.Is(unicode.Han, r)
}
