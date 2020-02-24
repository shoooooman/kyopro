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

// GenStack generates stack
func GenStack() *Stack {
	ss := Stack(make([]int, 0))
	return &ss
}

// func main() {
// 	ss := Stack(make([]int, 0))
// 	s := &ss
//
// 	s.Push(0)
// 	s.Push(1)
// 	fmt.Println(*s)
//
// 	if v, err := s.Pop(); err == nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
//
// 	if v, err := s.Peek(); err == nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
//
// 	fmt.Println(*s)
// 	fmt.Println(s.Size())
// 	fmt.Println(s.Empty())
// 	s.Pop()
// 	s.Pop()
// 	fmt.Println(s.Empty())
// }
