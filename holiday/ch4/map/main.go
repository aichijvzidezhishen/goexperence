package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//map的key-value 是无序的
	ages := make(map[string]string)
	ages["alcie"] = "sss"
	delete(ages, "alice") //!!
	fmt.Println(ages)
	for name, age := range ages {
		// fmt.Printf("%s\t%d\n", name, age)
		fmt.Printf("%s\t%s\n", name, age)
	}
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
	//读取多行输入，只打印第一行
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
