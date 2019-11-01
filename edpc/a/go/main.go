package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}

	dp := make([]float64, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		cost1 := dp[i-1] + math.Abs(h[i]-h[i-1])
		if i == 1 {
			dp[i] = cost1
			continue
		}
		cost2 := dp[i-2] + math.Abs(h[i]-h[i-2])
		if cost1 < cost2 {
			dp[i] = cost1
		} else {
			dp[i] = cost2
		}
	}
	fmt.Println(int64(dp[n-1]))
}
