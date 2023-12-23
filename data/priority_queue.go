package data

type PriorityQueue[T comparable] struct {
	heap Heap[T]
}

func (q *PriorityQueue[T]) ExtractMin() {
	q.heap.Pop(0)
}

func (q *PriorityQueue[T]) ChangeKey(t T, priority int) {
	q.heap.ChangeKey(t, priority)
}
