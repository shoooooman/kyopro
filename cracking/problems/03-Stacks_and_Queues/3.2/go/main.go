package main

import (
	"fmt"
	"os"
)

type data struct {
	val int
	min int
}

// Stack is wrapper of []data
type Stack []data

// Push adds an element
func (s *Stack) Push(v int) {
	min, err := s.Min()
	if err == nil && v < min {
		min = v
	}
	d := data{v, min}
	*s = append(*s, d)
}

// Pop removes the top element and return it
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	d := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return d.val, nil
}

// Peek returns the top value
func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1].val, nil
}

// Min returns a minimum value in stack
func (s *Stack) Min() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1].min, nil
}

// Size returns the length of stack
func (s *Stack) Size() int {
	return len(*s)
}

// Empty returns true when stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// GenStack generates stack
func GenStack() *Stack {
	ss := Stack(make([]data, 0))
	return &ss
}

func main() {
	s := GenStack()

	s.Push(0)
	s.Push(-1)
	s.Push(1)
	s.Push(2)
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

	if v, err := s.Min(); err == nil {
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
	s.Pop()
	fmt.Println(s.Empty())
}
