package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a, b, c := s[0], s[1], s[2]

	count := 0
	if a == '1' {
		count++
	}
	if b == '1' {
		count++
	}
	if c == '1' {
		count++
	}

	fmt.Println(count)
}
