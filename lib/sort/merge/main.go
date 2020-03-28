package main

import "fmt"

func merge(l, r []int) []int {
	rst := make([]int, len(l)+len(r))
	lj, rj := 0, 0
	for i := 0; i < len(rst); i++ {
		if lj > len(l)-1 {
			rst[i] = r[rj]
			rj++
		} else if rj > len(r)-1 {
			rst[i] = l[lj]
			lj++
		} else if l[lj] <= r[rj] {
			rst[i] = l[lj]
			lj++
		} else {
			rst[i] = r[rj]
			rj++
		}
	}
	return rst
}

func mergeSort(data []int, left, right int) []int {
	if len(data) == 0 {
		return nil
	}

	if left == right {
		return []int{data[left]}
	}
	mid := (left + right) / 2
	l := mergeSort(data, left, mid)
	r := mergeSort(data, mid+1, right)
	return merge(l, r)
}

func main() {
	data := []int{1, 2, -1, 8, 5, 20}
	data = mergeSort(data, 0, len(data)-1)
	fmt.Println(data)
}
