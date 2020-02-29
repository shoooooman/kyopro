package main

import (
	"fmt"
	"sort"
	"strings"
)

/* sort: O(nlogn) */
func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if len(s) != len(t) {
		fmt.Println("Np")
		return
	}

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(tt)
	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

/* counting: O(n) */
// func main() {
// 	var s, t string
// 	fmt.Scan(&s, &t)
//
// 	if len(s) != len(t) {
// 		fmt.Println("No")
// 		return
// 	}
//
// 	scounter := make(map[rune]int)
// 	for _, c := range s {
// 		scounter[c]++
// 	}
//
// 	tcounter := make(map[rune]int)
// 	for _, c := range t {
// 		tcounter[c]++
// 	}
//
// 	for key, cnt := range scounter {
// 		if tcounter[key] != cnt {
// 			fmt.Println("No")
// 			return
// 		}
// 	}
// 	fmt.Println("Yes")
// }
