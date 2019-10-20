package main

import "fmt"

/* 2-1-1 部分和問題 */
var n, k int
var a []int

func dfs(i int, sum int) bool {
	if i == n {
		return sum == k
	}

	if dfs(i+1, sum) {
		return true
	}
	if dfs(i+1, sum+a[i]) {
		return true
	}

	return false
}

func main() {
	fmt.Scan(&n)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)

	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
