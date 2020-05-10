package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	ans := 0
	if k <= a {
		ans += k
		fmt.Println(ans)
		return
	}
	ans += a
	k -= a
	if k <= b {
		fmt.Println(ans)
		return
	}
	k -= b
	ans -= k
	fmt.Println(ans)
}
