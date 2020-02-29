package main

import "fmt"

func main() {
	var x, y int64
	fmt.Scan(&x, &y)

	ans := 0
	for a := x; a <= y; a *= 2 {
		ans++
	}
	fmt.Println(ans)
}
