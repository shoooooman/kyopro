package main

import "fmt"

func solve(a, b []int) {
	l := len(a)
	ai := len(a) - len(b) - 1
	bi := len(b) - 1
	for i := l - 1; bi >= 0; i-- {
		if a[ai] > b[bi] {
			a[i] = a[ai]
			ai--
		} else {
			a[i] = b[bi]
			bi--
		}
	}
}

func main() {
	num := []int{1, 2, 3, 3, 4, 5, 8, 9, 10, 11}
	a := make([]int, 10, 20)
	b := make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		a[i] = num[i*2]
		b[i] = num[i*2+1]
	}
	solve(a, b)
	fmt.Println(a)
}
