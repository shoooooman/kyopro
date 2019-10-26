package main

import "fmt"

/* tを配置できる全ての場合を試して
   最後に辞書順で最も小さいものを取り出す */
func main() {
	var sp, t string
	fmt.Scan(&sp, &t)

	// tを配置できる場所の最初のインデックスを保持
	begin := make([]int, 0)
	for i := 0; i < len(sp); i++ {
		flag := false
		for j := 0; j < len(t) && i+j < len(sp); j++ {
			if !(sp[i+j] == t[j] || sp[i+j] == '?') {
				break
			}
			if j == len(t)-1 {
				flag = true
			}
		}
		if flag {
			begin = append(begin, i)
		}
	}

	if len(begin) == 0 {
		fmt.Println("UNRESTORABLE")
		return
	}

	// tを挿入し他の'?'の部分には'a'を挿れる
	candinates := make([]string, len(begin))
	for j, k := range begin {
		s := ""
		for i := range sp {
			if i >= k && i < k+len(t) {
				s += string(t[i-k])
			} else if sp[i] == '?' {
				s += string('a')
			} else {
				s += string(sp[i])
			}
		}
		candinates[j] = s
	}

	// 辞書順で最も小さいものを取り出す
	var min string
	for i, str := range candinates {
		if i == 0 || str < min {
			min = str
		}
	}
	fmt.Println(min)
}
