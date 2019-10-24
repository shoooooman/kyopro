package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	s := make([]string, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&s[i])
	}

	for i := 0; i < t; i++ {
		count := 0
		for j := 0; j < len(s[i])-4; j++ {
			if s[i][j:j+5] == "tokyo" || s[i][j:j+5] == "kyoto" {
				count++
				j += 4
			}
		}
		fmt.Println(count)
	}
}
