package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 100
	sum := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				a3 := math.Pow(float64(a), 3.0)
				b3 := math.Pow(float64(b), 3.0)
				c3 := math.Pow(float64(c), 3.0)
				d := int(math.Cbrt(a3 + b3 - c3))
				d3 := math.Pow(float64(d), 3.0)
				if a3+b3 == c3+d3 && d >= 1 && d <= n {
					fmt.Println(a, b, c, d)
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}
