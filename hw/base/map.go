package base

import (
	"fmt"
	"sync"
)

func SyncMap() {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

var m = struct {
	sync.RWMutex
	m1 map[string]int
}{m1: make(map[string]int)}

func e() {
	m.RLock()
	read := m.m1["key"]
	m.RUnlock()
	fmt.Println("")
}
