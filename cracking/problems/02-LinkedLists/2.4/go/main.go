package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/lib/data_structures/mylist"
)

/* xより小さい数は先頭から挿入していく
 * 挿入する場所を挿入するたびにずらしていく
 */
func main() {
	const n = 10
	list := mylist.GenLinkedList(n)
	list.Show()

	var x int
	fmt.Scan(&x)

	halfTop := 0
	runner := list.Head
	count := 0
	for runner != nil {
		if runner.Val < x {
			rv := runner.Val
			list.Remove(count)
			fmt.Printf("Removed %d at %d\n", rv, count)
			list.InsertBefore(halfTop, &mylist.Node{rv, nil})
			halfTop++
			fmt.Println(halfTop)
			list.Show()
		}
		runner = runner.Next
		count++
	}
	list.Show()
}
