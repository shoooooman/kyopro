package myheap

// Heap is a interface to use methods for heap.
// Any type that implements this interface
// can be used as a heap.
type Heap interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
}

// Init builds a heap from the argument
func Init(h Heap) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push adds an element to a heap
func Push(h Heap, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1) // move the added element to right place
}

// Pop removes the top element from a heap and return it
func Pop(h Heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)   // swap the top one and the last one
	down(h, 0, n)  // move the top one (former the last one) to right place (ignore the last element)
	return h.Pop() // remove the last one and return it
}

// Remove removes the ith element from a heap and return it
func Remove(h Heap, i int) interface{} {
	n := h.Len() - 1
	if i != n {
		h.Swap(i, n) // move removing element to the tail
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix rebuilds a heap after the ith element has changed its value.
func Fix(h Heap, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// Sort sorts the elements according to the Less method by heap sort
func Sort(h Heap) {
	for i := h.Len() - 1; i > 0; i-- {
		h.Swap(0, i)
		down(h, 0, i-1)
	}
}

// compare the ith element with its parent and swap them
// if the heap invariants is invalidated
func up(h Heap, i int) {
	for {
		parent := (i - 1) / 2
		if i == parent || !h.Less(i, parent) {
			break
		}
		h.Swap(parent, i)
		i = parent
	}
}

// compare the ith element with its children and swap it with smaller child
// if the heap invariants is invalidated
func down(h Heap, i0, n int) bool {
	i := i0
	for {
		left := 2*i + 1
		if left >= n || left < 0 { // left < 0 is saticefied when overflow
			break
		}
		smaller := left
		if right := left + 1; right < n && h.Less(right, left) {
			smaller = right
		}
		if !h.Less(smaller, i) {
			break
		}
		h.Swap(i, smaller)
		i = smaller
	}
	return i > i0 // true when h[i0] changed
}

/* recursion version */
// func up(h Heap, i int) {
// 	parent := (i - 1) / 2
// 	if i == parent || !h.Less(i, parent) { // i == parent when i is 0
// 		return
// 	}
// 	h.Swap(parent, i)
// 	up(h, parent)
// }
//
// func down(h Heap, i, n int) bool {
// 	left := 2*i + 1
// 	if left >= n || left < 0 {
// 		return false
// 	}
// 	smaller := left
// 	if right := left + 1; right < n && h.Less(right, left) {
// 		smaller = right
// 	}
// 	if h.Less(smaller, i) {
// 		h.Swap(i, smaller)
// 		down(h, smaller, n)
// 		return true
// 	}
// 	return false
// }
