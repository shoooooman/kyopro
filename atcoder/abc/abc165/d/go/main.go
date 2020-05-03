package main

import (
	"fmt"
	"math"
)

func solve(a, b, n float64) float64 {
	if b-1 <= n {
		return math.Trunc(a * (b - 1) / b)
	}
	return math.Trunc(a * n / b)
}

func main() {
	var a, b, n float64
	fmt.Scan(&a, &b, &n)

	ans := solve(a, b, n)
	fmt.Println(ans)
}
