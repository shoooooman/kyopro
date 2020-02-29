package main

import (
	"fmt"
	"m/mylist"
	"os"
)

/* hash: O(N) */
func solve(list *mylist.LinkedList) *mylist.LinkedList {
	now := list.Head
	if now == nil {
		fmt.Println("list is empty.")
		os.Exit(1)
	}

	index := 0
	checker := make(map[int]bool)
	for now != nil {
		v := now.Val
		if checker[v] {
			list.Remove(index)
		} else {
			checker[v] = true
			index++
		}
		now = now.Next
	}
	return list
}

/* runner: O(N^2) */
func solve2(list *mylist.LinkedList) *mylist.LinkedList {
	now := list.Head
	if now == nil {
		fmt.Println("list is empty.")
		os.Exit(1)
	}

	i := 0
	for now != nil {
		v := now.Val
		j := i + 1
		checker := now.Next
		for checker != nil {
			if checker.Val == v {
				list.Remove(j)
			} else {
				j++
			}
			checker = checker.Next
		}
		now = now.Next
		i++
	}
	return list
}

func main() {
	list := mylist.GenLinkedList(10)
	list.Show()

	list = solve(list)
	// list = solve2(list)
	list.Show()
}
