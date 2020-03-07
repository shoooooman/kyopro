package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= 1000; i++ {
		if int(0.08*float64(i)) == a && int(0.1*float64(i)) == b {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("-1")
}
