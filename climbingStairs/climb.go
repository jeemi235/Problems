package main

import "fmt"

func climb(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func main() {
	var n int
	fmt.Printf("enter the value of top: ")
	fmt.Scanln(&n)
	fmt.Println(climb(n))
}
