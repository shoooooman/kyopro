package myheap

// IntHeap is heap of int values
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Push adds an element to heap
func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(int))
}

// Pop removes min value from heap and return it
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Top returns min value of heap
func (h *IntHeap) Top() int {
	return (*h)[0]
}

// func main() {
// 	q := &IntHeap{1, 3, 2}
// 	heap.Push(q, 4)
// 	heap.Init(q)
// 	heap.Push(q, 6)
// 	heap.Push(q, 5)
// 	for len(*q) > 0 {
// 		fmt.Println(q.Top(), heap.Pop(q))
// 	}
// }
