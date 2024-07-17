package bst_test

import (
	"fmt"

	"github.com/dungtl2003/data-structure/bst"
)

func Example() {
	tree := bst.New[int]()
	tree.Add(40)
	tree.Add(30)
	tree.Add(25)
	tree.Add(35)
	tree.Add(50)
	tree.Add(45)
	tree.Add(60)
	tree.Add(40)

	fmt.Printf("%v\n", tree.InOrder())
	fmt.Printf("There are %d number %d\n", tree.Root().Occurs, tree.Root().Val)
	// Output:
	// [25 30 35 40 45 50 60]
	// There are 2 number 40
}
