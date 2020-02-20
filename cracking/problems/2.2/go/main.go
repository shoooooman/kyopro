package main

import (
	"fmt"
	"m/mylist"
	"os"
)

/* two pointers */
func solve(list *mylist.LinkedList, k int) {
	runner := list.Head
	chaser := list.Head
	for i := 0; i < k; i++ {
		runner = runner.Next
	}
	for runner.Next != nil {
		runner = runner.Next
		chaser = chaser.Next
	}
	vk := chaser.Val
	fmt.Println(vk)
}

/* recursive */
func solve2(list *mylist.LinkedList, k int) {
	p, _ := searchIth(list.Head, k, 0)
	fmt.Println(p.Val)
}

func searchIth(head *mylist.Node, k int, count int) (*mylist.Node, int) {
	if head == nil {
		return nil, 0
	}
	nd, c := searchIth(head.Next, k, count)
	c++
	if c-1 == k {
		return head, c
	}
	return nd, c
}

func main() {
	const n = 10
	list := mylist.GenLinkedList(n)
	list.Show()

	var k int
	fmt.Scan(&k)

	if k >= n {
		fmt.Println("k must be smaller than n.")
		os.Exit(1)
	}

	solve(list, k)
	// solve2(list, k)
}
