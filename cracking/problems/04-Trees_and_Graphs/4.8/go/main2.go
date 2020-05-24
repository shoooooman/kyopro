package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

// 見つけたノードの数(<=2)と共通のノードを返す
func helper(root, n1, n2 *mytree.Node) (int, *mytree.Node) {
	if root == nil {
		return 0, nil
	}

	lcnt, left := helper(root.Left, n1, n2)
	if lcnt == 2 {
		return 2, left
	}
	rcnt, right := helper(root.Right, n1, n2)
	if rcnt == 2 {
		return 2, right
	}
	if lcnt == 1 && rcnt == 1 {
		return 2, root
	}
	if lcnt == 1 {
		if root == n1 || root == n2 {
			return 2, root
		}
		return 1, nil
	}
	if rcnt == 1 {
		if root == n1 || root == n2 {
			return 2, root
		}
		return 1, nil
	}
	if root == n1 || root == n2 {
		if root == n1 && root == n2 {
			return 2, root
		}
		return 1, nil
	}
	return 0, nil
}

func commonParent(root, n1, n2 *mytree.Node) *mytree.Node {
	cnt, common := helper(root, n1, n2)
	if cnt == 2 {
		return common
	}
	return nil
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
	r := commonParent(tree.Root, n1, n2)
	if r != nil {
		fmt.Println(r.Val)
	} else {
		fmt.Println("no common")
	}
}
