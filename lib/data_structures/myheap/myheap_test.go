package myheap

import (
	"container/heap"
	"log"
	"testing"
)

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

func TestMyHeap(t *testing.T) {
	q := &IntHeap{1, 3, 4}
	Init(q)
	Push(q, 5)
	Push(q, 2)
	assertInts(*q, []int{1, 2, 4, 5, 3}, t)
	Remove(q, 2)
	assertInts(*q, []int{1, 2, 3, 5}, t)
	(*q)[2] = 0
	Fix(q, 2)
	assertInts(*q, []int{0, 2, 1, 5}, t)
	Push(q, 3)
	assertInts(*q, []int{0, 2, 1, 5, 3}, t)
	for len(*q) > 0 {
		top := (*q)[0]
		if heap.Pop(q) != top {
			log.Fatal("Pop error")
		}
	}
}

func TestSort(t *testing.T) {
	q := &IntHeap{1, 3, 2, 5, 4}
	Init(q)
	Sort(q)

	expect := []int{5, 4, 3, 2, 1}
	assertInts(*q, expect, t)
}

func assertInts(result, expect []int, t *testing.T) {
	if len(result) != len(expect) {
		t.Fatal("length are different")
	}

	for i := range result {
		if result[i] != expect[i] {
			log.Fatal("different ints")
		}
	}
}
