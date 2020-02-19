package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* hash */
// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	str := scanner.Text()
//
// 	lstr := strings.ToLower(str)
//
// 	counter := make(map[rune]int)
// 	for _, c := range lstr {
// 		if c != ' ' {
// 			counter[c]++
// 		}
// 	}
//
// 	odd := 0
// 	for _, cnt := range counter {
// 		if cnt%2 == 1 {
// 			odd++
// 			if odd > 1 {
// 				fmt.Println("False")
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println("True")
// }

/* bit vector */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	lstr := strings.ToLower(str)

	checker := 0
	for _, c := range lstr {
		if c != ' ' {
			// 排他的論理和で偶数個か奇数個かをチェック
			checker ^= (1 << uint(c-'a'))
		}
	}

	// ASCII
	// const n = 128
	// odd := 0
	// for i := 0; i < n; i++ {
	// 	// 小さい桁から順に見ていく
	// 	if (checker>>uint(i))&1 == 1 {
	// 		odd++
	// 		if odd > 1 {
	// 			fmt.Println("False")
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println("True")

	/* Tips: 1引いたものと論理和を取れば1が一個だけか分かる */
	if checker&(checker-1) == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
