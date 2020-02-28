package main

import "fmt"

// Point has x and y
type Point struct {
	x int
	y int
}

// Queue is for bfs
type Queue struct {
	points []Point
}

// Push adds a point to queue
func (q *Queue) Push(p Point) {
	q.points = append(q.points, p)
}

// Pop removes the last point and return it
func (q *Queue) Pop() (Point, error) {
	if q.Empty() {
		return Point{}, fmt.Errorf("queue is empty")
	}
	p := q.points[0]
	q.points = q.points[1:]
	return p, nil
}

// Empty returns true if queue is empty
func (q *Queue) Empty() bool {
	return len(q.points) == 0
}

// addNeigh adds point to queue if the point have not been visited and can pass
func addNeigh(q *Queue, p Point, pass map[Point]bool, dist map[Point]int) {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 0; i < 4; i++ {
		dp := Point{p.x + dx[i], p.y + dy[i]}
		_, exist := dist[dp]
		if !exist && pass[dp] {
			q.Push(dp)
			dist[dp] = dist[p] + 1
		}
	}
}

func solve(start, goal Point, pass map[Point]bool) int {
	q := &Queue{make([]Point, 0)}
	dist := make(map[Point]int)
	dist[start] = 0
	q.Push(start)
	for !q.Empty() {
		p, _ := q.Pop()
		addNeigh(q, p, pass, dist)
	}
	return dist[goal]
}

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var sx, sy, gx, gy int
	fmt.Scan(&sy, &sx, &gy, &gx)
	start := Point{sx - 1, sy - 1}
	goal := Point{gx - 1, gy - 1}

	pass := make(map[Point]bool)
	for i := 0; i < r; i++ {
		var row string
		fmt.Scan(&row)
		for j, c := range row {
			if c == '.' {
				pass[Point{j, i}] = true
			} else {
				if c != '#' {
					fmt.Errorf("input error")
				}
			}
		}
	}

	ans := solve(start, goal, pass)
	fmt.Println(ans)
}
