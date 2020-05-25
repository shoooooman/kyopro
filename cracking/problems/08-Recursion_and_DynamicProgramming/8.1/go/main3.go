package main

import "fmt"

// solution2 (dp, bottom up)
func calcSteps(n int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < n; i++ {
		dp[i+1] += dp[i]
		dp[i+2] += dp[i]
		dp[i+3] += dp[i]
	}
	return dp[n]
}

func main() {
	n := 3
	ans := calcSteps(n)
	fmt.Println(ans)
}
