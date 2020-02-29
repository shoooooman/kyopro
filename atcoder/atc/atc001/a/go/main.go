package main

import "fmt"

var h, w int
var c [][]rune
var sx, sy int

func dfs(i, j int, passed [][]bool) bool {
	if passed[i][j] {
		return false
	}
	passed[i][j] = true

	if c[i][j] == 'g' {
		return true
	} else if c[i][j] == '#' {
		return false
	}

	if i+1 < h && dfs(i+1, j, passed) {
		return true
	}
	if j+1 < w && dfs(i, j+1, passed) {
		return true
	}
	if i-1 >= 0 && dfs(i-1, j, passed) {
		return true
	}
	if j-1 >= 0 && dfs(i, j-1, passed) {
		return true
	}
	return false
}

func main() {
	fmt.Scan(&h, &w)

	c = make([][]rune, h)
	for i := 0; i < h; i++ {
		c[i] = make([]rune, w)
	}

	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		for j, v := range s {
			c[i][j] = v
			if v == 's' {
				sx, sy = i, j
			}
		}
	}

	passed := make([][]bool, h)
	for i := 0; i < h; i++ {
		passed[i] = make([]bool, w)
	}

	if dfs(sx, sy, passed) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
