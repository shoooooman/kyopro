package main

import "fmt"

var dp = make(map[int]int)

func calcSteps(n int) int {
	if n < 0 {
		return 0
	}

	if v, ok := dp[n]; ok {
		return v
	}

	s1 := calcSteps(n - 1)
	s2 := calcSteps(n - 2)
	s3 := calcSteps(n - 3)
	dp[n] = s1 + s2 + s3
	return dp[n]
}

func main() {
	dp[0] = 1
	n := 3
	ans := calcSteps(n)
	fmt.Println(ans)
}
