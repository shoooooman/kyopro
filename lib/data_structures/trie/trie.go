package trie

// Node is a node of Trie
type Node struct {
	Val      rune
	Children []*Node
}

// Trie has its root node
type Trie struct {
	Root *Node
}

// return true if r is included in nodes
func hasChild(r rune, nodes []*Node) (*Node, bool) {
	for _, n := range nodes {
		if n.Val == r {
			return n, true
		}
	}
	return nil, false
}

// Insert inserts a node with the word to trie
func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}

	root := t.Root
	for _, r := range word {
		if n, ok := hasChild(r, root.Children); ok {
			root = n
		} else {
			newNode := &Node{r, nil}
			root.Children = append(root.Children, newNode)
			root = newNode
		}
	}
	leaf := &Node{'*', nil}
	root.Children = append(root.Children, leaf)
}

// Search returns true if word is in trie
func (t *Trie) Search(word string) bool {
	root := t.Root
	for _, r := range word {
		if n, ok := hasChild(r, root.Children); ok {
			root = n
		} else {
			return false
		}
	}
	if _, ok := hasChild('*', root.Children); ok {
		return true
	}
	return false
}

// NewTrie generates a new trie with dict nodes
func NewTrie(dict []string) *Trie {
	t := &Trie{new(Node)}
	for _, word := range dict {
		t.Insert(word)
	}
	return t
}

// func main() {
// 	words := []string{"foo", "bar", "hoge", "buzz"}
// 	t := NewTrie(words)
//
// 	fmt.Println(t.Search("foo"))
// 	fmt.Println(t.Search("fao"))
// 	fmt.Println(t.Search("hoge"))
// 	fmt.Println(t.Search("hogee"))
// }
