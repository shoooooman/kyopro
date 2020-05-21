package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	maxVal   = 200
	maxSize  = 100
	numTests = 1000
)

func randomInts() []int {
	n := rand.Intn(maxSize) + 1
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(maxVal)
	}
	return nums
}

func TestQuickSelect(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numTests; i++ {
		nums := randomInts()
		k := rand.Intn(len(nums))
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			t.Parallel()

			actual := quickSelect(nums, 0, len(nums)-1, k)
			sort.Ints(nums)
			expected := nums[k]
			if actual != expected {
				t.Errorf("expected %d, actual %d", expected, actual)
			}
		})
	}
}
