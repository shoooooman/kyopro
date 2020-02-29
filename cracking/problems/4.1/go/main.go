package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shoooooman/kyopro/cracking/data_structures/myqueue"
)

// Graph is a set of nodes
type Graph struct {
	nodes []*Node
}

// Node has its value and neighbors
type Node struct {
	val    int
	neighs []*Node
}

// Path represents a path from start to goal
type Path struct {
	goal *Node
	path []*Node
}

// Show shows a graph
func (g *Graph) Show() {
	nodes := g.nodes
	for _, node := range nodes {
		fmt.Printf("%d: ", node.val)
		for _, neigh := range node.neighs {
			fmt.Printf("%d ", neigh.val)
		}
		fmt.Println()
	}
}

// FindPath finds a path from start to goal
func (g *Graph) FindPath(start, goal *Node) ([]*Node, error) {
	visited := make(map[*Node]bool)
	q := myqueue.NewQueue()
	q.Push(&Path{start, make([]*Node, 0)})
	for !q.Empty() {
		src, _ := q.Pop()
		if v, ok := src.(*Path); ok {
			if v.goal == goal {
				return v.path, nil
			}
			for _, neigh := range v.goal.neighs {
				if !visited[neigh] {
					path := append(v.path, neigh)
					q.Push(&Path{neigh, path})
					visited[neigh] = true
				}
			}
		}
	}
	return nil, fmt.Errorf("cannot find a path")
}

// NewGraph generates a new directed graph
func NewGraph(n int) *Graph {
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{i, make([]*Node, 0)}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i {
				rand.Seed(time.Now().UnixNano())
				rnd := rand.Intn(2)
				if rnd == 1 {
					nodes[i].neighs = append(nodes[i].neighs, nodes[j])
				}
			}
		}
	}
	return &Graph{nodes}
}

func main() {
	graph := NewGraph(5)
	graph.Show()

	start := graph.nodes[0]
	goal := graph.nodes[3]
	path, err := graph.FindPath(start, goal)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", start.val)
		for _, node := range path {
			fmt.Printf("->%v", node.val)
		}
		fmt.Println()
	}
}
