package main

import "fmt"

var e [][]int
var n, m int

func dfs(v int, visited []bool) int {
	comp := true
	for i := 1; i < n+1; i++ {
		if !visited[i] {
			comp = false
			break
		}
	}

	if comp {
		return 1
	}

	count := 0
	for _, ne := range e[v] {
		if visited[ne] {
			continue
		}
		visited[ne] = true
		count += dfs(ne, visited)
		visited[ne] = false
	}
	return count
}

func main() {
	fmt.Scan(&n, &m)

	e = make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		e[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}

	visited := make([]bool, n+1)
	visited[1] = true
	count := dfs(1, visited)
	fmt.Println(count)
}
