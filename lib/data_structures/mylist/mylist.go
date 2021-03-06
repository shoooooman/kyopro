package mylist

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Node is one element of the list
type Node struct {
	Val  int
	Next *Node
}

// LinkedList refers the head and the tail of the list
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// Add appends a element at the end of the list
func (list *LinkedList) Add(n *Node) {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}

	if list.Tail == nil {
		list.Head = n
	} else {
		list.Tail.Next = n
	}
	list.Tail = n
	list.Length++
}

// InsertAfter inserts a new node n after the node indexed k
func (list *LinkedList) InsertAfter(k int, n *Node) *Node {
	if list == nil {
		fmt.Println("list is nil.")
		os.Exit(1)
	}
	if k < 0 {
		fmt.Println("InsertAfter: the index must be positive.")
		os.Exit(1)
	} else if k > list.Length-1 {
		fmt.Println("InsertAfter: the index is out of the range.")
		os.Exit(1)
	}

	now := list.Head
	if now == nil {
		fmt.Println("InsertAfter: list is empty.")
		os.Exit(1)
	}

	if k == list.Length-1 {
		list.Tail.Next = n
		list.Tail = n
		list.Length++
		return n
	}

	for i := 0; i < k; i++ {
		now = now.Next
	}
	n.Next = now.Next
	now.Next = n
	list.Length++
	return n
}

// InsertBefore inserts a new node n before the node indexed k
func (list *LinkedList) InsertBefore(k int, n *Node) *Node {
	if list == nil {
		fmt.Println("list is nil.")
		os.Exit(1)
	}
	if k < 0 {
		fmt.Println("InsertAfter: the index must be positive.")
		os.Exit(1)
	} else if k > list.Length-1 {
		fmt.Println("InsertAfter: the index is out of the range.")
		os.Exit(1)
	}

	now := list.Head
	if now == nil {
		fmt.Println("InsertAfter: list is empty.")
		os.Exit(1)
	}

	if k == 0 {
		n.Next = list.Head
		list.Head = n
		list.Length++
		return n
	}

	for i := 0; i < k-1; i++ {
		now = now.Next
	}
	n.Next = now.Next
	now.Next = n
	list.Length++
	return n
}

// Remove deletes a element of index of the argument
func (list *LinkedList) Remove(k int) int {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}
	if k < 0 {
		fmt.Println("Remove: the index must be positive.")
		os.Exit(1)
	} else if k > list.Length-1 {
		fmt.Println("Remove: the index is out of the range.")
		os.Exit(1)
	}

	now := list.Head
	if now == nil {
		fmt.Println("Remove: list is empty.")
		os.Exit(1)
	}

	if k == 0 {
		list.Head = now.Next
		list.Length--
		return now.Val
	}

	for i := 0; i < k-1; i++ {
		now = now.Next
	}
	rv := now.Next.Val
	if k == list.Length-1 {
		list.Tail = now
		now.Next = nil
	} else {
		now.Next = now.Next.Next
	}

	list.Length--
	return rv
}

// Show prints the all the elements from the head to the tail
func (list *LinkedList) Show() {
	if list == nil {
		fmt.Println("list is nil")
		os.Exit(1)
	}

	now := list.Head
	for now != nil {
		fmt.Printf("%d ", now.Val)
		now = now.Next
	}
	fmt.Println()
}

// GenLinkedList generates an example list
func GenLinkedList(num int) *LinkedList {
	list := &LinkedList{nil, nil, 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		// n := &Node{i, nil}
		n := &Node{rand.Intn(num), nil}
		list.Add(n)
	}
	return list
}

// func main() {
// 	list := GenLinkedList(10)
// 	list.Show()
//
// 	list.Add(&Node{10, nil})
// 	list.Show()
//
// 	list.Remove(1)
// 	list.Show()
//
// 	list.InsertAfter(list.Length-1, &Node{11, nil})
// 	list.Show()
//
// 	list.InsertBefore(0, &Node{12, nil})
// 	list.Show()
// }
