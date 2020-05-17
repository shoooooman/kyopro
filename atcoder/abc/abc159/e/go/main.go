package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func cutVertical(s []string, groups map[int]int, numGroups, h, w, k int) int {
	vcuts := 0
	gok := true // この横の切り方で可能かどうか
	white := make([]int, numGroups)
	for j := 0; j < w; j++ {
		colWhite := make([]int, numGroups)
		ok := true
		for i := 0; i < h; i++ {
			colWhite[groups[i]] += int(rune(s[i][j]) - '0')
			white[groups[i]] += int(rune(s[i][j]) - '0')
			if colWhite[groups[i]] > k {
				gok = false
			}
			if white[groups[i]] > k {
				ok = false
			}
		}
		// 縦に切る必要がある場合
		if !ok {
			vcuts++
			white = colWhite
		}
	}
	if gok {
		return vcuts
	}
	return math.MaxInt64 / 2
}

// 横の切り方でチョコをグルーピング
func grouping(mask, h int) (map[int]int, int) {
	groups := make(map[int]int)
	groups[0] = 0
	num := 0
	for i := 1; i < h; i++ {
		if mask&(1<<uint(i-1)) == 0 {
			groups[i] = groups[i-1]
		} else {
			groups[i] = groups[i-1] + 1
			num++
		}
	}
	return groups, num
}

func solve(s []string, h, w, k int) int {
	minCuts := math.MaxInt64
	for mask := 0; mask < (1 << uint(h-1)); mask++ {
		groups, hcuts := grouping(mask, h)
		vcuts := cutVertical(s, groups, hcuts+1, h, w, k)
		if hcuts+vcuts < minCuts {
			minCuts = hcuts + vcuts
		}
	}
	return minCuts
}

func main() {
	h, w, k := nextInt(), nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	ans := solve(s, h, w, k)
	fmt.Println(ans)
}
