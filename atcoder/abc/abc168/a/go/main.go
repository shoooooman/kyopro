package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	r := rune(str[len(str)-1])
	switch r {
	case '2', '4', '5', '7', '9':
		fmt.Println("hon")
	case '0', '1', '6', '8':
		fmt.Println("pon")
	case '3':
		fmt.Println("bon")
	}
}
