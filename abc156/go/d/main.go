package main

import "fmt"

const m = 1000000007

func comb(n, k int) int {
	if k < 1 || k > n {
		return 0
	}
	nu, de := 1, 1
	for i := 0; i < k; i++ {
		nu = nu * (n - i) % m
		de = de * (i + 1) % m
	}
	return nu * repSquare(de, m-2) % m
}

func repSquare(n, p int) int {
	if p == 0 {
		return 1
	}
	if p%2 == 0 {
		t := repSquare(n, p/2)
		return t * t % m
	}
	return n * repSquare(n, p-1) % m
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	all := repSquare(2, n)

	ca := comb(n, a)
	cb := comb(n, b)
	ans := all - 1 - (ca + cb)

	ans = (ans%m + m) % m // negative mod -> positive mod
	fmt.Println(ans)
}
