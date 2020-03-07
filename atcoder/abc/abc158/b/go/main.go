package main

import (
	"fmt"
)

// int64で収まるのでfloat64は不要
// func main() {
// 	var n, a, b float64
// 	fmt.Scan(&n, &a, &b)
//
// 	set := math.Trunc(n / (a + b))
// 	blue := set * a
// 	mod := n - set*(a+b)
// 	blue += math.Min(mod, a)
// 	fmt.Println(strconv.FormatFloat(blue, 'f', 0, 64))
// }

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	set := n / (a + b)
	ans := set * a
	mod := n % (a + b)
	// ans += int(math.Min(float64(mod), float64(a))) // なぜか通らない
	ans += min(mod, a)
	fmt.Println(ans)
}
