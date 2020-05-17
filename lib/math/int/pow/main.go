package main

import "fmt"

// 再帰版
func powRec(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		d := powRec(a, n/2)
		return d * d
	}
	return a * powRec(a, n-1)
}

// bit演算版
func pow(a, n int) int {
	rst := 1
	for n > 0 {
		if n&1 != 0 {
			rst *= a
		}
		a *= a
		n >>= 1
	}
	return rst
}

func main() {
	fmt.Println(powRec(2, 10))
	fmt.Println(pow(2, 10))
}
