package main

import "fmt"

func main() {
	const mod = 1000000007

	var n int
	fmt.Scan(&n)

	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
		ans %= mod
	}
	fmt.Println(ans)
}
