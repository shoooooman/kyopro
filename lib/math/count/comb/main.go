package main

import "fmt"

func perm(n, k int) int {
	ans := 1
	if k > 0 && k <= n {
		for i := 0; i < k; i++ {
			ans *= n - i
		}
	} else if k > n {
		ans = 0
	}
	return ans
}

func fact(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

func fact2(n int) int {
	return perm(n, n-1)
}

func comb(n, k int) int {
	return perm(n, k) / fact(k)
}

func main() {
	n, k := 5, 3
	fmt.Println(fact(n))    // 120
	fmt.Println(fact2(n))   // 120
	fmt.Println(perm(n, k)) // 60
	fmt.Println(comb(n, k)) // 10
}
