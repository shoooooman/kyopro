package main

import (
	"bufio"
	"container/heap"
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

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type node struct {
	ID    int
	Val   int
	Edges []*edge
}

type edge struct {
	Child  *node
	Weight int
}

func initNodes(n int, edges [][]int) []*node {
	nodes := make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{i, math.MaxInt64 / 3, make([]*edge, 0)}
	}
	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]
		nodes[from].Edges = append(nodes[from].Edges, &edge{nodes[to], weight})
	}
	return nodes
}

type nodeHeap []*node

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*node)) }
func (h *nodeHeap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func dijkstra(n int, edges [][]int, src int) map[int]int {
	nodes := nodeHeap(initNodes(n, edges))
	// nodes[src].Val = 0
	// h := &nodes
	h := &nodeHeap{}
	heap.Push(h, &node{src, 0, nodes[src].Edges})
	heap.Init(h)

	dists := make(map[int]int)
	dists[src] = 0
	visited := make(map[int]struct{})
	for h.Len() > 0 {
		nd := heap.Pop(h).(*node)
		if _, ok := visited[nd.ID]; ok {
			continue
		}
		visited[nd.ID] = struct{}{}
		for _, e := range nd.Edges {
			if _, ok := visited[e.Child.ID]; ok {
				continue
			}
			if nd.Val+e.Weight < e.Child.Val {
				updated := &node{e.Child.ID, nd.Val + e.Weight, e.Child.Edges}
				dists[e.Child.ID] = updated.Val
				// e.Child.Val = nd.Val + e.Weight
				// dists[e.Child.ID] = e.Child.Val
				heap.Push(h, updated)
			}
		}
	}

	return dists
}

func main() {
	n, m, t := nextInt(), nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	edges := make([][]int, m)
	rvs := make([][]int, m)
	for i := 0; i < m; i++ {
		from, to, weight := nextInt()-1, nextInt()-1, nextInt()
		edges[i] = []int{from, to, weight}
		rvs[i] = []int{to, from, weight}
	}
	way := dijkstra(n, edges, 0)
	back := dijkstra(n, rvs, 0)
	max := 0
	for i := 0; i < n; i++ {
		moveTime := way[i] + back[i]
		stayTime := t - moveTime
		if stayTime > 0 {
			profit := stayTime * a[i]
			if profit > max {
				max = profit
			}
		}
	}
	fmt.Println(max)
}
