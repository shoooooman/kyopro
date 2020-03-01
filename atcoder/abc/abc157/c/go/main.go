package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	digit := make(map[int]int)
	flag := false
	for i := 0; i < m; i++ {
		var s, c int
		fmt.Scan(&s, &c)
		if v, exist := digit[s-1]; exist {
			if v != c {
				flag = true
			}
		} else {
			digit[s-1] = c
		}
	}

	if flag {
		fmt.Println(-1)
		return
	}

	var ans int
	d0, exist := digit[0]
	if !exist {
		if n == 1 {
			ans = 0
		} else {
			ans = 1
		}
	} else if n != 1 && d0 == 0 {
		fmt.Println(-1)
		return
	} else {
		ans = d0
	}

	for i := 1; i < n; i++ {
		ans *= 10
		if v, exist := digit[i]; exist {
			ans += v
		}
	}
	fmt.Println(ans)
}
