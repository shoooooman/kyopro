package main

import (
	"fmt"
	"os"
)

var dp map[int]int

func solve(n int) int {
	if n < 0 {
		fmt.Println("n must be more than or equal to 0.")
		os.Exit(1)
	}

	if val, ok := dp[n]; ok {
		return val
	}

	s1 := solve(n - 1)
	if n == 1 {
		dp[n] = s1
		return dp[n]
	}
	s2 := solve(n - 2)
	if n == 2 {
		dp[n] = s1 + s2
		return dp[n]
	}
	s3 := solve(n - 3)
	dp[n] = s1 + s2 + s3
	return dp[n]
}

func main() {
	var n int
	fmt.Scan(&n)

	dp = make(map[int]int)
	dp[0] = 1
	ans := solve(n)
	fmt.Println(ans)
}
