package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mytree"
)

/* solution1 */
// 再帰的に部分木がBSTであるかを確認
// 左の部分木の最大値 <= root <= 右の部分木の最小値 -> BST
// func isBST(node *mytree.Node) (bool, int, int) {
// 	if node == nil {
// 		return true, -1, -1
// 	}
//
// 	lbst, lmax, lmin := isBST(node.Left)
// 	if !lbst {
// 		return false, -1, -1
// 	}
//
// 	rbst, rmax, rmin := isBST(node.Right)
// 	fmt.Println(lmax, node.Val.(int), rmin)
// 	if !rbst ||
// 		!((lmax == -1 || lmax <= node.Val.(int)) && (rmin == -1 || node.Val.(int) <= rmin)) {
// 		return false, -1, -1
// 	}
//
// 	// 葉の場合
// 	if lmax == -1 && rmax == -1 {
// 		return true, node.Val.(int), node.Val.(int)
// 	}
// 	// 左の子が無い場合
// 	if lmax == -1 {
// 		// return true, node.Val.(int), rmin
// 		return true, rmax, node.Val.(int)
// 	}
// 	// 右の子が無い場合
// 	if rmax == -1 {
// 		// return true, lmax, node.Val.(int)
// 		return true, node.Val.(int), lmin
// 	}
// 	return true, rmax, lmin
// }

/* solution2 */
func isBST(node *mytree.Node, min, max int) bool {
	if node == nil {
		return true
	}
	v := node.Val.(int)
	if (min != -1 && v <= min) || (max != -1 && v > max) {
		return false
	}
	if !isBST(node.Left, min, v) || !isBST(node.Right, v, max) {
		return false
	}
	return true
}

func main() {
	data := map[int][]int{
		8:  []int{4, 10},
		4:  []int{2, 6},
		10: []int{-1, 20},
	}
	tree, _ := mytree.GenTree(data, 8)
	tree.Print()

	// if ok, _, _ := isBST(tree.Root); ok {
	if ok := isBST(tree.Root, -1, -1); ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
