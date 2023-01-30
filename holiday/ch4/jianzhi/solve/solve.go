package solve

import "fmt"

func Duplicate_1(numbers []int, duplication *[1]int) bool {

	var list = make(map[int]bool)
	for i, v := range list {
		fmt.Println("++++", i, v)
	}
	for i, v := range numbers {
		fmt.Println(i, v)
		if list[v] {
			duplication[0] = v
			fmt.Println(duplication[0])
			return true
		} else {
			list[v] = true
		}
	}
	for i, v := range list {
		fmt.Println(i, v)
	}
	duplication[0] = -1
	// fmt.Println(duplication[0])eeeeee
	return false
	length := len(numbers)
	for _, x := range numbers {
		if x >= length {
			x = x - length
		}
		if numbers[x] >= length {
			duplication[0] = x
			return true
		}
		numbers[x] = numbers[x] + length
	}
	duplication[0] = -1
	return false
}
func Duplicate_2(numbers []int, duplication *[1]int) bool {
	l := len(numbers)
	if l == 0 {
		duplication[0] = -1
		return false
	}
	for i := 0; i < l; i++ {
		if numbers[i] < 0 || numbers[i] > l-1 {
			duplication[0] = -1
			return false
		}
	}
	var hash = make([]int, len(numbers))
	for i := 0; i < l; i++ {
		hash[numbers[i]]++
	}
	for i := 0; i < l; i++ {
		if hash[i] > 1 {
			duplication[0] = i
			return true
		} else {
			duplication[0] = -1
		}
	}

	return false
}
