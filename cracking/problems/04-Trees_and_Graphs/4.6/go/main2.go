package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

func mostLeft(node *mytree.Node) *mytree.Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func parentOfLeft(node *mytree.Node) *mytree.Node {
	for node.Parent != nil {
		parent := node.Parent
		if parent.Left == node {
			return parent
		}
		node = parent
	}
	return nil
}

func nextNode(node *mytree.Node) *mytree.Node {
	if node.Right != nil {
		return mostLeft(node.Right)
	}

	return parentOfLeft(node)
}

func main() {
	data := map[int][]int{
		8:  {4, 10},
		4:  {2, 6},
		2:  {1, 3},
		6:  {5, 7},
		10: {9, 11},
	}
	t, _ := mytree.GenTree(data, 8)
	// node := t.Root.Left.Right.Right
	node := t.Root.Right.Right
	fmt.Printf("%d ", node.Val.(int))
	next := nextNode(node)
	if next != nil {
		fmt.Println(next.Val.(int))
	} else {
		fmt.Println("no next node")
	}
}
