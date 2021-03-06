package mybst

import (
	"testing"
)

func TestGenBST(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)
	bst.PrintTree()

	if bst.Root.Val != 3 {
		t.Fatal("gen error")
	}
	if bst.Root.Left.Val != 1 {
		t.Fatal("gen error")
	}
	if bst.Root.Right.Val != 5 {
		t.Fatal("gen error")
	}
}

func TestInsert01(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	inserted := bst.Insert(7)
	bst.PrintTree()
	if inserted.Parent.Val != 6 {
		t.Fatal("insert error")
	}

	inserted = bst.Insert(2)
	bst.PrintTree()
	if inserted.Parent.Val != 2 {
		t.Fatal("insert error")
	}
}

func TestInsert02(t *testing.T) {
	bst := Tree{nil}

	inserted := bst.Insert(0)
	bst.PrintTree()
	if inserted.Parent != nil {
		t.Fatal("insert error")
	}

	inserted = bst.Insert(1)
	bst.PrintTree()
	if inserted.Parent.Val != 0 {
		t.Fatal("insert error")
	}
}

func TestDelete01(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	bst.Delete(bst.Root.Left) // 1
	bst.PrintTree()
	if bst.Root.Left.Val != 2 {
		t.Fatal("delete error")
	}
	if bst.Root.Left.Left.Val != 0 {
		t.Fatal("delete error")
	}
	if bst.Root.Left.Right != nil {
		t.Fatal("delete error")
	}
}

func TestDelete02(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	bst.Delete(bst.Root.Left.Left) // 0
	bst.PrintTree()
	if bst.Root.Left.Val != 1 {
		t.Fatal("delete error")
	}
	if bst.Root.Left.Left != nil {
		t.Fatal("delete error")
	}
	if bst.Root.Left.Right.Val != 2 {
		t.Fatal("delete error")
	}
}

func TestDelete03(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	bst.Delete(bst.Root) // 3
	bst.PrintTree()
	if bst.Root.Val != 4 {
		t.Fatal("delete error")
	}
	if bst.Root.Right.Val != 5 {
		t.Fatal("delete error")
	}
	if bst.Root.Right.Left != nil {
		t.Fatal("delete error")
	}
	if bst.Root.Left.Val != 1 {
		t.Fatal("delete error")
	}
	if bst.Root.Right.Right.Val != 6 {
		t.Fatal("delete error")
	}
}

func TestFindNext(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	n1 := FindNext(bst.Root.Right)
	if n1.Val != 6 {
		t.Fatal("find next error")
	}

	n2 := FindNext(bst.Root.Left)
	if n2.Val != 2 {
		t.Fatal("find next error")
	}

	n3 := FindNext(bst.Root.Left.Right)
	if n3.Val != 3 {
		t.Fatal("find next error")
	}

	n4 := FindNext(bst.Root.Right.Left)
	if n4.Val != 5 {
		t.Fatal("find next error")
	}

	n5 := FindNext(bst.Root.Right.Right)
	if n5 != nil {
		t.Fatal("find next error")
	}

	n6 := FindNext(bst.Root.Right.Right.Right)
	if n6 != nil {
		t.Fatal("find next error")
	}
}

func TestFindPrev(t *testing.T) {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	bst := GenBST(values)

	n1 := FindPrev(bst.Root.Right)
	if n1.Val != 4 {
		t.Fatal("find next error")
	}

	n2 := FindPrev(bst.Root.Left)
	if n2.Val != 0 {
		t.Fatal("find next error")
	}

	n3 := FindPrev(bst.Root.Left.Right)
	if n3.Val != 1 {
		t.Fatal("find next error")
	}

	n4 := FindPrev(bst.Root.Right.Left)
	if n4.Val != 3 {
		t.Fatal("find next error")
	}

	n5 := FindPrev(bst.Root.Left.Left)
	if n5 != nil {
		t.Fatal("find next error")
	}

	n6 := FindPrev(bst.Root.Right.Right.Right)
	if n6 != nil {
		t.Fatal("find next error")
	}
}
