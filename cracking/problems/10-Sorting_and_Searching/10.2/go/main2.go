package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func solve(strs []string) {
	// sortした文字列をキーとしてmapに追加
	h := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		if _, ok := h[key]; !ok {
			h[key] = make([]string, 0)
		}
		h[key] = append(h[key], str)
	}

	// mapからkeyが同じ物順に取り出す
	index := 0
	for _, ss := range h {
		for _, s := range ss {
			strs[index] = s
			index++
		}
	}
}

func main() {
	strs := []string{"abc", "abb", "acb", "acc", "bab"}
	solve(strs)
	fmt.Println(strs)
}
