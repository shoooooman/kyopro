package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		xs[i] = tmp
	}

	const minX = 1
	const maxX = 100
	min := 10000000
	for p := minX; p <= maxX; p++ {
		sum := 0
		for _, x := range xs {
			sum += int(math.Pow(float64(x-p), 2.0))
		}
		if sum < min {
			min = sum
		}
	}
	fmt.Println(min)
}
