package main

import "fmt"

func main() {
	var price int
	fmt.Scan(&price)

	pay := 1000
	change := pay - price

	coins := [...]int{500, 100, 50, 10, 5, 1}
	ans := 0
	for _, coin := range coins {
		n := change / coin
		ans += n
		change -= n * coin
	}
	fmt.Println(ans)
}
