package main

import (
	"fmt"
	"math"
	"os"
)

// Node is a node of tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Tree has a root
type Tree struct {
	Root *Node
}

func dfs(n *Node) {
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
	dfs(left)
	dfs(right)
}

// Print prints the content of the tree by DFS
func (t *Tree) Print() {
	dfs(t.Root)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func checkBalanced(n *Node, height int) int {
	if n == nil {
		return 0
	}
	heightL := checkBalanced(n.Left, height)
	heightR := checkBalanced(n.Right, height)
	fmt.Println(n.Val, heightL, heightR)
	// 左右どちらかの部分木が平衡でない場合
	if heightL == -1 || heightR == -1 {
		return -1
	}
	// このノードの左右の部分木が平衡でない場合
	if math.Abs(float64(heightL)-float64(heightR)) > 1.0 {
		return -1
	}
	return max(heightL, heightR) + 1
}

func genRandomTree(m map[int][]int) (*Tree, error) {
	if len(m) == 0 {
		return nil, fmt.Errorf("genRandomTree: n must be larger than 0")
	}

	nodes := make(map[int]*Node)
	root := &Node{0, nil, nil}
	nodes[0] = root
	for parent, children := range m {
		pnode, ok := nodes[parent]
		// 親ノードがなければ作成
		if !ok {
			pnode = &Node{parent, nil, nil}
			nodes[parent] = pnode
		}
		for i, num := range children {
			if num == -1 {
				continue
			}
			node, ok := nodes[num]
			// 子ノードがなければ作成
			if !ok {
				node = &Node{num, nil, nil}
				nodes[num] = node
			}
			if i == 0 {
				pnode.Left = node
			} else {
				pnode.Right = node
			}
		}
	}
	return &Tree{root}, nil
}

func main() {
	data := map[int][]int{
		0: []int{1, 2},
		1: []int{3, 4},
		2: []int{5},
		3: []int{6},
		6: []int{7},
	}
	tree, err := genRandomTree(data)
	if err != nil {
		fmt.Println("cannot generate a random tree")
		os.Exit(1)
	}
	tree.Print()

	h := checkBalanced(tree.Root, 0)
	if h != -1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
