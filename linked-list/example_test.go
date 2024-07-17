package linkedlist

import "fmt"

func Example() {
	ll := New[int]()
	ll.Add(0, 12)
	ll.AddFirst(11)
	ll.AddFirst(10)
	ll.AddLast(15)
	ll.Add(3, 13)
	ll.Add(5, 16)

	for n := ll.Head(); n != nil; n = n.Next() {
		fmt.Println(n.Val)
	}

	//Output:
	//10
	//11
	//12
	//13
	//15
	//16
}
