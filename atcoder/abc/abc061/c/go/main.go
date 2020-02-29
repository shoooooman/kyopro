package main

import (
	"fmt"
	"strconv"
)

/* 深さ優先探索(非効率) */
/* 最後まで深さ優先探索をして、右側の葉だけsumを返す */
// func dfs(s string, i int, lasti int, sum int) int {
// 	fmt.Println(i, lasti, sum)
// 	if i == len(s) {
// 		if lasti == i {
// 			return sum
// 		}
// 		return 0
// 	}
//
// 	n, _ := strconv.Atoi(s[lasti : i+1])
// 	sum1 := dfs(s, i+1, lasti, sum)
// 	sum2 := dfs(s, i+1, i+1, sum+n)
//
// 	return sum1 + sum2
// }

/* 深さ優先探索 */
/* 最後の項までは探索せず、単純に足す */
func dfs(s string, i int, lasti int, sum int) int {
	fmt.Println(i, lasti, sum)
	if i == len(s)-1 {
		last, _ := strconv.Atoi(s[lasti:len(s)])
		return sum + last
	}

	n, _ := strconv.Atoi(s[lasti : i+1])
	sum1 := dfs(s, i+1, lasti, sum)
	sum2 := dfs(s, i+1, i+1, sum+n)

	return sum1 + sum2
}

/* bit全探索 */
// func solve(s string) int {
// 	sum := 0
// 	l := len(s) - 1
// 	for bit := 0; bit < (1 << uint(l)); bit++ {
// 		lasti := 0
// 		for i := 0; i < l; i++ {
// 			if bit&(1<<uint(i)) != 0 {
// 				n, _ := strconv.Atoi(s[lasti : i+1])
// 				sum += n
// 				lasti = i + 1
// 			}
// 		}
// 		last, _ := strconv.Atoi(s[lasti:])
// 		sum += last
// 	}
// 	return sum
// }

func main() {
	var s string
	fmt.Scan(&s)

	sum := dfs(s, 0, 0, 0)
	// sum := solve(s)
	fmt.Println(sum)
}
