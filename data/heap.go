package data

import (
	"fmt"
	"log/slog"
	"math"
	"strings"
)

type Heap[T comparable] struct {
	priorities []int
	size       int
	positions  []T
	capacity   int
}

func NewHeap[T comparable](capacity int) Heap[T] {
	return Heap[T]{
		make([]int, capacity),
		0,
		make([]T, capacity),
		capacity,
	}
}

func (h *Heap[T]) Add(t T, priority int) bool {
	if h.size == h.capacity {
		return false
	} else {
		h.priorities[h.size] = priority
		h.positions[h.size] = t
		h.size += 1
		h.HeapifyUp(h.size - 1)
		return true
	}
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
	if h.priorities[i] < h.priorities[parentIdx] {
		h.priorities[parentIdx], h.priorities[i] = h.priorities[i], h.priorities[parentIdx]
		h.positions[i], h.positions[parentIdx] = h.positions[parentIdx], h.positions[i]
		h.HeapifyUp(parentIdx)
	} else {
		slog.Info("HeapifyUp: Nothing to do")
	}
}

func (h *Heap[T]) Pop(i int) (*T, bool) {
	if h.size == 0 {
		return nil, false
	} else {
		h.size -= 1
		out := h.positions[i]
		h.priorities[i] = h.priorities[h.size]
		h.HeapifyUp(i)
		h.HeapifyDown(i)
		return &out, true
	}
}

func (h *Heap[T]) HeapifyDown(i int) {
	var childIdx int
	// check if we removed a leaf
	if (i+1)*2 >= h.size {
		return
	}
	// pick child with lower value
	if h.priorities[i*2+1] < h.priorities[i*2+2] {
		childIdx = i*2 + 1
	} else {
		childIdx = i*2 + 2
	}
	if h.priorities[childIdx] < h.priorities[i] {
		h.priorities[i], h.priorities[childIdx] = h.priorities[childIdx], h.priorities[i]
		h.positions[i], h.positions[childIdx] = h.positions[childIdx], h.positions[i]
		h.HeapifyDown(childIdx)
	}
}

func (h *Heap[T]) IsCorrupt() bool {
	for i := 0; i < h.size; i++ {
		parentP := h.priorities[i]
		if (i*2+1 < h.size && parentP > h.priorities[i*2+1]) ||
			(i*2+2 < h.size && parentP > h.priorities[i*2+2]) {
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
		chunks := make([]string, 0)
		for j := i; j <= i*2 && j < h.size; j++ {
			chunks = append(
				chunks,
				fmt.Sprintf("%v: %d", h.positions[j], h.priorities[j]),
			)
		}
		lines = append(lines, chunks)
		i = i<<1 + 1
	}
	// Create line strings
	spacesPrefix := int(math.Ceil(math.Log(float64(h.size)))) * 4
	spacesSep := int(math.Floor(math.Log(float64(h.size)))) * 3
	out := ""
	for i := 0; i < len(lines); i++ {
		prefix := strings.Repeat(" ", (len(lines)-1-i)*spacesPrefix)
		sep := strings.Repeat(" ", (len(lines)-i)*spacesSep)
		out += prefix + strings.Join(lines[i], sep) + "\n\n"
	}
	return out
}

func (h *Heap[T]) ChangeKey(t T, priority int) bool {
	pos, ok := IndexOf(h.positions, t)
	if !ok {
		return false
	} else {
		h.priorities[pos] = priority
		h.HeapifyUp(pos)
		h.HeapifyDown(pos)
		return true
	}
}

func IndexOf[T comparable](arr []T, needle T) (int, bool) {
	for i, t := range arr {
		if t == needle {
			return i, true
		}
	}
	return -1, false
}
