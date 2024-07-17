// Package that implements linked list
// You can iterate the linked list using for loop.
// Example:
//
//	for node := list.Head(); node != nil; node.Next() {
//	     value = node.Val
//	}
package linkedlist

import (
	ex "ilikeblue/ds/exception"
)

type Node[T comparable] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}

// Linked list data structure
type LinkedList[T comparable] struct {
	size uint
	head *Node[T]
	tail *Node[T]
}

func createNode[T comparable](data T) *Node[T] {
	n := new(Node[T])
	n.Val = data
	n.next = nil
	n.prev = nil
	return n
}

// Get the next node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Get the previous node
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// Initialize or clear linked list
func (list *LinkedList[T]) Init() *LinkedList[T] {
	list.head = nil
	list.tail = nil
	list.size = 0
	return list
}

// Create new linked list
func New[T comparable]() *LinkedList[T] {
	list := new(LinkedList[T])
	return list.Init()
}

// Get the number of data in linked list
func (list *LinkedList[T]) Size() uint {
	return list.size
}

// Get the first node of the linked list
func (list *LinkedList[T]) Head() *Node[T] {
	return list.head
}

// Get the last node of the linked list
func (list *LinkedList[T]) Tail() *Node[T] {
	return list.tail
}

// Add data to the beginning of the linked list
func (list *LinkedList[T]) AddFirst(data T) {
	n := createNode(data)
	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		n.next = list.head
		list.head.prev = n
		list.head = n
	}

	list.size++
}

// Add data to the end of the linked list
func (list *LinkedList[T]) AddLast(data T) {
	n := createNode(data)
	if list.head == nil {
		list.head = n
		list.tail = n
	} else {
		list.tail.next = n
		n.prev = list.tail
		list.tail = n
	}

	list.size++
}

// Add data to `position` in the linked list (from 0), if the `position` is larger than
// list's length, this will be the same as `AddLast(data T)`
func (list *LinkedList[T]) Add(position uint, data T) {
	if position >= list.size {
		list.AddLast(data)
		return
	}

	if position == 0 {
		list.AddFirst(data)
		return
	}

	n := createNode(data)
	curr := list.head.next
	for i := uint(1); i != position; {
		curr = curr.next
		i++
	}

	n.next = curr
	n.prev = curr.prev
	curr.prev.next = n
	curr.prev = n
	list.size++
}

// Remove data from `position` of the linked list. Throws `IndexOutOfBound` if index
// is out of range [0, length - 1]
func (list *LinkedList[T]) Remove(position uint) (data T, err *ex.IndexOutOfBound) {
	if position < 0 || position >= list.size {
		err = &ex.IndexOutOfBound{Start: 0, End: list.size - 1}
		return
	}

	if list.size == 1 {
		data = list.head.Val
		list.head = nil
		list.tail = nil
	} else if position == 0 {
		data = list.head.Val
		list.head = list.head.next
		list.head.prev = nil
	} else if position == list.size-1 {
		data = list.tail.Val
		list.tail = list.tail.prev
		list.tail.next = nil
	} else {
		curr := list.head.next
		for i := uint(1); i != position; {
			curr = curr.next
			i++
		}

		data = curr.Val
		curr.prev.next = curr.next
		curr.next.prev = curr.prev
	}

	list.size--
	return
}

// Get the first data of the linked list
func (list *LinkedList[T]) GetFirst() (data T) {
	if list.size > 0 {
		return list.head.Val
	}

	return
}

// Get the last data of the linked list
func (list *LinkedList[T]) GetLast() (data T) {
	if list.size > 0 {
		return list.tail.Val
	}

	return
}

// Get the data of the linked list at `position`. Throws `IndexOutOfBound` if index
// is out of range [0, length - 1]
func (list *LinkedList[T]) Get(position uint) (data T, err *ex.IndexOutOfBound) {
	if position < 0 || position >= list.size {
		err = &ex.IndexOutOfBound{Start: 0, End: list.size - 1}
		return
	}

	curr := list.head
	for i := uint(0); i != position; {
		curr = curr.next
		i++
	}

	data = curr.Val
	return
}

// Return `true` if linked list contains data
func (list *LinkedList[T]) Contains(data T) bool {
	for curr := list.head; curr != nil; curr = curr.next {
		if curr.Val == data {
			return true
		}
	}

	return false
}

// Return the first data's position of the linked list. If not exists, it will
// return -1
func (list *LinkedList[T]) IndexOf(data T) int {
	position := 0
	for curr := list.head; curr != nil; {
		if curr.Val == data {
			return position
		}

		curr = curr.next
		position++
	}

	return -1
}

// Return the last data's position of the linked list. If not exists, it will
// return -1
func (list *LinkedList[T]) LastIndexOf(data T) int {
	position := int(list.size - 1)
	for curr := list.tail; curr != nil; {
		if curr.Val == data {
			return position
		}

		curr = curr.prev
		position--
	}

	return -1
}
