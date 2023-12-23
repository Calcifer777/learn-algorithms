package data

type PriorityQueue[T comparable] struct {
	heap Heap[T]
}

func NewPriorityQueue[T comparable](capacity int) PriorityQueue[T] {
	return PriorityQueue[T]{NewHeap[T](capacity)}
}

func (q *PriorityQueue[T]) Push(t T, priority int) bool {
	return q.heap.Add(t, priority)
}

func (q *PriorityQueue[T]) Pop() (*T, int, bool) {
	return q.heap.Pop(0)
}

func (q *PriorityQueue[T]) Change(t T, priority int) {
	q.heap.ChangeKey(t, priority)
}

func (q *PriorityQueue[T]) Value(t T) (int, bool) {
	return q.heap.Value(t)
}
