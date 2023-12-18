package data

import (
	"fmt"
	"log/slog"
	"strings"
)

type Heap[T any] struct {
	arr  []T
	size int
	less func(t1, t2 T) bool
}

func NewHeap[T any](size int, less func(t1, t2 T) bool) Heap[T] {
	return Heap[T]{make([]T, size), 0, less}
}

func (h *Heap[T]) Add(t T) {
	h.arr[h.size] = t
	h.size += 1
	h.HeapifyUp(h.size - 1)
}

func (h *Heap[T]) HeapifyUp(i int) {
	var parentIdx int
	if i == 0 {
		return
	}
	if i%2 == 0 {
		parentIdx = i/2 - 1
	} else {
		parentIdx = i / 2
	}
	if h.less(h.arr[i], h.arr[parentIdx]) {
		h.arr[parentIdx], h.arr[i] = h.arr[i], h.arr[parentIdx]
		h.HeapifyUp(parentIdx)
	} else {
		slog.Info("HeapifyUp: Nothing to do")
	}
}

func (h *Heap[T]) Pop(i int) {
	h.size -= 1
	h.arr[i] = h.arr[h.size]
	h.HeapifyUp(i)
	h.HeapifyDown(i)
}

func (h *Heap[T]) HeapifyDown(i int) {
	var childIdx int
	// check if we removed a leaf
	if (i+1)*2 >= h.size {
		return
	}
	// pick child with lower value
	if h.less(h.arr[i*2+1], h.arr[i*2+2]) {
		childIdx = i*2 + 1
	} else {
		childIdx = i*2 + 2
	}
	slog.Info("HeapifyDown",
		slog.Int("i", i),
		slog.Any("v", h.arr[i]),
		slog.Any("c-l", h.arr[i*2]),
		slog.Any("c-r", h.arr[i*2+1]),
		slog.Any("cIdx", childIdx),
	)
	if h.less(h.arr[childIdx], h.arr[i]) {
		h.arr[i], h.arr[childIdx] = h.arr[childIdx], h.arr[i]
		h.HeapifyDown(childIdx)
	}
}

func (h *Heap[T]) IsCorrupt() bool {
	for i := 0; i < h.size; i++ {
		if i > i*2 || i > i*2+1 {
			return true
		}
	}
	return false
}

func (h *Heap[T]) String() string {
	lines := make([][]string, 0)
	i := 0
	// Create lines
	for i < h.size {
		line := make([]string, 0)
		for j := i; j <= i*2 && j < h.size; j++ {
			line = append(line, fmt.Sprintf("%4v", h.arr[j]))
		}
		lines = append(lines, line)
		i = i<<1 + 1
	}
	// Create line strings
	out := ""
	for i := 0; i < len(lines); i++ {
		sep := strings.Repeat(" ", (len(lines)-i)*4)
		prefix := strings.Repeat(" ", (len(lines)-1-i)*4*2)
		out += prefix + strings.Join(lines[i], sep) + "\n\n"
	}
	return out
}
