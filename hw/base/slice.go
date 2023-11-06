package base

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// ex1
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

// ex1
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

type Event struct {
	Eventid int32
	Idlist  []string
}

type Events struct {
	List []Event
}

func (es *Events) Refresh() {
	for _, v := range es.List {
		v.RefreshEvent()
	}
}

func (e *Event) RefreshEvent() {
	PrintSliceStruct(&e.Idlist, "RefreshEvent")
	for k := range e.Idlist {
		e.Idlist[k] = "test" + strconv.Itoa(int(e.Eventid))
	}
}

func PrintSliceStruct(s1 *[]string, from string) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(s1))
	fmt.Printf("from %v slice addr %+v\n", from, sh)
}

func DeepCopy(dist []int32) (s1 []int32) {
	s1 = []int32{3, 21, 1, 4}
	copyNum := copy(s1, dist)
	fmt.Println("copyNum", copyNum)
	return s1
}
