package base

import (
	"fmt"
	"sync"
	"time"
)

func MutexEx() {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	fmt.Println("Locking  (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	wg.Wait()
}

// func (*RWMutex) Lock // 写锁定
// func (*RWMutex) Unlock // 写解锁
// func (*RWMutex) RLock // 读锁定
// func (*RWMutex) RUnlock // 读解锁
func RWMutexEx() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	var rwMu sync.RWMutex
	data := 0
	for i := 0; i < 10; i++ {
		go func(t int) {
			rwMu.RLock()
			defer rwMu.RUnlock()
			fmt.Println("read data", data)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)
		go func(t int) {
			rwMu.Lock()
			defer rwMu.Unlock()
			data += t
			fmt.Println("write data: ", data, "t", t)
			wg.Done()
			time.Sleep(2 * time.Second)
		}(i)
	}
	time.Sleep(5 * time.Second)
	wg.Wait()

}
