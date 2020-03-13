package main

import (
	"fmt"
	"math"
)

func findMin(nums []int, n int) int {
	min := n
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

// n以下の素数を全て求める
func eratosthenes(n int) []int {
	if n < 2 {
		return nil
	}

	nums := make([]int, n-1)
	for i := 2; i <= n; i++ {
		nums[i-2] = i
	}

	primes := make([]int, 0)
	for true {
		p := findMin(nums, n)

		if p > int(math.Sqrt(float64(n))) {
			break
		}
		primes = append(primes, p)
		i := 0
		for i < len(nums) {
			if nums[i]%p == 0 {
				nums = append(nums[:i], nums[i+1:]...)
				continue
			}
			i++
		}
	}
	for _, v := range nums {
		primes = append(primes, v)
	}
	return primes
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(eratosthenes(n))
}
