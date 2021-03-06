package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 10000000
)

var (
	sc = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		sc.Split(bufio.ScanWords)
		return sc
	}()
)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type item struct {
	A, B, C, D int
}

var n, m int

// 条件を満たす数列を全て生成
func dfs(seq []int, rst *[][]int) {
	if len(seq) == n {
		*rst = append(*rst, seq)
		return
	}

	// 1つ前の要素の数 (最初の要素の場合は1)
	var last int
	if len(seq) == 0 {
		last = 1
	} else {
		last = seq[len(seq)-1]
	}
	// last以上M以下の数を追加
	for i := last; i <= m; i++ {
		dfs(append(seq, i), rst)
	}
}

// 数列の得点を制約より計算
func calcPoint(seq []int, items []item) int {
	point := 0
	for _, it := range items {
		if seq[it.B]-seq[it.A] == it.C {
			point += it.D
		}
	}
	return point
}

// 全ての数列の中から最大の得点を計算
func findMaxPoint(seqs [][]int, items []item) int {
	max := 0
	for _, seq := range seqs {
		if v := calcPoint(seq, items); v > max {
			max = v
		}
	}
	return max
}

func solve(items []item) int {
	seq := make([]int, 0)
	rst := make([][]int, 0)
	dfs(seq, &rst)
	return findMaxPoint(rst, items)
}

func main() {
	n, m = nextInt(), nextInt()
	q := nextInt()
	items := make([]item, q)
	for i := 0; i < q; i++ {
		items[i] = item{nextInt() - 1, nextInt() - 1, nextInt(), nextInt()}
	}
	ans := solve(items)
	fmt.Println(ans)
}
