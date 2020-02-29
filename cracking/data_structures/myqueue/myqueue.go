package myqueue

import "fmt"

// Queue is for bfs
type Queue struct {
	data []interface{}
}

// Push adds a value to queue
func (q *Queue) Push(v interface{}) {
	q.data = append(q.data, v)
}

// Pop removes the last element and return it
func (q *Queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("queue is empty")
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v, nil
}

// Empty returns true if queue is empty
func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

// NewQueue generates a new queue
func NewQueue() *Queue {
	return &Queue{make([]interface{}, 0)}
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
