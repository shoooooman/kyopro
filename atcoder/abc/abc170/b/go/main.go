package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if y >= 2*x && (y-2*x)%2 == 0 && 4*x >= y && (4*x-y)%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
