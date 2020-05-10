package main

import "fmt"

const mod = 998244353

func dfs(n, index, m, remain int) int {
	if index == n {
		return 1
	}

	sum := (dfs(n, index+1, m, remain) % mod * (m - 1)) % mod
	if remain > 0 {
		sum += dfs(n, index+1, m, remain-1) % mod
	}
	return sum
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	ans := m * dfs(n, 1, m, k) % mod
	fmt.Println(ans)
}
