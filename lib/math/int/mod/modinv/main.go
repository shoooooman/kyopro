package main

import "fmt"

// 拡張Euclidの互除法
// aの逆元(a^(-1) mod m)を求める
// 必要条件: aとmが互いに素
func modinv(a, m int) int {
	b, u, v := m, 1, 0
	for b != 0 {
		q := a / b
		a, b = b, a-q*b
		u, v = v, u-q*v
	}
	u %= m
	if u < 0 {
		u += m
	}
	return u
}

func main() {
	for i := 0; i < 13; i++ {
		fmt.Println(i, modinv(i, 13))
	}
}
