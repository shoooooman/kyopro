package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mybst"
)

func appendLists(orders [][]int, lists [][]int) [][]int {
	joined := make([][]int, len(orders)*len(lists))
	k := 0
	for i := range orders {
		for j := range lists {
			joined[k] = append(orders[i], lists[j]...)
			k++
		}
	}
	return joined
}

func appendAll(list []int, orders [][]int) [][]int {
	joined := make([][]int, len(orders))
	for i := range orders {
		joined[i] = append(list, orders[i]...)
	}
	return joined
}

func doubleDFS(node *mybst.Node, orders [][]int) [][]int {
	if node == nil {
		return nil
	}

	left := doubleDFS(node.Left, orders)
	right := doubleDFS(node.Right, orders)

	self := []int{node.Val}
	if left == nil && right == nil {
		return [][]int{self}
	} else if left != nil && right == nil {
		return appendAll(self, left)
	} else if right != nil && left == nil {
		return appendAll(self, right)
	}
	ltor := appendLists(left, right)
	rtol := appendLists(right, left)
	next1 := appendAll(self, ltor)
	next2 := appendAll(self, rtol)
	return append(next1, next2...)
}

func solve(bst *mybst.Tree) [][]int {
	orders := make([][]int, 0)
	return doubleDFS(bst.Root, orders)
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}
	// values := []int{1, 2, 3}
	bst := mybst.GenBST(values)
	bst.PrintTree()
	mybst.PrintValues(bst.Root)
	fmt.Println()

	arrays := solve(bst)
	fmt.Println(arrays)
}
