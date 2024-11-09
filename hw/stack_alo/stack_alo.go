package stackalo

import (
	"strconv"
	"strings"
)

// vaild string
func ValidString(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			//char match end，last char pop
			stack = stack[:len(stack)-1]
		} else {
			//起始符号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// Simplify Path
func SimplifyPath(path string) string {
	stack := []string{}
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			// "."表示当前目录，".."表示上级目录，""表示空字符串，这三种情况都不需要操作
			stack = append(stack, v)

		}

	}
	return "/" + strings.Join(stack, "/")
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err != nil {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		} else {
			stack = append(stack, val)
		}
	}
	return stack[0]
}
