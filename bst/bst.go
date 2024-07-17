// This package implements binary search tree
package bst

import (
	"cmp"
	"slices"
)

type Node[T cmp.Ordered] struct {
	Val    T
	Occurs uint
	left   *Node[T]
	right  *Node[T]
}

// Binary search tree data structure, useful for searching data
type BST[T cmp.Ordered] struct {
	root *Node[T]
	size uint
}

// Return the left child of this node
func (n *Node[T]) Left() *Node[T] {
	return n.left
}

// Return the right child of this node
func (n *Node[T]) Right() *Node[T] {
	return n.right
}

// Initialize or clear binary search tree
func (bst *BST[T]) Init() *BST[T] {
	bst.root = nil
	bst.size = 0
	return bst
}

// Get root of the binary search tree
func (bst *BST[T]) Root() *Node[T] {
	return bst.root
}

// Create a new binary search tree
func New[T cmp.Ordered]() *BST[T] {
	return new(BST[T]).Init()
}

// Return the number of nodes in the binary search tree
func (bst *BST[T]) Size() uint {
	return bst.size
}

// Add a new node with `data` to the binary search tree IF the `data` does not
// exist in the tree, else it will only increase the occurrences of the `data`
func (bst *BST[T]) Add(data T) {
	if bst.root == nil { // The tree is empty
		bst.root = createNode(data)
		bst.size++
		return
	}

	var parent *Node[T]
	var left bool
	curr := bst.root
	for curr != nil {
		parent = curr
		if curr.Val == data { // `data` is existed, increase occurences
			curr.Occurs++
			return
		}

		if data < curr.Val {
			curr = curr.left
			left = true
		} else {
			curr = curr.right
			left = false
		}
	}

	curr = createNode(data)
	if left {
		parent.left = curr
	} else {
		parent.right = curr
	}
	bst.size++
}

// Remove `data` in binary search tree. If `ignoreOccurs` is set to true, it will
// remove the node (if found) in the tree. Else, it only decrease the occurences of
// the `data` by 1, remove the node if `Occurs` = 0
func (bst *BST[T]) Remove(data T, ignoreOccurs bool) {
	var parent *Node[T]
	curr := bst.root

	// find the node
	for curr != nil {
		parent = curr
		if curr.Val == data {
			break
		}

		if data < curr.Val {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	// not found
	if curr == nil {
		return
	}

	// only need to decrease the occurences of the `data`
	if !ignoreOccurs && curr.Occurs > 1 {
		curr.Occurs--
		return
	}

	// if current node has both branches, then find the left most node of the next
	// right node and swap the value
	if curr.left != nil && curr.right != nil {
		deletedNode := curr // save the node that about to be deleted (change the value)
		parent = curr
		curr = curr.right

		for curr.left != nil {
			parent = curr
			curr = curr.left
		}

		deletedNode.Val = curr.Val
	}

	// now curr node is the node about to be deleted (and it cannot have 2 branches).
	// If it still have a branch then the branch will be connected with parent node,
	// else it will be a leaf, we just need to remove it
	var child *Node[T] = nil
	if curr.left != nil {
		child = curr.left
	} else {
		child = curr.right
	}

	if curr == bst.root {
		bst.root = child
	} else {
		if parent.left == curr {
			parent.left = child
		} else {
			parent.right = child
		}
	}

	bst.size--
}

// Check if the binary search tree contains this `data` or not
func (bst *BST[T]) Contains(data T) bool {
	curr := bst.root

	for curr != nil {
		if curr.Val == data {
			return true
		}

		if data < curr.Val {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return false
}

func (bst *BST[T]) InOrder() []T {
	if bst.size == 0 {
		return nil
	}

	stack := make([]*Node[T], bst.size)
	top := -1
	result := make([]T, bst.size)
	i := -1

	curr := bst.root
	for curr != nil || top > -1 {
		for curr != nil {
			push(stack, &top, curr)
			curr = curr.left
		}

		curr = pop(stack, &top)
		push(result, &i, curr.Val)
		curr = curr.right
	}

	return result
}

func (bst *BST[T]) PreOrder() []T {
	if bst.size == 0 {
		return nil
	}

	stack := make([]*Node[T], bst.size)
	top := -1
	result := make([]T, bst.size)
	i := -1

	curr := bst.root
	for curr != nil || top > -1 {
		for curr != nil {
			push(result, &i, curr.Val)
			push(stack, &top, curr)
			curr = curr.left
		}

		curr = pop(stack, &top)
		curr = curr.right
	}

	return result
}

func (bst *BST[T]) PostOrder() []T {
	if bst.size == 0 {
		return nil
	}

	stack := make([]*Node[T], bst.size)
	top := -1
	result := make([]T, bst.size)
	i := -1

	curr := bst.root
	for curr != nil || top > -1 {
		for curr != nil {
			push(result, &i, curr.Val)
			push(stack, &top, curr)
			curr = curr.right
		}

		curr = pop(stack, &top)
		curr = curr.left
	}

	slices.Reverse(result)
	return result
}

func pop[T any](s []T, top *int) T {
	data := s[*top]
	*top--
	return data
}

func push[T any](s []T, top *int, data T) {
	*top++
	s[*top] = data
}

func createNode[T cmp.Ordered](data T) *Node[T] {
	return &Node[T]{Val: data, Occurs: 1, left: nil, right: nil}
}
