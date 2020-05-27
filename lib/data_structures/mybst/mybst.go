package mybst

import "fmt"

// Node has its value and two children
type Node struct {
	Val    int
	Parent *Node
	Left   *Node
	Right  *Node
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

	pval := -1
	if node.Parent != nil {
		pval = node.Parent.Val
	}
	fmt.Printf("%v (p=%v): ", node.Val, pval)
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

// insertNodes inserts nodes with sorted values
// and keeps the balance
func insertNodes(values []int, parent *Node) *Node {
	l := len(values)
	if l == 0 {
		return nil
	}

	mid := l / 2
	root := &Node{values[mid], parent, nil, nil}
	if l == 1 {
		return root
	}
	// lが偶数の場合、左部分木が大きくなるようにする
	root.Left = insertNodes(values[:mid], root)
	if l > 2 {
		root.Right = insertNodes(values[mid+1:], root)
	}
	return root
}

// GenBST generates a binary tree with values
func GenBST(values []int) *Tree {
	root := insertNodes(values, nil)
	return &Tree{root}
}

// Insert inserts a new node with value val
func (t *Tree) Insert(val int) *Node {
	root := t.Root
	var p *Node
	n := root
	for n != nil {
		p = n
		if val <= n.Val {
			n = n.Left
		} else {
			n = n.Right
		}
	}

	newNode := &Node{val, p, nil, nil}
	if val <= p.Val {
		p.Left = newNode
	} else {
		p.Right = newNode
	}
	return newNode
}

// func main() {
// 	values := []int{0, 1, 2, 3, 4, 5, 6}
// 	// values := []int{0, 1, 2, 3, 4, 5, 6, 7}
// 	tree := GenBST(values)
// 	tree.PrintTree()
// 	PrintValues(tree.Root)
// }
