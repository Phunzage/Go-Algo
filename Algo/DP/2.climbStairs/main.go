package main

import "fmt"

func main() {
	var n int
	// 求第n个数对应的值
	fmt.Scan(&n)
	if n < 2 {
		fmt.Print(n)
		return
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	fmt.Print(dp[n])
}
