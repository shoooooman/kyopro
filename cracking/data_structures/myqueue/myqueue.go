package myqueue

import "fmt"

// Queue is for bfs
type Queue struct {
	data []int
}

// Push adds a value to queue
func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

// Pop removes the last element and return it
func (q *Queue) Pop() (int, error) {
	if q.Empty() {
		return 0, fmt.Errorf("queue is empty")
	}
	p := q.data[0]
	q.data = q.data[1:]
	return p, nil
}

// Empty returns true if queue is empty
func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

// NewQueue generates a new queue
func NewQueue() *Queue {
	return &Queue{make([]int, 0)}
}

// func main() {
// 	q := NewQueue()
//
// 	q.Push(0)
// 	q.Push(1)
// 	fmt.Println(q)
//
// 	if v, err := q.Pop(); err == nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Errorf("queue is empty")
// 	}
// 	fmt.Println(q)
// }
