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

// Queue is for bfs
type Queue struct {
	data []interface{}
}

// Push adds a value to queue
func (q *Queue) Push(v interface{}) {
	q.data = append(q.data, v)
}

// Pop removes the last element and return it
func (q *Queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}

// Empty returns true if queue is empty
func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

// NewQueue generates a new queue
func NewQueue() *Queue {
	return &Queue{make([]interface{}, 0)}
}

type node struct {
	Val    int
	Neigh  []*node
	Parent *node
}

func main() {
	n, m := nextInt(), nextInt()
	nodes := make(map[int]*node)
	for i := 0; i < n; i++ {
		nodes[i] = &node{i, nil, nil}
	}

	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		nodes[a].Neigh = append(nodes[a].Neigh, nodes[b])
		nodes[b].Neigh = append(nodes[b].Neigh, nodes[a])
	}

	q := NewQueue()
	q.Push(nodes[0])
	visited := make(map[int]bool)
	for !q.Empty() {
		t, _ := q.Pop()
		top := t.(*node)
		visited[top.Val] = true
		for _, ne := range top.Neigh {
			if !visited[ne.Val] {
				ne.Parent = top
				q.Push(ne)
				visited[ne.Val] = true
			}
		}
	}
	if len(visited) != n {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
	for i := 1; i < n; i++ {
		fmt.Println(nodes[i].Parent.Val + 1)
	}
}
