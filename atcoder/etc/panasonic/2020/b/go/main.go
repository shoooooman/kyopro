package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	ans := 0
	ans += ((w + 1) / 2) * ((h + 1) / 2)
	ans += (w / 2) * (h / 2)
	fmt.Println(ans)
}
