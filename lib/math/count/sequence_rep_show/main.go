package main

import (
	"fmt"
)

func perm(pat string, data []rune, ret *[]string) {
	if len(pat) == len(data) {
		// fmt.Println(pat)
		*ret = append(*ret, pat)
		return
	}

	for i := 0; i < len(data); i++ {
		perm(pat+string(data[i]), data, ret)
	}
}

// "abc"の重複を許した順列を全て列挙 (aaa, aab, ..., ccb, ccc)
func main() {
	data := []rune{'a', 'b', 'c'}
	ret := make([]string, 0, len(data))
	perm("", data, &ret)
	fmt.Println(ret)
}
