package main

import (
	"fmt"
	"math"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// solution1
// 葉ノードから再帰的に確認(button up)
// 最大値、最小値を保存
func isBST(root *mytree.Node) (int, int, bool) {
	if root == nil {
		return math.MinInt64, math.MaxInt64, true
	}

	lmax, lmin, lok := isBST(root.Left)
	if !lok {
		return -1, -1, false
	}

	rmax, rmin, rok := isBST(root.Right)
	if !rok {
		return -1, -1, false
	}

	v := root.Val.(int)
	fmt.Println(lmax, v, rmin)
	if v >= lmax && v <= rmin {
		vmax := v
		if rmax != math.MaxInt64 {
			vmax = max(v, rmax)
		}
		vmin := v
		if lmin != math.MinInt64 {
			vmin = min(v, lmin)
		}
		return vmax, vmin, true
	}
	return -1, -1, false
}

func helper(root *mytree.Node, minv, maxv int) bool {
	if root == nil {
		return true
	}

	v := root.Val.(int)
	if (minv != -1 && v < minv) || (maxv != -1 && v > maxv) {
		return false
	}

	if helper(root.Left, minv, v) && helper(root.Right, v, maxv) {
		return true
	}
	return false
}

// solution2
// 根ノードから再帰的に確認(top down)
func isBST2(root *mytree.Node) bool {
	return helper(root, -1, -1)
}

func main() {
	data := map[int][]int{
		5: {3, 7},
		3: {2, 4},
		7: {6, 8},
		2: {1, -1},

		// 0: {-1, -1},

		// 20: {10, 30},
		// 10: {-1, 25},

		// 20: {10, 30},
		// 10: {5, 15},
		// 5:  {3, 7},
		// 15: {-1, 17},
	}
	nonbst, _ := mytree.GenTree(data, 5)
	// _, _, ans := isBST(nonbst.Root)
	ans := isBST2(nonbst.Root)
	fmt.Println(ans)
}
