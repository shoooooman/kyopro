package main

import (
	"fmt"
	"strconv"
)

func fact(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

func exclude(str string, i int) string {
	ret := str[:i]
	ret += str[i+1:]
	return ret
}

func perm(pat, str string, ret *[]string) {
	if len(str) == 0 {
		*ret = append(*ret, pat)
		return
	}

	for i := 0; i < len(str); i++ {
		perm(pat+string(str[i]), exclude(str, i), ret)
	}
}

// "12345"の順列を全て列挙
func main() {
	n := 5
	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(i + 1)
	}

	l := fact(n)
	ret := make([]string, 0, l)
	perm("", str, &ret)
	fmt.Println(ret)
}
