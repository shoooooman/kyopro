package main

import "fmt"

/* hash */
func main() {
	var str string
	fmt.Scan(&str)

	hash := make(map[rune]bool)
	for _, c := range str {
		if _, ok := hash[c]; ok {
			fmt.Println("No")
			return
		}
		hash[c] = true
	}
	fmt.Println("Yes")
}

/* bit vector */
// func main() {
// 	var str string
// 	fmt.Scan(&str)
//
// 	checker := 0
// 	for _, c := range str {
// 		val := uint(c - 'a')
// 		if checker&(1<<val) != 0 {
// 			fmt.Println("No")
// 			return
// 		}
// 		checker |= 1 << val
// 	}
// 	fmt.Println("Yes")
// }
