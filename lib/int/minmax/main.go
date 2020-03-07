package main

import "fmt"

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

func main() {
	x, y := 1, 2
	fmt.Println(min(x, y))
	fmt.Println(max(x, y))
}
