package main

import "fmt"

func main() {
	var s, t, u string
	var a, b int
	fmt.Scan(&s, &t, &a, &b, &u)

	if s == u {
		fmt.Println(a-1, b)
	} else if t == u {
		fmt.Println(a, b-1)
	}
}
