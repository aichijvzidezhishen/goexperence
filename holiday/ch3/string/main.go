package main

import (
	"fmt"
	"strconv"
	"time"
)

//strings包提供了字符串的查询、替换、比较、截断、拆分、合并

func main() {

	// //int,uint,rune,byte的使用
	// var b uint = 1
	// var c byte = 255
	// fmt.Println(b, c)
	// fmt.Printf("b的类型是 %T", b)
	// fmt.Println("占用的字节数", unsafe.Sizeof(b))
	// // var n2 int = 10
	// // fmt.Println("字节数", unsafe.Sizeof(n2))
	// var ch byte = 's'
	// fmt.Printf("ch=%c", ch)
	// var c3 int = '呗'
	// fmt.Printf("c3=%c,c3对应的码=%d", c3, c3)
	// var n1 = 20 + 'a'
	// fmt.Println("n1=", n1)
	// str := `反引号，以字符串的原生形式输出，包括换行和特殊字符，可实现防止攻击，输出源代码等效果`
	// fmt.Println(str)
	// //处理字符串中有中文的问题,rune int32的别名，几乎在所有方面等同于int32
	// //它用来区分字符值和整数值
	// str1 := "hello老铁"
	// r := []rune(str1)
	// for i := 0; i < len(r); i++ {
	// 	fmt.Printf("字符=%c\n", r[i])
	// }
	// fmt.Println()
	// // 字符串转整数
	// n, err := strconv.Atoi("12")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(n)
	// }
	// // 整数转字符串
	// str2 := strconv.Itoa(1234)
	// fmt.Printf("str=%v str datatype %T", str2, str2)
	// //字符串转[]byte
	// // var bytes = []byte("golang row!")
	// // []byte转字符串
	// str3 := string([]byte{97, 98, 99})
	// fmt.Printf("str=%v", str3) // 结果abc
	// //10进制转2,8,16进制
	// strconv.FormatInt(num, radix) // radix 进制数
	// // 子串是否在指定的字符串中
	// strings.Contains("seafood", "foo") //true
	// // 统计一个字符串有几个指定的子串
	// strings.Count("ceekee","e") //4
	// //不区分大小写的字符串比较
	// strings.EqualFold("ABC","abc")//true
	// // 返回子串在字符串中第一次出现的index值，如果没有返回-1
	// index ：=strings.Index("nlt_abcdsadabc","abc")//4
	// // 返回子串在字符串中最后一次出现的index值，如果没有返回-1
	// index ：=strings.LastIndex("go golang","go")//3
	// // 将指定的子串替换成另外一个子串
	// // n可以指定替换的个数，-1表示全部替换
	// str4 :="go go hello"
	// strings.Replace(str4,"go","北京",n)
	now := time.Now()
	fmt.Printf("now=%v type=%T", now, now)
	//printf \sprintf
	fmt.Printf("%d-%d-%d %d:%d:%d \n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行时间%v", end-start)

}

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
