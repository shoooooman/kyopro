package main

import "fmt"

func binary(data []string, left, right int, find string) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if data[mid] == find {
		return mid
	}
	if data[mid] == "" {
		if index := binary(data, left, mid-1, find); index != -1 {
			return binary(data, left, mid-1, find)
		}
		return binary(data, mid+1, right, find)
	}
	if data[mid] < find {
		return binary(data, left, mid-1, find)
	}
	if data[mid] > find {
		return binary(data, mid+1, right, find)
	}
	return -1
}

func search(data []string, find string) int {
	return binary(data, 0, len(data)-1, find)
}

func main() {
	data := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	index := search(data, "ball")
	fmt.Println(index)
}
