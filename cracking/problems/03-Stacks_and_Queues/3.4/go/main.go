package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/cracking/data_structures/mystack"
)

// MyQueue is a queue which is implemented with two stacks
type MyQueue struct {
	StackNew *mystack.Stack
	StackOld *mystack.Stack
}

// Add adds a new value
func (q *MyQueue) Add(v int) {
	q.StackNew.Push(v)
}

// Shift moves all elements from StackNew to StackOld
func (q *MyQueue) Shift() {
	sn := q.StackNew
	so := q.StackOld
	for !sn.Empty() {
		v, _ := sn.Pop()
		so.Push(v)
	}
}

// Remove removes the last element
func (q *MyQueue) Remove() (int, error) {
	if q.StackOld.Empty() {
		q.Shift()
	}
	return q.StackOld.Pop()
}

func (q *MyQueue) String() string {
	return fmt.Sprintf("%v %v", *q.StackNew, *q.StackOld)

}

func main() {
	sn := mystack.NewStack()
	so := mystack.NewStack()
	q := &MyQueue{sn, so}

	q.Add(0)
	q.Add(1)
	fmt.Println(q)
	if v, err := q.Remove(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	fmt.Println(q)

	q.Add(2)
	q.Add(3)
	fmt.Println(q)

	if v, err := q.Remove(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	fmt.Println(q)

	if v, err := q.Remove(); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	fmt.Println(q)
}
