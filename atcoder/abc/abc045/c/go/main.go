package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	// e.g. 123
	// 1 2 3
	//  ↑ ↑    ← '+'を入れるところ
	n := len(s) - 1
	sum := 0
	// 00, 01, 10, 11が上の矢印の有無を表す
	for bit := 0; bit < (1 << uint(n)); bit++ {
		// 最後に'+'が入ったところ
		last := 0
		for i := 0; i < n; i++ {
			// bit内で立っているビットを列挙していく
			// bit=0b10ならi=1のとき!=0となる
			// つまり2と3の間に'+'が入り，12+3となる
			if bit&(1<<uint(i)) != 0 {
				term, _ := strconv.Atoi(s[last : i+1])
				sum += term
				last = i + 1
			}
		}
		// 最後の'+'以降の項を加える
		term, _ := strconv.Atoi(s[last:])
		sum += term
	}
	fmt.Println(sum)
}
