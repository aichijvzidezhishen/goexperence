package main

import (
	"fmt"
	"time"
)

var timestr = "2006-01-02 15:04:05"

type Slice[T int | float32 | float64] []T
type Test1 struct {
	Age  int
	Name string
}
type Test2 struct {
	A []Test1
}

func GetTe1(arr []*Test1) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr", arr[i].Age)
	}
}

func main() {
	// var te1 *Test1
	te1 := []*Test1{}
	te1 = append(te1)
	GetTe1(te1)
	// Test()
	// fmt.Println("te1", GetTe1(te1))
	// fmt.Println("ss")
	// var a Slice[int] = []int{1, 2, 3}
	// fmt.Printf("Type Name: %T", a) //输出：Type Name: Slice[int]
	// var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	// fmt.Printf("Type Name: %T", b) //输出：Type Name: Slice[float32]
	// var aa int32
	// for i := 1; i <= 500; i += 10 {
	// 	fmt.Println("start", i, "end", i+9)
	// 	aa++

	// }
	// fmt.Println("a1", aa)
	// // fmt.Println("a1", a1)
	// c := 45
	// a := c / 10
	// b := c % 10
	// if b == 0 {
	// 	fmt.Println("gr", a)
	// } else {
	// 	fmt.Println("gr", a+1)
	// }
	// fmt.Println("", 45/10)
	// fmt.Println("a", a)
	// fmt.Println(IsSameDay(0, time.Now().Unix()))
	fmt.Println(IsAfterDay(time.Now(), GetTimeByUnixTimeStamp(0)))
	// str := "2022-08-01 4:00:00"
	// t, err := GetTimeByFmtString(str)
	// if err != nil {
	// 	fmt.Println("err", err.Error())
	// }
	// var b1 = Test1{}
	// fmt.Println("b1", b1)

	// var a1 Test2
	// a1.A = append(a1.A, Test1{1, "2"}, Test1{2, "3"})
	// fmt.Println("a1", a1)
	// a2 := a1
	// fmt.Println("a2", a2)
	// // a := Test2{{Test1{1, "ss"}}, {Test1}}
	// fmt.Println("t", t)
	// m := map[int32]dd{}
	// m[1] = dd{1}
	// m[2] = dd{2}
	// m[3] = dd{3}
	// for _, v := range m {
	// 	fmt.Println("ssss", v)
	// }
}

func GetTimeByFmtString(fs string) (time.Time, error) {
	return time.ParseInLocation(timestr, fs, time.Local)
}

// IsSameDay 是否是同一天
func IsSameDay(t1, t2 int64) bool {
	gt1 := time.Unix(t1, 0)
	gt2 := time.Unix(t2, 0)
	return gt1.Year() == gt2.Year() && gt1.Month() == gt2.Month() && gt1.Day() == gt2.Day()

}
func IsAfterDay(t1, t2 time.Time) bool {
	fmt.Println("t1.YearDay()", t1.YearDay(), t1.Hour())

	fmt.Println("t2.YearDay()", t2.YearDay(), t2.Hour())
	if t1.YearDay() != t2.YearDay() && t2.Hour() >= 4 {
		return true
	}

	return false
}

// GetTimeByUnixTimeStamp 通过unix时间获得time
func GetTimeByUnixTimeStamp(t int64) time.Time {
	return time.Unix(t, 0)
}

type dd struct {
	A int
}
