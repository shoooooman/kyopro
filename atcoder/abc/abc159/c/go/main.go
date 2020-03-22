package main

import (
	"fmt"
	"math"
)

func main() {
	var l int
	fmt.Scan(&l)

	a := float64(l) / 3.0
	fmt.Printf("%.12f\n", math.Pow(a, 3.0))
}
