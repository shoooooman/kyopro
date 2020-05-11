package main

import "fmt"

/* 参考: https://drken1215.hatenablog.com/entry/2018/06/08/210000 */

const (
	mod   = 1000000007
	limit = 500001 // nの最大値
)

var (
	fac  = make([]int, limit)
	finv = make([]int, limit)
	inv  = make([]int, limit)
)

// 前処理: O(n)
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

// クエリ: O(1)
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
	combModInit()
	fmt.Println(combMod(100000, 50000))
}
