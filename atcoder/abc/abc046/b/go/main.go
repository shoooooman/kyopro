package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)

	ans := k * math.Pow(k-1, n-1)
	fmt.Println(int64(ans))
}
