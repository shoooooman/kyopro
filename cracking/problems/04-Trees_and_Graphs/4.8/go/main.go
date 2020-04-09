package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

func findCommonRoot(node *mytree.Node, n1, n2 *mytree.Node) *mytree.Node {
	if node == nil {
		return nil
	}

	if node == n1 || node == n2 {
		return node
	}

	left := findCommonRoot(node.Left, n1, n2)
	right := findCommonRoot(node.Right, n1, n2)
	if left == nil && right == nil {
		return nil
	}
	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	return node
}

func main() {
	data := map[int][]int{
		0: []int{1, 2},
		1: []int{3, 4},
		2: []int{5, 6},
		3: []int{},
		4: []int{7, 8},
		5: []int{},
		6: []int{-1, 9},
		7: []int{},
		8: []int{10},
	}
	tree, _ := mytree.GenTree(data, 0)

	n1 := tree.GetNode(3)
	n2 := tree.GetNode(10)
	r := findCommonRoot(tree.Root, n1, n2)
	fmt.Println(r.Val)
}
