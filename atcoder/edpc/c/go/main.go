package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	v := make([][]int, n)
	for i := 0; i < n; i++ {
		v[i] = make([]int, 3)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i][0], &v[i][1], &v[i][2])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 3)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				next := dp[i][j] + v[i][k]
				if next > dp[i+1][k] {
					dp[i+1][k] = next
				}
			}
		}
	}

	ans := -1
	for k := 0; k < 3; k++ {
		if dp[n][k] > ans {
			ans = dp[n][k]
		}
	}
	fmt.Println(ans)
}
