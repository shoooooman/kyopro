package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/cracking/data_structures/mytree"
)

func findMostLeft(node *mytree.Node) *mytree.Node {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return findMostLeft(node.Left)
}

func findRootOf(node *mytree.Node) *mytree.Node {
	pnode := node.Parent
	// 木の右端の場合にはnilを返す
	if pnode == nil {
		return nil
	}
	if node == pnode.Left {
		return pnode
	}
	return findRootOf(pnode)
}

func findNext(node *mytree.Node) *mytree.Node {
	// 右の部分木がある場合は、その中で最も左のノード
	if node.Right != nil {
		return findMostLeft(node.Right)
	}
	// ない場合は、nodeが属する部分木を左の子としてもつルートノード
	return findRootOf(node)
}

func main() {
	data := map[int][]int{
		8:  []int{4, 10},
		4:  []int{2, 6},
		10: []int{-1, 20},
	}
	tree, _ := mytree.GenTree(data, 8)
	tree.Print()

	node := tree.GetNode(6)
	next := findNext(node)
	if next != nil {
		fmt.Println(next.Val)
	} else {
		fmt.Println("nil")
	}
}
