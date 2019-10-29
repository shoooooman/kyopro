package main

import (
	"fmt"
	"math"
)

/* solution1 */
// func main() {
// 	var a, b, c, d int
// 	fmt.Scan(&a, &b, &c, &d)
//
// 	ans := 0
// 	for i := 0; i <= 100; i++ {
// 		if i >= a && i < b && i >= c && i < d {
// 			ans++
// 		}
// 	}
// 	fmt.Println(ans)
// }

/* solution2 */
func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

	ans := math.Min(b, d) - math.Max(a, c)
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
}
