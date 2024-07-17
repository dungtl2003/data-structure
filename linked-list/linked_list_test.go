package linkedlist

import (
	"testing"
)

func TestAddDataToLinkedList(t *testing.T) {
	ll := New[int]()
	ll.Add(0, 2)
	ll.AddFirst(1)
	ll.AddLast(3)
	ll.AddFirst(0)
	ll.AddLast(4)
	ll.Add(5, 5)

	want := [6]int{0, 1, 2, 3, 4, 5}
	values := [6]int{}
	i := 0
	for n := ll.Head(); n != nil; {
		values[i] = n.Val

		i++
		n = n.Next()
	}

	if want != values {
		t.Fatalf("Unexpected result, want %v, got %v", want, values)
	}
}

func TestRemoveDataFromLinkedList(t *testing.T) {
	ll := New[int]()
	ll.AddLast(0)
	ll.AddLast(1)
	ll.AddLast(2)
	ll.AddLast(3)
	ll.AddLast(4)

	_, err := ll.Remove(5)
	if err == nil {
		t.Fatalf("ll.Remove(5) did not return IndexOutOfBound error")
	}

	ll.Remove(1)
	ll.Remove(3)
	ll.Remove(0)

	want := [2]int{2, 3}
	values := [2]int{}
	i := 0
	for n := ll.Head(); n != nil; {
		values[i] = n.Val

		i++
		n = n.Next()
	}

	if want != values {
		t.Fatalf("Unexpected result, want %v, got %v", want, values)
	}
}

func TestGetInformationFromLinkedList(t *testing.T) {
	ll := New[int]()
	ll.AddLast(0)
	ll.AddLast(4)
	ll.AddLast(2)
	ll.AddLast(3)
	ll.AddLast(4)
	ll.AddLast(5)

	vf := ll.GetFirst()
	if vf != 0 {
		t.Fatalf("ll.GetFirst() returned unexpected result, want 0, got %v", vf)
	}

	vl := ll.GetLast()
	if vl != 5 {
		t.Fatalf("ll.GetLast() returned unexpected result, want 5, got %v", vl)
	}

	_, err := ll.Get(6)
	if err == nil {
		t.Fatalf("ll.Get(6) did not return IndexOutOfBound error")
	}

	if !ll.Contains(3) {
		t.Fatalf("ll.Contains(3) should return true")
	}

	if ll.Contains(-3) {
		t.Fatalf("ll.Contains(-3) should return false")
	}

	pf := ll.IndexOf(4)
	if pf != 1 {
		t.Fatalf("ll.IndexOf(4) should return 1, not %d", pf)
	}

	pf = ll.IndexOf(7)
	if pf != -1 {
		t.Fatalf("ll.IndexOf(7) should return -1, not %d", pf)
	}

	pl := ll.LastIndexOf(4)
	if pl != 4 {
		t.Fatalf("ll.IndexOf(4) should return 4, not %d", pl)
	}

	pl = ll.LastIndexOf(7)
	if pl != -1 {
		t.Fatalf("ll.IndexOf(7) should return -1, not %d", pl)
	}
}

func TestLinkedListLength(t *testing.T) {
	ll := New[int]()
	var len uint

	for i := uint(0); i < 6; {
		len = ll.Size()
		if len != i {
			t.Fatalf("ll.Size() should return %d, not %d", i, len)
		}

		ll.AddLast(int(i))
		i++
	}

	for i := uint(6); i > 0; {
		len = ll.Size()
		if len != i {
			t.Fatalf("ll.Size() should return %d, not %d", i, len)
		}

		ll.Remove(0)
		i--
	}
}
