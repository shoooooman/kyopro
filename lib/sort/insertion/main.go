package main

import "fmt"

func insertionSort(data []int) {
	l := len(data)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if data[j] > data[i] {
				di := data[i]
				// j番目からi-1番目を一つずつ右にずらす
				for k := i; k > j; k-- {
					data[k] = data[k-1]
				}
				data[j] = di
			}
		}
	}
}

func main() {
	data := []int{1, 2, -1, 5, 20, 8}
	insertionSort(data)
	fmt.Println(data)
}
