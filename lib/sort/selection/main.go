package main

import "fmt"

func selectionSort(data []int) {
	l := len(data)
	for i := 0; i < l-1; i++ {
		minj := i
		for j := i + 1; j < l; j++ {
			if data[j] < data[minj] {
				minj = j
			}
		}
		data[i], data[minj] = data[minj], data[i]
	}
}

func main() {
	data := []int{1, 3, -2, 5, 20, 8}
	selectionSort(data)
	fmt.Println(data)
}
