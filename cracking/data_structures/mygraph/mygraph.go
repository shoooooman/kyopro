package mygraph

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/shoooooman/kyopro/cracking/data_structures/myqueue"
)

// Graph is a set of nodes
type Graph struct {
	Nodes []*Node
}

// Node has its value and neighbors
type Node struct {
	Val    interface{}
	Neighs []*Node
}

// Path represents a path from start to goal
type Path struct {
	Goal *Node
	Path []*Node
}

// Show shows a graph
func (g *Graph) Show() {
	nodes := g.Nodes
	for _, node := range nodes {
		fmt.Printf("%v: ", node.Val)
		for _, neigh := range node.Neighs {
			fmt.Printf("%v ", neigh.Val)
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
			if v.Goal == goal {
				return v.Path, nil
			}
			for _, neigh := range v.Goal.Neighs {
				if !visited[neigh] {
					path := append(v.Path, neigh)
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
		name := "n" + strconv.Itoa(i)
		nodes[i] = &Node{name, make([]*Node, 0)}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i {
				rand.Seed(time.Now().UnixNano())
				rnd := rand.Intn(2)
				if rnd == 1 {
					nodes[i].Neighs = append(nodes[i].Neighs, nodes[j])
				}
			}
		}
	}
	return &Graph{nodes}
}

// func main() {
// 	graph := NewGraph(5)
// 	graph.Show()
//
// 	start := graph.Nodes[0]
// 	goal := graph.Nodes[3]
// 	path, err := graph.FindPath(start, goal)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%v", start.Val)
// 		for _, node := range path {
// 			fmt.Printf("->%v", node.Val)
// 		}
// 		fmt.Println()
// 	}
// }
