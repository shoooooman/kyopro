package main

import "fmt"

type dependency struct {
	Child  string
	Parent string
}

type node struct {
	Name     string
	Children []*node
	InBound  int
}

func solve(projects []string, dependencies []dependency) []string {
	nodes := make(map[string]*node)
	for _, project := range projects {
		nodes[project] = &node{project, nil, 0}
	}

	for _, dependency := range dependencies {
		parent := dependency.Parent
		child := dependency.Child
		nodes[parent].Children = append(nodes[parent].Children, nodes[child])
		nodes[child].InBound++
	}

	// 親を持たないノード
	tops := make([]*node, 0)
	for _, n := range nodes {
		if n.InBound == 0 {
			tops = append(tops, n)
		}
	}

	ans := make([]string, 0)
	visited := make(map[string]bool)
	for len(tops) != 0 {
		nextTops := make([]*node, 0)
		for _, top := range tops {
			ans = append(ans, top.Name)
			visited[top.Name] = true
			for _, child := range top.Children {
				child.InBound--
				if child.InBound == 0 {
					nextTops = append(nextTops, child)
				}
			}
		}
		tops = nextTops
	}

	// 閉路がある場合(topsに含まれることはないため訪問されない)
	if len(visited) != len(projects) {
		return nil
	}

	return ans
}

func main() {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := []dependency{
		dependency{"d", "a"},
		dependency{"b", "f"},
		dependency{"d", "b"},
		dependency{"a", "f"},
		dependency{"c", "d"},

		// dependency{"a", "f"},
		// dependency{"d", "a"},
		// dependency{"d", "b"},
		// dependency{"c", "d"},

		// dependency{"a", "f"},
		// dependency{"d", "a"},
		// dependency{"b", "a"},
		// dependency{"c", "d"},

		// dependency{"b", "a"},
		// dependency{"c", "b"},
		// dependency{"a", "c"},
	}
	rst := solve(projects, dependencies)
	fmt.Println(rst)
}
