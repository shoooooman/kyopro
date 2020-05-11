package main

import "fmt"

// 再帰版
func powModRec(a, n, mod int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		d := powModRec(a, n/2, mod) % mod
		return d * d % mod
	}
	return a * powModRec(a, n-1, mod) % mod
}

// bit演算版
func powMod(a, n, mod int) int {
	rst := 1
	for n > 0 {
		if n&1 != 0 {
			rst = rst * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return rst
}

func main() {
	mod := 1000000007
	fmt.Println(powModRec(3, 45, mod))
	fmt.Println(powMod(3, 45, mod))
}
