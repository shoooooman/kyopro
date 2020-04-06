package main

import "fmt"

func inRange(data []int, index int) bool {
	return index >= 0 && index <= len(data)-1
}

func elementAt(data []int, index int) int {
	if !inRange(data, index) {
		return -1
	}
	return data[index]
}

func estimateSize(data []int) int {
	size := 1
	for {
		if elementAt(data, size-1) == -1 {
			return size
		}
		size = size << 1
	}
	return -1
}

func search(data []int, find int) int {
	size := estimateSize(data)
	left := 0
	right := size - 1
	for left <= right {
		mid := left + (right-left)/2
		midv := elementAt(data, mid)

		if midv == find {
			return mid
		}

		if midv == -1 {
			right = mid - 1
		} else if find < midv {
			right = mid - 1
		} else if find > midv {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	data := []int{1, 3, 4, 7, 10, 14, 16, 18, 21}
	index := search(data, 21)
	fmt.Println(index)
}
