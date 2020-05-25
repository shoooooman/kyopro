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
	data := map[int][]int{
		// 2: {1, 3},
		5: {3, 8},
		3: {2, 4},
		// 8: {7, 9},
	}

	t, _ := mytree.GenTree(data, 5)
	ans := bstArrays(t.Root)
	fmt.Println(ans)
}
