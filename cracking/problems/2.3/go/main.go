package main

import (
	"fmt"
	"os"

	"github.com/shoooooman/kyopro/cracking/problems/2.3/go/mylist"
)

func exactRemove(node *mylist.Node) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}

func main() {
	const n = 10
	list := mylist.GenLinkedList(10)
	list.Show()

	var k int
	fmt.Scan(&k)
	if k == n-1 {
		fmt.Println("The removed element must be other than the last.")
		os.Exit(1)
	}

	runner := list.Head
	for i := 0; i < k; i++ {
		runner = runner.Next
	}
	exactRemove(runner)
	list.Show()
}
