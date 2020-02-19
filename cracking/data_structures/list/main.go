package main

import (
	"fmt"
	"os"
)

// Node is one element of the list
type Node struct {
	val  int
	next *Node
}

// LinkedList refers the head and the tail of the list
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// Add appends a element at the end of the list
func (list *LinkedList) Add(n *Node) {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}

	if list.tail == nil {
		list.head = n
	} else {
		list.tail.next = n
	}
	list.tail = n
	list.length++
}

// Remove deletes a element of index of the argument
func (list *LinkedList) Remove(k int) {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}
	if k < 0 {
		fmt.Println("Remove: the index must be positive.")
		os.Exit(1)
	} else if k > list.length-1 {
		fmt.Println("Remove: the index is out of the range.")
		os.Exit(1)
	}

	now := list.head
	if now == nil {
		fmt.Println("Remove: list is empty.")
		os.Exit(1)
	}

	if k == 0 {
		list.head = now.next
	}

	for i := 0; i < k-1; i++ {
		now = now.next
	}
	if k == list.length-1 {
		list.tail = now
		now.next = nil
	} else {
		now.next = now.next.next
	}

	list.length--
}

// Show prints the all the elements from the head to the tail
func (list *LinkedList) Show() {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}

	now := list.head
	for now != nil {
		fmt.Printf("%d ", now.val)
		now = now.next
	}
	fmt.Println()
}

func genLinkedList(num int) *LinkedList {
	list := &LinkedList{nil, nil, 0}
	for i := 0; i < num; i++ {
		n := &Node{i, nil}
		list.Add(n)
	}
	return list
}

func main() {
	list := genLinkedList(10)
	list.Show()

	list.Add(&Node{10, nil})
	list.Show()

	list.Remove(1)
	list.Show()
}
