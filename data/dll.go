package data

type Node[T comparable] struct {
	v          T
	prev, next *Node[T]
}

func NewNode[T comparable](t T, prev, next *Node[T]) Node[T] {
	return Node[T]{t, prev, next}
}

type DoublyLinkedList[T comparable] struct {
	first, last *Node[T]
}

func NewDLL[T comparable]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{nil, nil}
}

func (l *DoublyLinkedList[T]) Push(t T, first bool) {
	if l.first == nil && l.last == nil {
		n := NewNode(t, nil, nil)
		l.first = &n
		l.last = &n
	} else if first {
		n := NewNode(t, l.last, l.first)
		l.first.prev = &n
		l.last.next = &n
		l.first = &n
	} else {
		n := NewNode(t, l.first, l.last)
		l.last.next = &n
		l.first.prev = &n
		l.last = &n
	}
}

func (l *DoublyLinkedList[T]) Pop(first bool) bool {
	if l.first == nil {
		return false
	} else if l.first == l.last {
		l.first = nil
		l.last = nil
		return true
	} else if first {
		l.first = l.first.next
		l.first.prev = l.last
		l.last.next = l.first
		return true
	} else {
		l.last = l.last.prev
		l.last.next = l.first
		l.first.prev = l.last
		return true
	}
}
