package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* Solution using UnionFind */

// UnionFind is union find structure
type UnionFind struct {
	Par   []int
	Ranks []int
	Sizes []int
}

// NewUnionFind generates a new union find
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{make([]int, n), make([]int, n), make([]int, n)}
	for i := range uf.Par {
		uf.Par[i] = i
		uf.Ranks[i] = 0
		uf.Sizes[i] = 1
	}
	return uf
}

// Root searches the root of the arguement
func (uf *UnionFind) Root(x int) int {
	if x == uf.Par[x] {
		return x
	}
	uf.Par[x] = uf.Root(uf.Par[x]) // 経路圧縮
	return uf.Par[x]
}

// Unite unites two trees by adding one root to another
func (uf *UnionFind) Unite(x, y int) {
	rx := uf.Root(x)
	ry := uf.Root(y)
	if rx == ry {
		return
	}
	if uf.Ranks[rx] < uf.Ranks[ry] {
		uf.Par[rx] = ry
		uf.Sizes[ry] += uf.Sizes[rx]
	} else {
		uf.Par[ry] = rx
		uf.Sizes[rx] += uf.Sizes[ry]
		if uf.Ranks[rx] == uf.Ranks[ry] {
			uf.Ranks[rx]++
		}
	}
}

// Same returns true when the two are in the same group
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// Size returns size of a group where x belongs
func (uf *UnionFind) Size(x int) int {
	return uf.Sizes[uf.Root(x)]
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m, k := nextInt(), nextInt(), nextInt()
	uf := NewUnionFind(n)

	friends := make(map[int]int)
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		uf.Unite(a, b)
		friends[a]++
		friends[b]++
	}

	// sizes := make([]int, n)
	// for i := range sizes {
	// 	r := uf.Root(uf.Par[i])
	// 	sizes[r]++
	// }

	blocks := make([]int, n)
	for i := 0; i < k; i++ {
		c, d := nextInt()-1, nextInt()-1
		if uf.Same(c, d) {
			blocks[c]++
			blocks[d]++
		}
	}

	for i := 0; i < n; i++ {
		ans := uf.Size(i) - friends[i] - blocks[i] - 1
		fmt.Printf("%v ", ans)
	}
	fmt.Println()
}
