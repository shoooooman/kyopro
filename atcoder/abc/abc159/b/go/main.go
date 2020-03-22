package main

import "fmt"

func kaibun(str string) bool {
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)

	l := len(str)
	if kaibun(str) && kaibun(str[0:(l-1)/2]) && kaibun(str[(l+3)/2-1:]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
