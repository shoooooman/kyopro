package main

import (
	"fmt"
)

/* solution 1 */
// func main() {
// 	var n, a, b int
// 	fmt.Scan(&n, &a, &b)
//
// 	ans := 0
// 	for i := 0; i <= n; i++ {
// 		strI := strconv.Itoa(i)
// 		digitSum := 0
// 		for j := 0; j < len(strI); j++ {
// 			num, _ := strconv.Atoi(string(strI[j]))
// 			digitSum += num
// 		}
// 		if digitSum >= a && digitSum <= b {
// 			ans += i
// 		}
// 	}
// 	fmt.Println(ans)
// }

/* solution 2 */
func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := 0
	for i := 0; i <= n; i++ {
		num := i
		sum := 0
		for num > 0 {
			mod := num % 10
			sum += mod
			num /= 10
		}
		if sum >= a && sum <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
