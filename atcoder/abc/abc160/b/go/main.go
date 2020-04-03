package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	c500 := x / 500
	x -= c500 * 500
	c5 := x / 5
	fmt.Println(c500*1000 + c5*5)
}
