package main

import (
	"fmt"
	"strings"
)

type point struct {
	x int
	y int
}

var dp map[point][]string
var locked map[point]bool

func appendToAll(bases []string, as string) []string {
	strs := make([]string, 0)
	for _, base := range bases {
		// lockedが含まれていたら更新しない&&含まない
		if !strings.Contains(base, "[locked]") {
			strs = append(strs, base+as)
		}
	}
	return strs
}

// 解説では一つの道順を見つける解法を紹介しているが、
// このアルゴリズムは全ての道順を見つける
func solve(p point, routes []string) []string {
	if val, ok := dp[p]; ok {
		return val
	}

	x := p.x
	y := p.y
	dx := point{x - 1, y}
	dy := point{x, y - 1}

	if locked[dx] && locked[dy] {
		rts := appendToAll(routes, "[locked]")
		dp[p] = rts
		return dp[p]
	} else if locked[dy] {
		rtsdx := solve(dx, routes)
		rtsX := appendToAll(rtsdx, "→")
		rtsY := appendToAll(rtsdx, "[locked]")
		dp[p] = append(rtsX, rtsY...)
		return dp[p]
	} else if locked[dx] {
		rtsdy := solve(dy, routes)
		rtsX := appendToAll(rtsdy, "[locked]")
		rtsY := appendToAll(rtsdy, "↓")
		dp[p] = append(rtsX, rtsY...)
		return dp[p]
	}

	rtsdx := solve(dx, routes)
	rtsX := appendToAll(rtsdx, "→")
	rtsdy := solve(dy, routes)
	rtsY := appendToAll(rtsdy, "↓")
	dp[p] = append(rtsX, rtsY...)
	return dp[p]
}

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var k int
	fmt.Scan(&k)
	locked = make(map[point]bool, k)
	for i := 0; i < k; i++ {
		var lx, ly int
		fmt.Scan(&lx, &ly)
		locked[point{lx, ly}] = true
	}
	// 枠外もlockedに加える
	for i := 0; i < c; i++ {
		locked[point{-1, i}] = true
	}
	for i := 0; i < r; i++ {
		locked[point{i, -1}] = true
	}

	dp = make(map[point][]string)
	dp[point{0, 0}] = make([]string, 1)
	routes := solve(point{c - 1, r - 1}, make([]string, 1))
	for _, route := range routes {
		fmt.Println(route)
	}
}
