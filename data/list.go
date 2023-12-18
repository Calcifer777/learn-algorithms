package data

type List[T any] struct {
	v    T
	next *List[T]
}
