package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Greedy: search all magic indexes
func solve(a []int) []int {
	magics := make([]int, 0)
	for i, v := range a {
		if i < v {
			return magics
		}
		if i == v {
			magics = append(magics, i)
		}
	}
	return magics
}

// Binary search: search only one magic index
func solve2(a []int) int {
	left := 0
	right := len(a) - 1
	for left <= right {
		mid := (right - left) / 2
		if a[mid] == mid {
			return mid
		}
		if a[mid] < mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearch(a []int, left, right int) int {
	if left > right {
		return -1
	}

	mid := (right - left) / 2
	if a[mid] == mid {
		return mid
	}

	// left side
	newRight := int(math.Min(float64(mid-1), float64(a[mid])))
	leftMagic := binarySearch(a, left, newRight)
	if leftMagic != -1 {
		return leftMagic
	}

	// right side
	newLeft := int(math.Max(float64(mid-1), float64(a[mid])))
	rightMagic := binarySearch(a, newLeft, mid-1)
	if rightMagic != -1 {
		return rightMagic
	}

	return -1
}

// Binary search: search one magic index even when the elements overlap
func solve3(a []int) int {
	return binarySearch(a, 0, len(a)-1)
}

func getArray(n int) []int {
	a := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	exist := make(map[int]bool)
	for i := 0; i < n; i++ {
		for {
			rnd := rand.Intn(15)
			if !exist[rnd] {
				a[i] = rnd
				exist[rnd] = true
				break
			}
		}
	}
	sort.Ints(a)
	return a
}

func getOverlapArray(n int) []int {
	a := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(15)
	}
	sort.Ints(a)
	return a
}

func main() {
	const n = 10
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	fmt.Println(ids)

	a := getArray(n)
	fmt.Println(a)

	magics := solve(a)
	fmt.Println(magics)
	magic2 := solve2(a)
	fmt.Println(magic2)

	fmt.Println()

	fmt.Println(ids)
	b := getOverlapArray(n)
	fmt.Println(b)

	magic3 := solve3(b)
	fmt.Println(magic3)
}
