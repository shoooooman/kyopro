package mystack

import (
	"fmt"
)

// Stack is wrapper of []int
type Stack []int

// Push adds an element
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop removes the top element and return it
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Peek returns the top value
func (s *Stack) Peek() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

// Size returns the length of stack
func (s *Stack) Size() int {
	return len(*s)
}

// Empty returns true when stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// NewStack generates stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

// func main() {
// 	s := NewStack()
//
// 	s.Push(0)
// 	s.Push(1)
// 	fmt.Println(*s) // [0 1]
//
// 	fmt.Println(s.Size())  // 2
// 	fmt.Println(s.Empty()) // false
//
// 	if v, err := s.Pop(); err == nil {
// 		fmt.Println(v) // 1
// 	} else {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
//
// 	fmt.Println(*s) // [0]
//
// 	if v, err := s.Peek(); err == nil {
// 		fmt.Println(v) // 0
// 	} else {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
//
// 	s.Pop()
// 	fmt.Println(s.Empty()) // true
// }
