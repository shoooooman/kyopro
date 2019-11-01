package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	ans := 0
	dp := make(map[int]bool)
	dp[0] = true
	for i := 0; i < n; i++ {
		// 降順にループするのは同じp[i]を
		// 2回採用することを避けるため
		// i.e. sum+p[i]が後にsumとして
		// 使用されることを避ける
		for sum := 10000; sum >= 0; sum-- {
			if _, ok := dp[sum]; ok {
				dp[sum+p[i]] = true
			}
		}
	}

	for i := 0; i <= 10000; i++ {
		if _, ok := dp[i]; ok {
			ans++
		}
	}
	fmt.Println(ans)
}
