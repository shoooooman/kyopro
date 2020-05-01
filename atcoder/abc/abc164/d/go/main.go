package main

import (
	"fmt"
	"strconv"
)

func solve(str string) int {
	const mod = 2019
	ans := 0
	validEnd := make(map[int]int)
	for i := 0; i < len(str)-3; i++ {
		sizes := []int{3, 4}
		for _, size := range sizes {
			j := i + size
			if j > len(str)-1 {
				continue
			}
			num, _ := strconv.Atoi(str[i : j+1])
			if num%mod == 0 {
				validEnd[j]++
				ans++
				if v, ok := validEnd[i-1]; ok {
					ans += v
					validEnd[j] += v
				}
			}
		}
	}
	return ans
}

func main() {
	var str string
	fmt.Scan(&str)

	ans := solve(str)
	fmt.Println(ans)
}
