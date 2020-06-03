package main

import "fmt"

func pow(a, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		tmp := pow(a, n/2)
		return tmp * tmp
	}
	return a * pow(a, n-1)
}

// 計数ソート(安定ソートの例)
func countingSort(data []int, digit, max int) {
	counter := make([]int, max+1)
	div := pow(10, digit)
	for _, v := range data {
		v /= div
		v %= 10
		counter[v]++
	}
	for i := 1; i <= max; i++ {
		counter[i] += counter[i-1]
	}
	ans := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		v := data[i]
		v /= div
		v %= 10
		ans[counter[v]-1] = data[i]
		counter[v]--
	}
	copy(data, ans)
}

// 基数ソート
func radixSort(data []int, d int) {
	for i := 0; i < d; i++ {
		countingSort(data, i, 9)
	}
}

func main() {
	data := []int{329, 457, 657, 839, 436, 720, 355}
	radixSort(data, 3)
	fmt.Println(data)
}
