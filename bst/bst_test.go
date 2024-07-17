package bst_test

import (
	"testing"

	"github.com/dungtl2003/data-structure/bst"
)

// https://www.google.com/url?sa=i&url=https%3A%2F%2Fbyjus.com%2Fgate%2Fbinary-search-trees-notes%2F&psig=AOvVaw0yA-ujUk0itrm-NiHiupUM&ust=1721216558895000&source=images&opi=89978449
func createSampleBst() *bst.BST[int] {
	tree := bst.New[int]()
	tree.Add(8)
	tree.Add(3)
	tree.Add(1)
	tree.Add(6)
	tree.Add(4)
	tree.Add(7)
	tree.Add(10)
	tree.Add(14)
	tree.Add(13)

	return tree
}

func TestBstPreOrderIterate(t *testing.T) {
	tree := createSampleBst()

	want := [9]int{8, 3, 1, 6, 4, 7, 10, 14, 13}
	result := tree.PreOrder()

	if want != ([9]int)(result) {
		t.Fatalf("tree.PreOrder() did not give the expected result, want %v, got %v", want, result)
	}
}

func TestBstInOrderIterate(t *testing.T) {
	tree := createSampleBst()

	want := [9]int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	result := tree.InOrder()

	if want != ([9]int)(result) {
		t.Fatalf("tree.InOrder() did not give the expected result, want %v, got %v", want, result)
	}
}

func TestBstPostOrderIterate(t *testing.T) {
	tree := createSampleBst()

	want := [9]int{1, 4, 7, 6, 3, 13, 14, 10, 8}
	result := tree.PostOrder()

	if want != ([9]int)(result) {
		t.Fatalf("tree.PostOrder() did not give the expected result, want %v, got %v", want, result)
	}
}

func TestBstContainsData(t *testing.T) {
	tree := bst.New[int]()
	tree.Add(12)
	tree.Add(13)
	tree.Add(10)

	if !tree.Contains(13) {
		t.Fatalf("tree.Contains(13) should return true")
	}
	if tree.Contains(3) {
		t.Fatalf("tree.Contains(3) should return false")
	}
}

func TestSizeOfBst(t *testing.T) {
	tree := createSampleBst()

	var size uint
	var want uint
	size = tree.Size()
	want = 9

	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}

	tree.Add(6) // 6 is already in tree, size should not change
	size = tree.Size()
	want = 9
	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}

	tree.Add(69) // 69 is not in the tree, size should increase
	size = tree.Size()
	want = 10
	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}

	tree.Add(6)
	tree.Add(6)
	tree.Remove(6, false) // there are 3 number 6, remove 1 should not change the size
	size = tree.Size()
	want = 10
	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}

	tree.Remove(6, true) // there are 2 number 6, but we ignore occurences, size should decrease
	size = tree.Size()
	want = 9
	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}

	tree.Init() // this should clear the tree
	size = tree.Size()
	want = 0
	if size != want {
		t.Fatalf("tree.Size() did not give the expected result, want %d, got %d", want, size)
	}
}
