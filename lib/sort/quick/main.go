package main

import "fmt"

func med3(x, y, z int) int {
	if x < y {
		if y < z {
			return y
		} else if x < z {
			return z
		} else {
			return x
		}
	} else {
		if x < z {
			return x
		} else if y < z {
			return z
		} else {
			return y
		}
	}
}

func partition(data []int, left, right int) int {
	pivot := med3(data[left], data[(left+right)/2], data[right])
	for left <= right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
	return left
}

func quickSort(data []int, left, right int) {
	index := partition(data, left, right)
	if left < index-1 {
		quickSort(data, left, index-1)
	}
	if index < right {
		quickSort(data, index, right)
	}
}

func main() {
	data := []int{1, 2, -1, 8, 5, 20}
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)
}
