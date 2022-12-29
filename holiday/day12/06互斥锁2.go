package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	银行账户，存取时必须同步（不能并发操作
	其余操作可以并发
*/
type Account struct {
	money int
	mt    sync.Mutex
}

/*
	存钱
*/
func (a *Account) Save(amount int) int {
	a.mt.Lock()
	fmt.Println("正在存钱----")
	<-time.After(3 * time.Second)
	a.money += amount
	a.mt.Unlock()
	return a.money
}

/*
	取钱
*/
func (a *Account) Get(amount int) int {
	a.mt.Lock()
	fmt.Println("正在取钱----")
	<-time.After(3 * time.Second)
	a.money -= amount
	a.mt.Unlock()
	return a.money
}

/*
	打印上个月的流水
*/
func (a *Account) Print() {
	fmt.Println("打印上个月的流水----")
}

/*
	查询余额
*/
func (a *Account) Query() int {
	fmt.Println("当前余额----", a.money)
	return a.money
}

func main() {
	var wg sync.WaitGroup
	account := new(Account)
	account.Save(100000)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.Save(300)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			account.Get(500)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			account.Print()
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			account.Query()
			wg.Done()
		}()
	}
	wg.Wait()
}
