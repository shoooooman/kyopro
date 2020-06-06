package main

import "fmt"

func main() {
	var a int64
	var b float64
	fmt.Scan(&a, &b)
	fmt.Println((a * int64(100.0*b+0.001)) / 100)
}
