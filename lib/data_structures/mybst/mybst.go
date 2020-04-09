package mybst

import "fmt"

// Node has its value and two children
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Tree has a root
type Tree struct {
	Root *Node
}

// PrintValues prints sub-tree in in-order
func PrintValues(node *Node) {
	if node == nil {
		return
	}
	PrintValues(node.Left)
	fmt.Printf("%v ", node.Val)
	PrintValues(node.Right)
}

// printNode prints its value and children
func printNode(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%v: ", node.Val)
	if node.Left != nil {
		fmt.Printf("%v ", node.Left.Val)
	}
	if node.Right != nil {
		fmt.Printf("%v ", node.Right.Val)
	}
	fmt.Println()
	printNode(node.Left)
	printNode(node.Right)
}

// PrintTree prints tree
func (t *Tree) PrintTree() {
	root := t.Root
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
	root.Left = addChildren(values[:mid])
	if l > 2 {
		root.Right = addChildren(values[mid+1:])
	}
	return root
}

// GenBST generates a binary tree with values
func GenBST(values []int) *Tree {
	root := addChildren(values)
	return &Tree{root}
}

// func main() {
// 	values := []int{0, 1, 2, 3, 4, 5, 6}
// 	// values := []int{0, 1, 2, 3, 4, 5, 6, 7}
// 	tree := GenBST(values)
// 	tree.PrintTree()
// 	PrintValues(tree.Root)
// }
