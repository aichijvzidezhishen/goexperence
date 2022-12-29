package main

import "fmt"

const (
	sunday = iota * 2
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
	// [Running] go run "e:\workplace\Go\src\holiday\day1\iota定义常量组.go"
	// 0 1 2 3 4 5 6
}
