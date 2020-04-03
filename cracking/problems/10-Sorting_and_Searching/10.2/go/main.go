package main

import (
	"fmt"
	"sort"
	"strings"
)

type strWithID struct {
	Str string
	ID  int
}

type strWithIDs []strWithID

func (s strWithIDs) Less(i, j int) bool {
	return s[i].Str < s[j].Str
}

func (s strWithIDs) Len() int {
	return len(s)
}

func (s strWithIDs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func convert(strs []string) strWithIDs {
	strIDs := make(strWithIDs, len(strs))
	for i, str := range strs {
		strIDs[i] = strWithID{sortString(str), i}
	}
	return strIDs
}

func solve(strs []string) []string {
	strIDs := convert(strs)
	sort.Sort(strIDs)
	sorted := make([]string, len(strs))
	for i, strID := range strIDs {
		sorted[i] = strs[strID.ID]
	}
	return sorted
}

func main() {
	strs := []string{"abc", "abb", "acb", "acc", "bab"}
	strs = solve(strs)
	fmt.Println(strs)
}
