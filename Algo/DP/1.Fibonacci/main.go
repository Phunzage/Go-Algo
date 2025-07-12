//斐波拉契数列

package main

import "fmt"

func main() {
	fmt.Println(fib(4))
	fmt.Println(fib2(4))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 压缩（优化）状态
func fib2(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
