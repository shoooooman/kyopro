package main

import (
	"fmt"
	"strconv"
)

type key struct {
	Head rune
	Tail rune
}

/* solution1 */
// func perm(n, k int) int {
// 	ans := 1
// 	if k > 0 && k <= n {
// 		for i := 0; i < k; i++ {
// 			ans *= n - i
// 		}
// 	} else if k > n {
// 		ans = 0
// 	}
// 	return ans
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)
//
// 	mh := make(map[key][]int)
// 	mt := make(map[key][]int)
// 	ms := make(map[key][]int)
// 	for i := 1; i <= n; i++ {
// 		s := strconv.Itoa(i)
// 		head := rune(s[0]) - '0'
// 		tail := rune(s[len(s)-1]) - '0'
// 		if head == 0 || tail == 0 {
// 			continue
// 		}
// 		if head < tail {
// 			k := key{head, tail}
// 			mh[k] = append(mh[k], i)
// 		} else if head > tail {
// 			k := key{tail, head}
// 			mt[k] = append(mt[k], i)
// 		} else {
// 			k := key{head, tail}
// 			ms[k] = append(ms[k], i)
// 		}
// 	}
//
// 	ans := 0
// 	for _, v := range ms {
// 		ans += perm(len(v), 1)
// 		ans += perm(len(v), 2)
// 	}
// 	for k := range mh {
// 		if _, ok := mt[k]; ok {
// 			ans += len(mh[k]) * len(mt[k]) * 2
// 		}
// 	}
// 	fmt.Println(ans)
// }

/* solution2 */
func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[key][]int)
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		head := rune(s[0]) - '0'
		tail := rune(s[len(s)-1]) - '0'
		if tail == 0 {
			continue
		}
		k := key{head, tail}
		m[k] = append(m[k], i)
	}

	ans := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			v1 := m[key{rune(i), rune(j)}]
			v2 := m[key{rune(j), rune(i)}]
			ans += len(v1) * len(v2)
		}
	}
	fmt.Println(ans)
}
