package main

import "testing"

var tests = []struct {
	name         string
	expected     [][]interface{}
	all          []interface{}
	dependencies []dependency
}{
	{
		"normal1",
		[][]interface{}{[]interface{}{0, 1}},
		[]interface{}{0, 1},
		[]dependency{dependency{1, 0}},
	},
	{
		"normal2",
		[][]interface{}{[]interface{}{0, 1, 2}, []interface{}{0, 2, 1}},
		[]interface{}{0, 1, 2},
		[]dependency{dependency{1, 0}, dependency{2, 0}},
	},
	{
		"normal3",
		[][]interface{}{[]interface{}{0, 1, 2, 3}, []interface{}{0, 2, 1, 3}},
		[]interface{}{0, 1, 2, 3},
		[]dependency{dependency{1, 0}, dependency{2, 0}, dependency{3, 1}, dependency{3, 2}},
	},
	{
		"normal4",
		[][]interface{}{[]interface{}{0, 2, 1, 3}, []interface{}{0, 2, 3, 1}, []interface{}{2, 0, 1, 3}, []interface{}{2, 0, 3, 1}},
		[]interface{}{0, 1, 2, 3},
		[]dependency{dependency{1, 0}, dependency{3, 2}},
	},
	{
		"normal5",
		[][]interface{}{[]interface{}{0}},
		[]interface{}{0},
		[]dependency{},
	},
	{
		"cycle1",
		nil,
		[]interface{}{0, 1, 2},
		[]dependency{dependency{1, 0}, dependency{2, 1}, dependency{0, 2}},
	},
	{
		"cycle2",
		nil,
		[]interface{}{0, 1, 2, 3},
		[]dependency{dependency{1, 0}, dependency{2, 1}, dependency{0, 2}, dependency{2, 3}},
	},
	{
		"empty",
		nil,
		[]interface{}{},
		[]dependency{},
	},
}

// expectedの要素に1つでもactualが合致すればOK
func assertSlices(actual []interface{}, expected [][]interface{}, t *testing.T) {
	if len(actual) == 0 && len(expected) == 0 {
		return
	}

L_FOR:
	for _, e := range expected {
		if len(actual) != len(e) {
			continue L_FOR
		}

		for i := range actual {
			if actual[i] != e[i] {
				continue L_FOR
			}
		}
		t.Logf("match: %v", e)
		return
	}
	t.Fatalf("expected %v, actual %v", expected, actual)
}

// []*nodeをnodeのID列に変換
func convertNodesToIDs(ns []*node) []interface{} {
	ids := make([]interface{}, len(ns))
	for i, n := range ns {
		ids[i] = n.ID
	}
	return ids
}

func TestTopologicalSort(t *testing.T) {
	for _, tt := range tests {
		nodes := createNodes(tt.all)
		t.Run(tt.name, func(t *testing.T) {
			result := topologicalSort(nodes, tt.dependencies)
			actual := convertNodesToIDs(result)
			assertSlices(actual, tt.expected, t)
		})
	}
}
