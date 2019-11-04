package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mod := make(map[int]bool)
	for i := 0; ; i++ {
		sum := a * i
		sum %= b
		if sum == c {
			fmt.Println("YES")
			return
		} else if mod[sum] {
			fmt.Println("NO")
			return
		}
		mod[sum] = true
	}
}
