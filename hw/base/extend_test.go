package base

import (
	"fmt"
	"sync"
	"testing"
)

func TestExec(t *testing.T) {
	ta := TaskA{}

	tb := TaskB{}
	Exec(&ta)
	Exec(&tb)
	FuncA()
	T1()
}

func FuncA() {
	a := 1
	f := func() int {
		a += 1
		return a
	}
	fmt.Println("a", f())
}

func T1() {
	s := []string{"a", "b", "c"}
	var wg sync.WaitGroup
	for i, v := range s {
		wg.Add(1)
		go func() {
			fmt.Println(i, v)
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
2 c
2 c
2 c
*/
