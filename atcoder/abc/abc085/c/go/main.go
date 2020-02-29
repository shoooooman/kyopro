package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; j+i <= n; j++ {
			sum := 10000*i + 5000*j + 1000*(n-i-j)
			if sum == y {
				fmt.Printf("%d %d %d\n", i, j, n-i-j)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
