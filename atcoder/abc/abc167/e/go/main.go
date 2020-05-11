package main

import "fmt"

const (
	mod   = 998244353
	limit = 200001
)

var (
	fac  = make([]int, limit)
	finv = make([]int, limit)
	inv  = make([]int, limit)
)

func powMod(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		d := powMod(a, n/2) % mod
		return d * d % mod
	}
	return a * powMod(a, n-1) % mod
}

func combModInit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < limit; i++ {
		fac[i] = fac[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		finv[i] = finv[i-1] * inv[i] % mod
	}
}

func combMod(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % mod) % mod
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	combModInit()
	ans := 0
	for i := 0; i <= k; i++ {
		ans += (((m * (combMod(n-1, i) % mod)) % mod) * powMod(m-1, n-i-1)) % mod
		ans %= mod
	}
	fmt.Println(ans)
}
