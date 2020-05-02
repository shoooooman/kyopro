package main

import "testing"

func assertInts(expect, result []int, t *testing.T) {
	if len(expect) != len(result) {
		t.Errorf("different length")
	}
	for i := range result {
		if result[i] != expect[i] {
			t.Errorf("not expected element")
		}
	}
}

func TestNextPermutation01(t *testing.T) {
	a := []int{1, 2, 3}
	expect := []int{1, 3, 2}
	result := nextPermutation(a)
	assertInts(expect, result, t)
}

func TestNextPermutation02(t *testing.T) {
	a := []int{3, 2, 1}
	expect := []int{1, 2, 3}
	result := nextPermutation(a)
	assertInts(expect, result, t)
}

func TestNextPermutation03(t *testing.T) {
	a := []int{1, 2, 3}
	expect := []int{1, 2, 3}
	var result []int
	for i := 0; i < fact(len(a)); i++ {
		result = nextPermutation(a)
	}
	assertInts(expect, result, t)
}

func TestNextPermutation04(t *testing.T) {
	a := []int{1}
	expect := []int{1}
	var result []int
	for i := 0; i < fact(len(a)); i++ {
		result = nextPermutation(a)
	}
	assertInts(expect, result, t)
}

func TestNextPermutation05(t *testing.T) {
	a := []int{}
	expect := []int{}
	var result []int
	for i := 0; i < fact(len(a)); i++ {
		result = nextPermutation(a)
	}
	assertInts(expect, result, t)
}
