package main

import "fmt"

// 計数ソート
func countingSort(data []int, max int) {
	counter := make([]int, max+1)
	for _, v := range data {
		counter[v]++
	}
	// i以下の要素の数を計算(累積和)
	for i := 1; i <= max; i++ {
		counter[i] += counter[i-1]
	}
	ans := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		v := data[i]
		ans[counter[v]-1] = v
		// 同じ値がある場合、右から埋めていく
		counter[v]--
	}
	copy(data, ans)
}

func main() {
	data := []int{2, 5, 3, 0, 2, 3, 0, 3}
	countingSort(data, 5)
	fmt.Println(data)
}
