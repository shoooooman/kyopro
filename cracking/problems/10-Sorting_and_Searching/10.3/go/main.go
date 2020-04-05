package main

import "fmt"

func search(data []int, find, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if data[mid] == find {
		return mid
	}

	if data[left] < data[mid] {
		if find >= data[left] && find < data[mid] {
			return search(data, find, left, mid-1)
		}
		return search(data, find, mid+1, right)
	}

	if data[left] > data[mid] {
		if find > data[mid] && find <= data[right] {
			return search(data, find, mid+1, right)
		}
		return search(data, find, left, mid-1)
	}

	if data[left] == data[mid] {
		if data[right] != data[mid] {
			return search(data, find, mid+1, right)
		}
		if rst := search(data, find, left, mid-1); rst != -1 {
			return rst
		}
		if rst := search(data, find, mid+1, right); rst != -1 {
			return rst
		}
	}

	return -1
}

func solve(data []int, find int) int {
	return search(data, find, 0, len(data)-1)
}

func main() {
	data := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
	index := solve(data, 5)
	fmt.Println(index)
}
