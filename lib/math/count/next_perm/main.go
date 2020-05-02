package main

import "fmt"

// i以降(i含む)のaの要素を反転させる
func reverse(a []int, i int) {
	l := len(a)
	for j := i; j <= (l+i-1)/2; j++ {
		a[j], a[l-j+i-1] = a[l-j+i-1], a[j]
	}
}

func nextPermutation(a []int) []int {
	l := len(a)
	index := 0
	for i := l - 1; i > 0; i-- {
		if a[i] > a[i-1] {
			index = i
			break
		}
	}
	if index > 0 {
		// index以降の要素の内、a[index-1]よりも大きい最小の要素
		toSwap := index
		for j := index + 1; j < l; j++ {
			if a[j] <= a[index-1] {
				break
			}
			toSwap = j
		}
		a[index-1], a[toSwap] = a[toSwap], a[index-1]
	}
	reverse(a, index)
	return a
}

func fact(n int) int {
	if n <= 0 {
		return 0
	}
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	return ans
}

func main() {
	a := []int{1, 2, 3}
	for i := 0; i < fact(len(a)); i++ {
		next := nextPermutation(a)
		fmt.Println(next)
	}
}
