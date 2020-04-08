package mytree

import (
	"fmt"
)

// Node is a node of tree
type Node struct {
	Val    interface{}
	Parent *Node
	Left   *Node
	Right  *Node
}

// Tree has a root
type Tree struct {
	Root *Node
}

func dfsQuery(node *Node, query interface{}) *Node {
	if node == nil {
		return nil
	}

	if node.Val == query {
		return node
	}
	if n := dfsQuery(node.Left, query); n != nil {
		return n
	}
	if n := dfsQuery(node.Right, query); n != nil {
		return n
	}
	return nil
}

// GetNode return the node whose value equals to query
func (t *Tree) GetNode(query interface{}) *Node {
	return dfsQuery(t.Root, query)
}

func dfsPrint(n *Node) {
	if n == nil {
		return
	}
	left := n.Left
	right := n.Right
	fmt.Printf("%v: ", n.Val)
	if left != nil {
		fmt.Printf("%v ", left.Val)
	}
	if right != nil {
		fmt.Printf("%v ", right.Val)
	}
	fmt.Println()
	dfsPrint(left)
	dfsPrint(right)
}

// Print prints the content of the tree by DFS
func (t *Tree) Print() {
	dfsPrint(t.Root)
}

// GenTree generates a tree from data
func GenTree(m map[int][]int, r int) (*Tree, error) {
	if len(m) == 0 {
		return nil, fmt.Errorf("genRandomTree: n must be larger than 0")
	}

	nodes := make(map[int]*Node)
	root := &Node{r, nil, nil, nil}
	nodes[r] = root
	for parent, children := range m {
		pnode, ok := nodes[parent]
		// 親ノードがなければ作成
		if !ok {
			pnode = &Node{parent, nil, nil, nil}
			nodes[parent] = pnode
		}
		for i, num := range children {
			// -1は無視される; e.g. []int{-1, 5} -> 右の子だけ作られる
			if num == -1 {
				continue
			}
			node, ok := nodes[num]
			// 子ノードがなければ作成
			if !ok {
				node = &Node{num, nil, nil, nil}
				nodes[num] = node
			}
			if i == 0 {
				pnode.Left = node
			} else {
				pnode.Right = node
			}
			node.Parent = pnode
		}
	}
	return &Tree{root}, nil
}

// func main() {
// 	data := map[int][]int{
// 		0: []int{1, 2},
// 		1: []int{3, 4},
// 		2: []int{5},
// 	}
// 	tree, err := GenTree(data, 0)
// 	if err != nil {
// 		fmt.Println("cannot generate a random tree")
// 		os.Exit(1)
// 	}
// 	tree.Print()
// 	node := tree.GetNode(2)
// 	fmt.Println(node.Val)
// }
