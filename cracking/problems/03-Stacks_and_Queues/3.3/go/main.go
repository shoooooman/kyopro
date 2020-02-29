package main

import (
	"fmt"

	"github.com/shoooooman/kyopro/cracking/data_structures/mystack"
)

// SetOfStacks is a set of stacks
type SetOfStacks struct {
	Stacks   []*mystack.Stack
	Capacity int
	NowStack int
}

// Push adds a new element
func (ss *SetOfStacks) Push(v int) {
	ns := ss.Stacks[ss.NowStack]
	if ns.Size() == ss.Capacity {
		newStack := mystack.NewStack()
		ss.Stacks = append(ss.Stacks, newStack)
		ns = newStack
		ss.NowStack++
	}
	ns.Push(v)
}

// Pop removes the top element and return it
func (ss *SetOfStacks) Pop() (int, error) {
	ns := ss.Stacks[ss.NowStack]
	if ns.Empty() {
		if ss.NowStack == 0 {
			return 0, fmt.Errorf("stack is empty")
		}
		ss.Stacks = ss.Stacks[:len(ss.Stacks)-1]
		ns = ss.Stacks[len(ss.Stacks)-1]
		ss.NowStack--
	}
	return ns.Pop()
}

// PopAt returns the top value of the argument stack
func (ss *SetOfStacks) PopAt(index int) (int, error) {
	if index < 0 || index > ss.NowStack {
		return 0, fmt.Errorf("PopAt: index is out of range")
	}
	is := ss.Stacks[index]
	return is.Pop()
}

func (ss *SetOfStacks) String() string {
	str := ""
	for i := 0; i <= ss.NowStack; i++ {
		str += fmt.Sprintf("%v ", *ss.Stacks[i])
	}
	return str
}

func main() {
	stacks := []*mystack.Stack{mystack.NewStack()}
	ss := &SetOfStacks{stacks, 5, 0}
	fmt.Println(ss)
	ss.Push(0)
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	fmt.Println(ss)
	if v, err := ss.Pop(); err == nil {
		fmt.Println(v)
	}
	if v, err := ss.Pop(); err == nil {
		fmt.Println(v)
	}

	ss.Push(6)
	ss.Push(7)
	fmt.Println(ss)

	if v, err := ss.PopAt(0); err == nil {
		fmt.Println(v)
	}
	fmt.Println(ss)
}
