package main

import "fmt"

// Node has its value and two children
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Tree has a root
type Tree struct {
	root *Node
}

// printNode prints elements in in-order
func printNode(node *Node) {
	if node == nil {
		return
	}
	printNode(node.left)
	fmt.Printf("%v ", node.val)
	printNode(node.right)
}

func printTree(tree Tree) {
	root := tree.root
	printNode(root)
	fmt.Println()
}

// addChildren inserts nodes and create children recursively
func addChildren(values []int) *Node {
	l := len(values)
	if l == 0 {
		return nil
	}

	mid := l / 2
	root := &Node{values[mid], nil, nil}
	if l == 1 {
		return root
	}
	// lが偶数の場合、左部分木が大きくなるようにする
	root.left = addChildren(values[:mid])
	if l > 2 {
		root.right = addChildren(values[mid+1:])
	}
	return root
}

func getBinaryTree(values []int) Tree {
	root := addChildren(values)
	return Tree{root}
}

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	// values := []int{0, 1, 2, 3, 4, 5, 6, 7}
	tree := getBinaryTree(values)
	printTree(tree)
}
