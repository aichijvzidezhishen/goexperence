package arrayz

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//Fibonacci
//二维数组查找值

var (
	stack1 []int
	stack2 []int
)

// func main() {
// 	s := "dsada"
// 	strings.Replace(s, "a", "b", -1)
// 	a := Fibonacci(5)
// 	fmt.Println(a)
// 	result1 := stringToIntb("132")
// 	fmt.Println("result1", result1)

// }
func push(node int) {
	stack1 = append(stack1, node)
}
func pop() int {
	if len(stack2) == 0 {
		n := len(stack1)
		for i := 0; i < n; i++ {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	n := len(stack2)
	if n == 0 {
		return -1
	}
	ans := stack2[0]
	stack2 = stack2[1:]
	return ans
}

// 斐波那契
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	i := 2
	pre, cur := 1, 1
	for i < n {
		next := cur + pre
		pre = cur
		cur = next
		i++
	}

	return cur
}

// 二维数组查找值
func Find(target int, array [][]int) bool {
	// write code here
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 二维数组查找值 参数有数组是 注意考虑数组长度为0
func Find_a(target int, array [][]int) bool {
	// write code here
	l := len(array[0])
	if l == 0 {
		return false
	}
	for i, _ := range array {
		if array[i][l-1] < target {
			continue
		}
		if array[i][0] > target {
			continue
		}
		for j := range array[i] {
			if array[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
func printListFromTailToHead(head *ListNode) []int {
	// write code here
	vals := []int{}
	for ; head != nil; head = head.Next {

		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i := range vals[:n/2] {
		tem := vals[i]
		vals[i] = vals[n-i-1]
		vals[n-i-1] = tem

	}
	return vals
}

// 递归求解
func printListFromTailToHead_recursion(head *ListNode) []int {
	// write code here
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	ans := []int{}
	i := printListFromTailToHead(head.Next)
	ans = append(ans, i...)
	ans = append(ans, head.Val)
	return ans
}
func printListFromTailToHead_optimization(head *ListNode) []int {
	// write code here
	if head == nil {
		return []int{}
	}
	return append(printListFromTailToHead(head.Next), head.Val)
}
func ReverseList(phead *ListNode) *ListNode {
	return phead
}

// 是否为丑数
func isUgly(num int) bool {
	if num == 0 {
		return false
	} else if num == 1 {
		return true
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	}
	if num%3 == 0 {
		return isUgly(num / 3)
	}
	if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}

// 非递归丑数
func isUglyA(num int) bool {
	if num == 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%5 == 0 {
			num /= 5
		} else if num%3 == 0 {
			num /= 3
		} else {
			return false
		}
	}
	return true

}

// 把字符串转成整数
func stringToInta(s string) int {
	p := 0
	num := 0
	flag := false
	if len(s) == 0 {
		return 0
	}
	if s[0] == '+' {
		p = 1
	} else if s[0] == '-' {
		p = 1
		flag = true
	}
	for p < len(s) {
		if s[p] > '9' || s[p] < '0' {
			return 0
		}
		num *= 10
		//s中存放的是数字字符，减去‘0’，将其转成数值类型
		num += int(s[p] - '0')
		fmt.Println("num:", num)
		//p控制数值的位数
		p++
	}

	if flag == true {
		num = -num
	}
	return num
}
func stringToIntb(s string) int {
	flag := false
	val := 0
	for _, ch := range []byte(s) {
		fmt.Println("ch:", ch)
		switch ch {
		case '-':
			flag = true
		case '+':
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			val = val*10 + int(ch-'0')
		default:
			return 0
			w
		}
	}
	if flag {
		return -val
	}
	return val
}
