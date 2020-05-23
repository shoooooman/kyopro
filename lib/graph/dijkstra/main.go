package main

import (
	"container/heap"
	"fmt"
	"math"
)

type node struct {
	ID    int
	Val   int
	Edges []*edge
}

type edge struct {
	Child  *node
	Weight int
}

// func initNodes(n int, edges [][]int) map[int]*node {
func initNodes(n int, edges [][]int) []*node {
	// nodes := make(map[int]*node)
	nodes := make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{i, math.MaxInt64, make([]*edge, 0)}
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

func printHeap(h *nodeHeap) {
	for i := 0; i < h.Len(); i++ {
		fmt.Printf("(%d, %d) ", (*h)[i].ID, (*h)[i].Val)
	}
	fmt.Println()
}

func dijkstra(n int, edges [][]int, src int) map[int]int {
	nodes := nodeHeap(initNodes(n, edges))
	nodes[src].Val = 0
	h := &nodes
	heap.Init(h)

	visited := make(map[*node]struct{})
	for h.Len() > 0 {
		n := heap.Pop(h).(*node)
		visited[n] = struct{}{}
		for _, e := range n.Edges {
			if n.Val+e.Weight < e.Child.Val {
				e.Child.Val = n.Val + e.Weight
				// cannot know indices whose item changed
				// heap.Fix(h, e.Child.HeapIndex)
			}
		}
		// not efficient
		heap.Init(h)
		// printHeap(h)
	}

	dists := make(map[int]int)
	for n := range visited {
		dists[n.ID] = n.Val
	}
	return dists
}

func main() {
	edges := [][]int{
		// {0, 1, 10},
		// {1, 2, 1},
		// {2, 3, 4},
		// {3, 2, 6},
		// {4, 3, 2},
		// {0, 4, 5},
		// {1, 4, 2},
		// {4, 1, 3},
		// {3, 0, 7},
		// {4, 2, 9},

		{0, 1, 100},
		{0, 2, 1},
		{0, 3, 10},
		{2, 1, 1},
	}
	n := 4
	src := 0
	ans := dijkstra(n, edges, src)
	fmt.Println(ans)
}
