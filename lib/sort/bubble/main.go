package main

import "fmt"

// バブルソート: 大きい数から順に右端に寄せていく
func bubbleSort(data []int) {
	l := len(data)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func main() {
	data := []int{3, 1, -2, 5, 20, 8}
	bubbleSort(data)
	fmt.Println(data)
}
