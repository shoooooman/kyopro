package main

import "fmt"

/* solution1 */
func reverse(str string) string {
	rev := ""
	for i := range str {
		rev += string(str[len(str)-i-1])
	}
	return rev
}

func main() {
	var s string
	fmt.Scan(&s)

	rev := reverse(s)
	for {
		if rev == "" {
			fmt.Println("YES")
			break
		} else if len(rev) >= 7 && rev[:7] == "remaerd" {
			rev = rev[7:]
		} else if len(rev) >= 6 && rev[:6] == "resare" {
			rev = rev[6:]
		} else if len(rev) >= 5 && rev[:5] == "maerd" || rev[:5] == "esare" {
			rev = rev[5:]
		} else {
			fmt.Println("NO")
			break
		}
	}
}

/* solution2 */
// func main() {
// 	var s string
// 	fmt.Scan(&s)
//
// 	divider := []string{"dream", "dreamer", "erase", "eraser"}
// 	dp := [100000]bool{}
// 	dp[0] = true
// 	for i := range s {
// 		if !dp[i] {
// 			continue
// 		}
// 		for _, d := range divider {
// 			l := len(d)
// 			if i+l <= len(s) && s[i:i+l] == d {
// 				dp[i+l] = true
// 			}
// 		}
// 	}
//
// 	if dp[len(s)] {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }
