package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	l := len(s)
	for i := 0; i < l; i++ {
		fmt.Printf("x")
	}
	fmt.Println()
}
