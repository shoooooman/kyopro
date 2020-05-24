package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mybst"
)

type item struct {
	Node  *mybst.Node
	Depth int
}

type queue []*item

func (q *queue) Push(it *item) {
	*q = append(*q, it)
}

func (q *queue) Pop() *item {
	top := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return top
}

func (q queue) Empty() bool {
	return len(q) == 0
}

func depthLists(root *mybst.Node) map[int][]*mybst.Node {
	q := new(queue)
	q.Push(&item{root, 0})
	lists := make(map[int][]*mybst.Node)
	for !q.Empty() {
		top := q.Pop()
		depth := top.Depth
		nd := top.Node
		lists[depth] = append(lists[depth], nd)
		if nd.Left != nil {
			q.Push(&item{nd.Left, depth + 1})
		}
		if nd.Right != nil {
			q.Push(&item{nd.Right, depth + 1})
		}
	}
	return lists
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	bst := mybst.GenBST(nums)
	ans := depthLists(bst.Root)
	for d, nodes := range ans {
		fmt.Printf("%d: ", d)
		for _, n := range nodes {
			fmt.Printf("%d ", n.Val)
		}
		fmt.Println()
	}
}
