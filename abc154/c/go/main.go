package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	flags := make(map[int]bool)
	ans := false
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if exist := flags[a]; exist {
			ans = true
		}
		flags[a] = true
	}

	if ans {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
