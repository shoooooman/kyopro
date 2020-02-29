package main

import (
	"fmt"
	"math"

	"github.com/shoooooman/kyopro/cracking/data_structures/mybst"
	"github.com/shoooooman/kyopro/cracking/data_structures/mylist"
	"github.com/shoooooman/kyopro/cracking/data_structures/myqueue"
)

// QData is data for queue
type QData struct {
	node  *mybst.Node
	depth int
}

// Solution using BFS
func solveBFS(tree mybst.Tree, n int) []mylist.LinkedList {
	depth := int(math.Log2(float64(n))) + 1
	lists := make([]mylist.LinkedList, depth)

	q := myqueue.NewQueue()
	root := tree.Root
	q.Push(&QData{root, 0})
	for !q.Empty() {
		src, _ := q.Pop()
		if data, ok := src.(*QData); ok {
			node := data.node
			depth := data.depth
			lists[depth].Add(&mylist.Node{node.Val, nil})
			if node.Left != nil {
				q.Push(&QData{node.Left, depth + 1})
			}
			if node.Right != nil {
				q.Push(&QData{node.Right, depth + 1})
			}
		}
	}
	return lists
}

func dfs(lists []mylist.LinkedList, node *mybst.Node, depth int) {
	if node == nil {
		return
	}
	lists[depth].Add(&mylist.Node{node.Val, nil})
	dfs(lists, node.Left, depth+1)
	dfs(lists, node.Right, depth+1)
}

// Solution using DFS
func solveDFS(tree mybst.Tree, n int) []mylist.LinkedList {
	depth := int(math.Log2(float64(n))) + 1
	lists := make([]mylist.LinkedList, depth)

	root := tree.Root
	dfs(lists, root, 0)
	return lists
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	tree := mybst.GenBST(values)

	lists1 := solveBFS(tree, len(values))
	for i, list := range lists1 {
		fmt.Printf("depth %d: ", i)
		list.Show()
	}

	fmt.Println()

	lists2 := solveBFS(tree, len(values))
	for i, list := range lists2 {
		fmt.Printf("depth %d: ", i)
		list.Show()
	}
}
