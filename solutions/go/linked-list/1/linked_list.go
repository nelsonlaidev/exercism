package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type List struct {
	first  *Node
	last   *Node
	length int
}

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

func NewList(elements ...any) *List {
	l := &List{}

	for _, el := range elements {
		l.Push(el)
	}

	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	newNode := Node{Value: v}

	if l.last == nil {
		l.first = &newNode
		l.last = &newNode
	} else {
		l.first.prev = &newNode
		newNode.next = l.first
		l.first = &newNode
	}

	l.length += 1
}

func (l *List) Push(v any) {
	newNode := Node{Value: v}

	if l.last == nil {
		l.first = &newNode
		l.last = &newNode
	} else {
		newNode.prev = l.last
		l.last.next = &newNode
		l.last = &newNode
	}

	l.length += 1
}

func (l *List) Shift() (any, error) {
	var removed any

	if l.last == nil {
		return nil, errors.New("list is empty")
	} else {
		removed = l.first.Value
		l.first = l.first.next
		if l.first == nil {
			l.first = nil
			l.last = nil
		} else {
			l.first.prev = nil
		}
	}

	l.length -= 1

	return removed, nil
}

func (l *List) Pop() (any, error) {
	var removed any

	if l.last == nil {
		return nil, errors.New("list is empty")
	} else {
		removed = l.last.Value
		l.last = l.last.prev
		if l.last == nil {
			l.first = nil
			l.last = nil
		} else {
			l.last.next = nil
		}
	}

	l.length -= 1

	return removed, nil
}

func (l *List) Reverse() {
	for current := l.first; current != nil; {
		next := current.Next()
		current.next, current.prev = current.prev, current.next
		current = next
	}
	l.first, l.last = l.last, l.first
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Count() int {
	return l.length
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	var current *Node

	for current = ll.first; current != nil; current = current.next {
		if current.Value == v {
			break
		} else if ll.length == 1 {
			return false
		}
	}

	if current == nil {
		return false
	}

	if ll.length == 1 {
		ll.first = nil
		ll.last = nil
		ll.length = 0

		return true
	}

	if ll.first == current {
		ll.first = ll.first.next
		ll.first.prev = nil
	} else if ll.last == current {
		ll.last = ll.last.prev
		ll.last.next = nil
	} else {
		current.prev.next = current.next
		current.next.prev = current.prev
	}

	ll.length -= 1

	return true
}
