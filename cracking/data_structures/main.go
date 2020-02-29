package main

import (
	"fmt"
	"os"

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

	/* ----------myqueue---------- */
	q := myqueue.NewQueue()

	q.Push(0)
	q.Push(1)
	fmt.Println(*q)

	if v, err := q.Pop(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Errorf("queue is empty")
	}
	fmt.Println(*q)
}
