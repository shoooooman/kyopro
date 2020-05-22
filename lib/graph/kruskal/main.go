package main

import (
	"fmt"
	"sort"

	"github.com/shoooooman/kyopro/lib/data_structures/unionfind"
)

func kruskal(n int, edges [][]int) [][]int {
	uf := unionfind.NewUnionFind(n)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	ans := make([][]int, 0, n-1)
	// 重みが小さいエッジから閉路ができなよう追加していく
	for _, edge := range edges {
		i := edge[0]
		j := edge[1]
		if !uf.Same(i, j) {
			uf.Unite(i, j)
			ans = append(ans, []int{i, j})
		}
	}
	return ans
}

func main() {
	n := 9
	edges := [][]int{
		{0, 1, 4},
		{1, 2, 8},
		{2, 3, 7},
		{3, 4, 9},
		{4, 5, 10},
		{5, 6, 2},
		{6, 7, 1},
		{7, 8, 7},
		{8, 2, 2},
		{6, 8, 6},
		{2, 5, 4},
		{3, 5, 14},
	}
	ans := kruskal(n, edges)
	fmt.Println(ans)
}
