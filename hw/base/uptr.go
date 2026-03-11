package base

import (
	"fmt"
	"slices"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

func DiffUintptr() {
	var w *W = new(W)
	// fmt.Println(w.b, w.c)
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	// log.Printf("sss%v %v", w.b, w.c)
	// log.Fatal()
	fmt.Println("code", w.b)
	s := []string{"程序员", "陈明勇"}
	repeat := slices.Repeat(s, 2)

	fmt.Println(w.b, w.c, "repeat", repeat)
	fmt.Println(w.b, w.c)
}
