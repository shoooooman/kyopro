package main

import (
	"fmt"
	"math"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)

	if s1 == s2 {
		fmt.Println("True")
		return
	}

	l1 := len(s1)
	l2 := len(s2)
	if math.Abs(float64(l1-l2)) > 1.0 {
		fmt.Println("False")
		return
	}
	// 長い方をs1としておく(挿入と削除を一緒に扱うため)
	if l1 < l2 {
		tmp := s1
		s1 = s2
		s2 = tmp
	}

	i := 0
	j := 0
	flag := false
	for {
		if j == len(s2) {
			// pales, paleの場合、最後を削除する必要がある
			if i != len(s1) && flag {
				fmt.Println("False")
			} else {
				fmt.Println("True")
			}
			return
		}

		c1 := rune(s1[i])
		c2 := rune(s2[j])
		if c1 == c2 {
			i++
			j++
		} else {
			// 2回目以降の操作は禁止
			if flag {
				fmt.Println("False")
				return
			}
			if l1 == l2 {
				i++
				j++
				flag = true
			} else {
				// 長い方の文字列だけ進める
				i++
				flag = true
			}
		}
	}
}
