package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q", data)
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("elem of slice is", slice, "len of slice", len(slice))
	fmt.Println("slice的容量", cap(slice)) //4
	data1 := make([]int, 10)
	data1 = append(data1, 1, 2, 3)
	for _, v := range data1 {
		fmt.Println(v)
	}
	// arr := []int{3, 4, 5, 1, 2}
	// q = len(arr)/2

	// for i := 0; i < len(arr); i++ {
	// 	arr[i]=arr                                                                                                                                                                                                                                                                                                                                                                                                  []
	// }
	Sum([]int{1, -1, 56, 9})

}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
func nonempty_2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
func Sum(intArr []int) int {
	sum := 0
	for val := range intArr {
		sum += val
	}
	return sum
}
