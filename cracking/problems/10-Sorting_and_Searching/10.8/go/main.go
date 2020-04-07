package main

import "fmt"

func solve(data []int) {
	bitVec := 0
	for _, v := range data {
		if bitVec&(1<<uint(v)) != 0 {
			fmt.Println(v)
		} else {
			bitVec |= 1 << uint(v)
		}
	}
}

func main() {
	data := []int{1, 1, 5, 2, 4, 3, 6, 7, 6, 5}
	solve(data)
}
