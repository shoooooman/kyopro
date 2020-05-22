package trie

import "testing"

func TestTrie(t *testing.T) {
	var tests = []struct {
		name     string
		dict     []string
		query    string
		expected bool
	}{
		{"test1", []string{"foo", "bar"}, "foo", true},
		{"test2", []string{"foo", "bar"}, "for", false},
		{"test3", []string{"foo", "bar"}, "fooo", false},
		{"test4", []string{"foo", "bar"}, "fo", false},
		{"test5", []string{"foo", "fooo"}, "fooo", true},
		{"test6", []string{"foo", "fooo"}, "foo", true},
		{"test7", []string{"foo", "fooo"}, "fo", false},
		{"test8", []string{"foo", "fou"}, "fouo", false},
		{"test9", []string{"foo", "fou"}, "fo", false},
		{"test10", []string{}, "foo", false},
		{"test11", []string{""}, "", true},
		{"test12", []string{"foo", "foo"}, "foo", true},
	}

	for _, tt := range tests {
		tr := NewTrie(tt.dict)
		t.Run(tt.name, func(t *testing.T) {
			actual := tr.Search(tt.query)
			if actual != tt.expected {
				t.Errorf("dict(%v): expected %t, actual %t", tt.dict, tt.expected, actual)
			}
		})
	}
}
