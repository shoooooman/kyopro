package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minInts(values ...int) int {
	mv := math.MaxInt64
	for _, v := range values {
		if v < mv {
			mv = v
		}
	}
	return mv
}

func maxInts(values ...int) int {
	mv := math.MinInt64
	for _, v := range values {
		if v > mv {
			mv = v
		}
	}
	return mv
}

func main() {
	x, y := 1, 2
	fmt.Println(min(x, y))
	fmt.Println(max(x, y))
	fmt.Println(minInts(1, 2, 3))
	fmt.Println(maxInts(1, 2, 3))
}
