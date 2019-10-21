package main

import "fmt"

// Point is
type Point struct {
	x, y, depth int
}

// Queue is
type Queue struct {
	ps []Point
}

func (q *Queue) push(x, y, depth int) {
	p := Point{x, y, depth}
	q.ps = append(q.ps, p)
}

func (q *Queue) pop() (int, int, int) {
	p := q.ps[0]
	q.ps = q.ps[1:]
	return p.x, p.y, p.depth
}

func (q *Queue) size() int {
	return len(q.ps)
}

func main() {
	var r, c int
	var sx, sy, gy, gx int

	fmt.Scan(&r, &c)
	fmt.Scan(&sy, &sx, &gy, &gx)

	s := make([][]rune, r)
	for i := 0; i < r; i++ {
		s[i] = make([]rune, c)
	}

	for i := 0; i < r; i++ {
		var str string
		fmt.Scan(&str)
		for j, v := range str {
			s[i][j] = v
		}
	}

	visited := make([][]bool, r)
	for i := 0; i < r; i++ {
		visited[i] = make([]bool, c)
	}

	q := Queue{make([]Point, 0)}
	q.push(sx-1, sy-1, 0)
	for q.size() != 0 {
		x, y, depth := q.pop()

		if visited[y][x] {
			continue
		}
		visited[y][x] = true

		if x == gx-1 && y == gy-1 {
			fmt.Println(depth)
			return
		}

		if x+1 < c && s[y][x+1] == '.' {
			q.push(x+1, y, depth+1)
		}
		if y+1 < r && s[y+1][x] == '.' {
			q.push(x, y+1, depth+1)
		}
		if x-1 >= 0 && s[y][x-1] == '.' {
			q.push(x-1, y, depth+1)
		}
		if y-1 >= 0 && s[y-1][x] == '.' {
			q.push(x, y-1, depth+1)
		}
	}
}
