package unionfind

// UnionFind is union find structure
type UnionFind struct {
	Par   []int
	Ranks []int
	Sizes []int
}

// NewUnionFind generates a new union find
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{make([]int, n), make([]int, n), make([]int, n)}
	for i := range uf.Par {
		uf.Par[i] = i
		uf.Ranks[i] = 0
		uf.Sizes[i] = 1
	}
	return uf
}

// Root searches the root of the arguement
func (uf *UnionFind) Root(x int) int {
	if x == uf.Par[x] {
		return x
	}
	uf.Par[x] = uf.Root(uf.Par[x]) // 経路圧縮
	return uf.Par[x]
}

// Unite unites two trees by adding one root to another
func (uf *UnionFind) Unite(x, y int) {
	rx := uf.Root(x)
	ry := uf.Root(y)
	if rx == ry {
		return
	}
	if uf.Ranks[rx] < uf.Ranks[ry] {
		uf.Par[rx] = ry
		uf.Sizes[ry] += uf.Sizes[rx]
	} else {
		uf.Par[ry] = rx
		uf.Sizes[rx] += uf.Sizes[ry]
		if uf.Ranks[rx] == uf.Ranks[ry] {
			uf.Ranks[rx]++
		}
	}
}

// Same returns true when the two are in the same group
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

// Size returns size of a group where x belongs
func (uf *UnionFind) Size(x int) int {
	return uf.Sizes[uf.Root(x)]
}

// func main() {
// 	n := 5
// 	uf := NewUnionFind(n)
//
// 	fmt.Println(uf.Same(0, 1))
// 	uf.Unite(0, 1)
// 	fmt.Println(uf.Same(0, 1))
// 	uf.Unite(2, 3)
// 	fmt.Println(uf.Same(0, 2))
// 	uf.Unite(0, 3)
// 	fmt.Println(uf.Same(0, 2))
// 	fmt.Println(uf.Size(0))
// }
