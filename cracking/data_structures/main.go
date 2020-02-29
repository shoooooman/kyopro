package main

import (
	"fmt"
	"os"

	"github.com/shoooooman/kyopro/cracking/data_structures/mybst"
	"github.com/shoooooman/kyopro/cracking/data_structures/mygraph"
	"github.com/shoooooman/kyopro/cracking/data_structures/mylist"
	"github.com/shoooooman/kyopro/cracking/data_structures/myqueue"
	"github.com/shoooooman/kyopro/cracking/data_structures/mystack"
)

func main() {
	/* ----------mylist---------- */
	list := mylist.GenLinkedList(10)
	list.Show()

	list.Add(&mylist.Node{10, nil})
	list.Show()

	list.Remove(1)
	list.Show()

	list.InsertAfter(list.Length-1, &mylist.Node{11, nil})
	list.Show()

	list.InsertBefore(0, &mylist.Node{12, nil})
	list.Show()

	fmt.Println()

	/* ----------mystack---------- */
	s := mystack.NewStack()

	s.Push(0)
	s.Push(1)
	fmt.Println(*s)

	if v, err := s.Pop(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	if v, err := s.Peek(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(*s)
	fmt.Println(s.Size())
	fmt.Println(s.Empty())
	s.Pop()
	s.Pop()
	fmt.Println(s.Empty())

	fmt.Println()

	/* ----------myqueue---------- */
	q := myqueue.NewQueue()

	q.Push(0)
	q.Push("foo")
	fmt.Println(*q)

	if v, err := q.Pop(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Errorf("queue is empty")
	}
	fmt.Println(*q)

	fmt.Println()

	/* ----------mygraph---------- */
	graph := mygraph.NewGraph(5)
	graph.Show()

	start := graph.Nodes[0]
	goal := graph.Nodes[3]
	path, err := graph.FindPath(start, goal)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", start.Val)
		for _, node := range path {
			fmt.Printf("->%v", node.Val)
		}
		fmt.Println()
	}

	fmt.Println()

	/* ----------mybst---------- */
	values := []int{0, 1, 2, 3, 4, 5, 6}
	// values := []int{0, 1, 2, 3, 4, 5, 6, 7}
	tree := mybst.GenBST(values)
	mybst.PrintTree(tree)
	mybst.PrintValues(tree.Root)
	fmt.Println()
}
