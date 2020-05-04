package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	if str[1] == 'B' {
		fmt.Println("ARC")
	} else {
		fmt.Println("ABC")
	}
}
