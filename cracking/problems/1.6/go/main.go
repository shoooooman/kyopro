package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)

	l := len(str)
	if l < 1 {
		fmt.Println("")
		return
	}

	var buf bytes.Buffer
	prev := rune(str[0])
	num := 1
	for i := 1; i < len(str); i++ {
		c := rune(str[i])
		if prev == c {
			num++
		} else {
			buf.WriteRune(prev)
			buf.WriteString(strconv.Itoa(num))
			num = 1
			prev = c
		}
	}
	buf.WriteRune(prev)
	buf.WriteString(strconv.Itoa(num))

	strAfter := buf.String()
	if l > len(strAfter) {
		fmt.Println(strAfter)
	} else {
		fmt.Println(str)
	}
}
