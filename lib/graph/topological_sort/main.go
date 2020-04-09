package main

import "fmt"

type dependency struct {
	Child  interface{}
	Parent interface{}
}

type node struct {
	Name     interface{}
	Children []*node
	InBound  int // 入次数
}

func topologicalSort(nodes map[interface{}]*node, dependencies []dependency) []*node {
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

	ans := make([]*node, 0)
	visited := make(map[interface{}]bool)
	for len(tops) != 0 {
		nextTops := make([]*node, 0)
		for _, top := range tops {
			ans = append(ans, top)
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
	if len(visited) != len(nodes) {
		return nil
	}

	return ans
}

func createNodes(data []interface{}) map[interface{}]*node {
	nodes := make(map[interface{}]*node)
	for _, v := range data {
		nodes[v] = &node{v, nil, 0}
	}
	return nodes
}

func main() {
	data := []interface{}{"a", "b", "c", "d", "e", "f"}
	nodes := createNodes(data)

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

		// circle
		// dependency{"b", "a"},
		// dependency{"c", "b"},
		// dependency{"a", "c"},
	}
	rst := topologicalSort(nodes, dependencies)
	if rst == nil {
		fmt.Println("There exists a circle.")
		return
	}

	for _, n := range rst {
		fmt.Printf("%v ", n.Name)
	}
	fmt.Println()
}
