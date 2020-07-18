package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

func weave(left, right, list []int, rst *[][]int) {
	if len(left) == 0 || len(right) == 0 {
		if len(left) == 0 {
			list = append(list, right...)
			*rst = append(*rst, list)
			return
		}
		list = append(list, left...)
		*rst = append(*rst, list)
		return
	}

	weave(left[1:], right, append(list, left[0]), rst)
	weave(left, right[1:], append(list, right[0]), rst)
}

func bstArrays(root *mytree.Node) [][]int {
	if root == nil {
		return make([][]int, 1)
	}

	ans := make([][]int, 0)
	left := bstArrays(root.Left)
	right := bstArrays(root.Right)
	for _, l := range left {
		for _, r := range right {
			rst := make([][]int, 0)
			weave(l, r, []int{root.Val.(int)}, &rst)
			ans = append(ans, rst...)
		}
	}
	return ans
}

func main() {
	// data := map[int][]int{
	// 	// 2: {1, 3},
	// 	// 5: {3, 8},
	// 	// 3: {2, 4},
	// 	// 8: {7, 9},
	// 	8:  {4, 12},
	// 	4:  {2, 6},
	// 	2:  {1, 3},
	// 	6:  {5, 7},
	// 	12: {10, 14},
	// 	10: {9, 11},
	// 	14: {13, 15},
	// }
	//
	// t, _ := mytree.GenTree(data, 8)
	// ans := bstArrays(t.Root)

	root := &mytree.Node{
		4,
		nil,
		&mytree.Node{
			2, nil, &mytree.Node{1, nil, nil, nil}, &mytree.Node{3, nil, nil, nil},
		},
		&mytree.Node{
			6, nil, &mytree.Node{5, nil, nil, nil}, &mytree.Node{7, nil, nil, nil},
		},
	}
	ans := bstArrays(root)

	fmt.Println(ans)
}
