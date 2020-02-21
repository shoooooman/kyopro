package main

import "github.com/shoooooman/kyopro/cracking/data_structures/mylist"

func main() {
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
}
