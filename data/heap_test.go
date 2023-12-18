package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSingle(t *testing.T) {
	less := func(i, j int) bool { return i < j }
	h := NewHeap[int](10, less)
	h.Add(3)
	assert.Equal(t, 3, h.arr[0])
}

func TestAddHeapifyUp(t *testing.T) {
	less := func(i, j int) bool { return i < j }
	h := NewHeap[int](10, less)
	h.Add(3)
	h.Add(2)
	h.Add(5)
	h.Add(1)
	h.Add(7)
	assert.Equal(t, 1, h.arr[0])
	assert.Equal(t, h.IsCorrupt(), false)
}

func TestPop(t *testing.T) {
	less := func(i, j int) bool { return i < j }
	h := NewHeap[int](10, less)
	h.Add(1)
	h.Add(2)
	h.Add(5)
	h.Pop(2)
	assert.Equal(t, 2, h.size)
	assert.Equal(t, h.IsCorrupt(), false)
}

func TestPopHeapifyDown(t *testing.T) {
	less := func(i, j int) bool { return i < j }
	h := NewHeap[int](15, less)
	h.Add(4)
	h.Add(7)
	h.Add(7)
	h.Add(10)
	h.Add(16)
	h.Add(7)
	h.Add(11)
	h.Add(15)
	h.Add(17)
	h.Add(20)
	h.Add(17)
	h.Add(15)
	h.Add(8)
	h.Add(16)
	h.Add(21)
	fmt.Println(h.String())
	assert.Equal(t, h.IsCorrupt(), false)
	h.Pop(2)
	fmt.Println(h.String())
}
