package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)

	year := 0
	amount := 100.0
	for {
		if amount >= x {
			fmt.Println(year)
			return
		}
		amount = math.Trunc(amount * 1.01)
		year++
	}
}
