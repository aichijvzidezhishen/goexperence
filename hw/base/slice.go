package base

import "fmt"

//ex1
func Entrance1() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse1(s)
	fmt.Println("s", s)
}

func reverse1(s []int) {
	// s = append(s, 999)
	fmt.Println("s", s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

//ex1
func Entrance2() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse2(s)
	fmt.Println("s", s)
}

func reverse2(s []int) {
	s = append(s, 999)
	fmt.Println("s", s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Entrance3() {
	var s []int
	for i := 1; i <= 3; i++ {
		s = append(s, i)
	}
	reverse3(s)
	fmt.Println("s", s)
}

func reverse3(s []int) {
	s = append(s, 999, 10, 11)
	fmt.Println("s", s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}
