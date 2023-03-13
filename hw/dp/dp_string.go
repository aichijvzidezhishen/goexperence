package main

import (
	"fmt"
)

//插入一个字符
//删除一个字符
//替换一个字符

// 字符串编辑距离
func DpAloLevenshtein(param []string) {
	// var param []string
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	param = append(param, input.Text())
	// }
	for loop := 0; loop < len(param); loop += 2 {
		res := minDistanc(param[loop], param[loop+1])
		fmt.Println(res)
	}
}

func minDistanc(a, b string) int {
	var len1, len2 = len(a), len(b)
	dp := make([][]int, len1+1)
	for loop := range dp {
		dp[loop] = make([]int, len2+1)
	}
	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	return dp[len1][len2]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
