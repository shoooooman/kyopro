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

// Move v to u's position
// Note that u's info won't be changed
func transparent(tree *Tree, u, v *Node) {
	if u.Parent == nil {
		tree.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// FindMin returns the min node of a subtree of the root
func FindMin(root *Node) *Node {
	if root == nil {
		return nil
	}

	n := root
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// Delete deletes a node
func (t *Tree) Delete(node *Node) {
	if node.Left == nil {
		transparent(t, node, node.Right)
	} else if node.Right == nil {
		transparent(t, node, node.Left)
	} else {
		minNode := FindMin(node.Right)
		if minNode.Parent != node {
			transparent(t, minNode, minNode.Right)
			minNode.Right = node.Right
			minNode.Right.Parent = minNode
		}
		transparent(t, node, minNode)
		minNode.Left = node.Left
		minNode.Left.Parent = minNode
	}
}

// FindNext returns next node to the given node
func FindNext(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return FindMin(node.Right)
	}

	n := node
	p := node.Parent
	for p != nil && p.Left != n {
		n = p
		p = p.Parent
	}
	return p
}

// func main() {
// 	values := []int{0, 1, 2, 3, 4, 5, 6}
// 	// values := []int{0, 1, 2, 3, 4, 5, 6, 7}
// 	tree := GenBST(values)
// 	tree.PrintTree()
// 	PrintValues(tree.Root)
// }
