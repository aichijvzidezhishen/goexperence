package base

import (
	"fmt"
	"log"
	"sync"
	"time"
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
}{
	m1: make(map[string]int),
}

func e() {
	m.RLock()
	read := m.m1["key"]
	m.RUnlock()
	fmt.Println("read", read)
}

type SMap struct {
	rmx sync.RWMutex
	c   map[string]*Entery
}

type Entery struct {
	ch      chan struct{}
	val     interface{}
	isexist bool
}

// 会阻塞 等待key 或者超时
func (m *SMap) Rd(key string, tt time.Duration) interface{} {
	m.rmx.RLock()
	fmt.Println("m.c[key]", m.c[key])
	if e, ok := m.c[key]; ok && e.isexist {
		log.Println("key exist", e.val)
		m.rmx.RUnlock()
		return e.val
	} else if !ok {
		m.rmx.RUnlock()
		m.rmx.Lock()
		e = &Entery{
			ch:      make(chan struct{}),
			isexist: true,
		}
		m.c[key] = e
		m.rmx.Unlock()

		log.Println("协程阻塞 —>", key)
		select {
		case <-e.ch:
			return e.val
		case <-time.After(tt):
			log.Println("协程超时 ->", key)
			return nil
		}
	}
	for {
		time.Sleep(time.Second)
	}
	return nil
}

// mapInternalStruct 是一个函数，用于演示如何通过反射和unsafe包访问Go语言中map的内部结构
// func mapInternalStruct() {
// 	// 创建一个初始容量为100的map，键为string类型，值为int类型
// 	m := make(map[string]int, 100)

// 	// 通过反射获取map的反射值对象
// 	rm := reflect.ValueOf(m)
// 	// 通过unsafe.Pointer将map的地址转换为*hmap指针类型
// 	// hmap是Go语言内部表示map的结构体类型
// 	// ptr := (*hmap)(unsafe.Pointer(rm.UnsafeAddr()))
// }
