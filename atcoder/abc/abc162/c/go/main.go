package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcd3(a, b, c int) int {
	d := gcd(a, b)
	return gcd(c, d)
}

func main() {
	var k int
	fmt.Scan(&k)

	sum := 0
	for a := 1; a <= k; a++ {
		for b := 1; b <= k; b++ {
			for c := 1; c <= k; c++ {
				sum += gcd3(a, b, c)
			}
		}
	}
	fmt.Println(sum)
}
