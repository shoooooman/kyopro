package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	count := 0
L:
	for {
		for i := 0; i < n; i++ {
			if a[i]%2 == 1 {
				break L
			} else {
				a[i] /= 2
			}
		}
		count++
	}
	fmt.Println(count)
}
