package hash

import (
	"fmt"
	"sort"
	"strconv"
)

// 数据表记录包含表索引index和数值value（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照index值升序进行输
func MergeTableRecord() {
	hash := make(map[int]int)
	var key, value, num int

	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Scan(&key, &value)
		hash[key] += value
	}
	//
	nums := make([][]int, len(hash))
	numIndex := 0
	for k, v := range hash {
		nums[numIndex] = []int{k, v}
		numIndex++
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	//输出
	for _, v := range nums {
		fmt.Printf("%d %d \n", v[0], v[1])
	}
}

func NorepeatNum(num int) {
	var (
		// num int
		str    string
		record = make(map[string]bool)
		out    string
	)

	// fmt.Scan(&num)
	str = strconv.Itoa(num)
	for i := len(str) - 1; i >= 0; i-- {

		if record[string(str[i])] {
			continue
		} else {
			record[string(str[i])] = true
		}
		out += string(str[i])
	}
	fmt.Println(out)

}

func CharNum(n string) {
	var (
		record = make(map[rune]bool)
		count  int
	)

	for _, v := range n {
		if !record[v] {
			record[v] = true
			count++
		}
	}
	fmt.Println("count", count)
}

func ReverseNum(str string) {
	// var str string
	// fmt.Scan(&str)
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%s", string(str[i]))
	}
}
