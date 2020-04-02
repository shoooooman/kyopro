package main

import "fmt"

func binarySearch(data []int, query int) int {
	left, right := 0, len(data)
	for left <= right {
		mid := left + (right-left)/2
		if data[mid] == query {
			return mid
		}

		if data[mid] > query {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	data := []int{1, 2, 5, 12, 20}
	query := 2
	ret := binarySearch(data, query)
	fmt.Println(ret)
}
