package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		query    int
		expected int
	}{
		{"test1", []int{1, 2, 3}, 1, 0},
		{"test2", []int{1, 2, 3}, 2, 1},
		{"test3", []int{1, 2, 3}, 0, 0},
		{"test4", []int{1, 2, 3, 4}, 1, 0},
		{"test5", []int{1, 2, 3, 4}, 2, 1},
		{"test6", []int{1}, 1, 0},
		{"test7", []int{1}, 0, 0},
		{"test8", []int{}, 0, 0},
		{"test9", []int{1, 2}, 1, 0},
		{"test10", []int{1, 2}, 2, 1},
		{"test11", []int{1, 2}, 0, 0},
		{"test12", []int{1, 2, 2, 3, 5, 6}, 2, 1},
		{"test13", []int{1, 2, 2, 2, 5, 6}, 2, 1},
		{"test14", []int{1, 2, 2, 2, 5, 6}, 3, 4},
		{"test15", []int{1, 2, 2, 2, 5, 6}, 1, 0},
		{"test16", []int{1, 2, 2, 2, 5, 6}, 6, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := binarySearch(tt.nums, tt.query)
			if actual != tt.expected {
				t.Errorf("given(%v): expected %v, actual %v", tt.nums, tt.expected, actual)
			}
		})
	}
}
