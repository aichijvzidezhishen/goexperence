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
	return nil
}
