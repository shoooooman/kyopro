package main

import "fmt"

func getCount(str string) int {
	flag := make(map[rune]bool)
	cnt := 0
	for _, c := range str {
		if _, ok := flag[c]; !ok {
			cnt++
			flag[c] = true
		}
	}
	return cnt
}

func main() {
	var n int
	fmt.Scan(&n)

	standards := make(map[int][]string, 0)
	standards[1] = append(standards[1], "a")
	for i := 2; i <= n; i++ {
		for _, standard := range standards[i-1] {
			cnt := getCount(standard)
			for j := 0; j <= cnt; j++ {
				c := 'a' + j
				standards[i] = append(standards[i], standard+string(c))
			}
		}
	}
	for _, v := range standards[n] {
		fmt.Println(v)
	}
}
