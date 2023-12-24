package data

import "fmt"

type BTree[T any] struct {
	v    T
	l, r *BTree[T]
}

func NewBTreeNode[T any](v T, l, r *BTree[T]) BTree[T] {
	return BTree[T]{v, l, r}
}

func (bt *BTree[T]) Value() T {
	return bt.v
}

func (bt *BTree[T]) String() string {
	if bt == nil {
		return "nil"
	} else {
		return fmt.Sprintf(
			"(v: %v, l: %s, r: %s)",
			bt.v,
			bt.l.String(),
			bt.r.String(),
		)
	}
}
