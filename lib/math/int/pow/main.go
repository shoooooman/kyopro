package main

import "fmt"

func pow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		d := pow(a, n/2)
		return d * d
	}
	return a * pow(a, n-1)
}

func main() {
	fmt.Println(pow(2, 10))
}
