package main

import (
	"fmt"
	"math"
)

type pair struct {
	first  int
	second int
}

// a^2+b^2の値を全て計算し、あとでその中の組み合わせを求める
func main() {
	const n = 100

	h := make(map[int]int)
	pairs := make(map[int][]pair)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			p := int(math.Pow(float64(i), 3.0) + math.Pow(float64(j), 3.0))
			v, ok := h[p]
			if !ok {
				h[p] = 1
				pairs[p] = []pair{{i, j}}
			} else {
				h[p] = v + 1
				pairs[p] = append(pairs[p], pair{i, j})
			}
		}
	}

	for _, ps := range pairs {
		for _, p1 := range ps {
			for _, p2 := range ps {
				fmt.Println(p1.first, p1.second, p2.first, p2.second)
			}
		}
	}

	sum := 0
	for _, v := range h {
		sum += int(math.Pow(float64(v), 2.0))
	}
	fmt.Println(sum)
}
