package main

import (
	"fmt"
	"sort"
)

/* 最も重複が多いところから取っていく */
/* TLE & 最適解でないかも */
// func remove(requests [][]int, request int) {
// 	for i, v := range requests {
// 		result := make([]int, 0)
// 		for _, r := range v {
// 			if r != request {
// 				result = append(result, r)
// 			}
// 		}
// 		requests[i] = result
// 	}
// }
//
// // Requests is
// type Requests [][]int
//
// func (e Requests) Len() int {
// 	return len(e)
// }
//
// func (e Requests) Swap(i, j int) {
// 	e[i], e[j] = e[j], e[i]
// }
//
// func (e Requests) Less(i, j int) bool {
// 	return len(e[i]) > len(e[j])
// }
//
// func main() {
// 	var n, m int
// 	fmt.Scan(&n, &m)
//
// 	a := make([]int, m)
// 	b := make([]int, m)
// 	for i := 0; i < m; i++ {
// 		fmt.Scan(&a[i], &b[i])
// 		a[i]--
// 		b[i]--
// 	}
//
// 	requests := make(Requests, n-1)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < m; j++ {
// 			if i >= a[j] && i < b[j] {
// 				requests[i] = append(requests[i], j)
// 			}
// 		}
// 	}
//
// 	sort.Sort(requests)
//
// 	done := make([]bool, m)
// 	count := 0
// L:
// 	for i := 0; i < n-1; i++ {
// 		for _, request := range requests[i] {
// 			if !done[request] {
// 				done[request] = true
// 				count++
// 				if count == m {
// 					fmt.Println(i + 1)
// 					break L
// 				}
// 				if i != n-1 {
// 					remove(requests[i+1:], request)
// 				}
// 			}
// 		}
// 		sort.Sort(requests[i+1:])
// 	}
// }

type request struct {
	a, b int
}

type requests []request

func (e requests) Len() int {
	return len(e)
}

func (e requests) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e requests) Less(i, j int) bool {
	return e[i].b < e[j].b
}

/* 終端が大きいものから順に取っていく */
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	reqs := make(requests, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		r := request{a, b}
		reqs[i] = r
	}

	// sort.Slice(reqs, func(i, j int) bool { return requests[i].b < requests[j].b })
	sort.Sort(reqs)
	ans := 0
	end := 0
	for i := range reqs {
		if reqs[i].a >= end {
			ans++
			end = reqs[i].b
		}
	}
	fmt.Println(ans)
}
