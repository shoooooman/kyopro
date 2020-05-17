package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mybst"
)

func weaveLists(left, right, prefix []int, weaved *[][]int) {
	if len(left) == 0 || len(right) == 0 {
		// weavedにはprefixのポインタが格納されるので
		// コピーしないと呼び出し元でprefixを変更したときに値が壊れる
		cp := make([]int, len(prefix))
		copy(cp, prefix)
		cp = append(cp, left...)
		cp = append(cp, right...)
		*weaved = append(*weaved, cp)
		return
	}

	leftHead := left[0]
	prefix = append(prefix, leftHead)
	weaveLists(left[1:], right, prefix, weaved)
	prefix = prefix[:len(prefix)-1]

	rightHead := right[0]
	prefix = append(prefix, rightHead)
	weaveLists(left, right[1:], prefix, weaved)
	prefix = prefix[:len(prefix)-1]
}

func doubleDFS(node *mybst.Node) [][]int {
	result := make([][]int, 0)

	if node == nil {
		// 要素を入れておくことで下の
		// left, rightについてのループが回る
		result = append(result, make([]int, 0))
		return result
	}

	prefix := []int{node.Val}

	left := doubleDFS(node.Left)
	right := doubleDFS(node.Right)
	for _, ll := range left {
		for _, rl := range right {
			weaved := make([][]int, 0)
			weaveLists(ll, rl, prefix, &weaved)
			result = append(result, weaved...)
		}
	}
	return result
}

func solve(bst *mybst.Tree) [][]int {
	return doubleDFS(bst.Root)
}

func main() {
	// values := []int{1, 2, 3, 4, 5, 6, 7}
	// values := []int{1, 2, 3}
	values := []int{1, 2, 3, 4}
	bst := mybst.GenBST(values)
	bst.PrintTree()
	mybst.PrintValues(bst.Root)
	fmt.Println()

	arrays := solve(bst)
	fmt.Println(arrays)
}
