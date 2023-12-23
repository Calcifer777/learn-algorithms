package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSingle(t *testing.T) {
	h := NewHeap[string](10)
	h.Add("a", 3)
	assert.Equal(t, 3, h.priorities[0])
}

func TestAddHeapifyUp(t *testing.T) {
	h := NewHeap[string](10)
	h.Add("e", 7)
	h.Add("a", 3)
	h.Add("b", 2)
	h.Add("c", 5)
	h.Add("d", 1)
	fmt.Sprintln(h.String())
	assert.Equal(t, 1, h.priorities[0])
	assert.Equal(t, false, h.IsCorrupt())
}

func TestPop(t *testing.T) {
	h := NewHeap[string](10)
	h.Add("a", 1)
	h.Add("b", 2)
	h.Add("c", 5)
	h.Pop(2)
	assert.Equal(t, 2, h.size)
	assert.Equal(t, h.IsCorrupt(), false)
}

func TestPopHeapifyDown(t *testing.T) {
	h := NewHeap[string](15)
	h.Add("a", 4)
	h.Add("b", 7)
	h.Add("c", 7)
	h.Add("d", 10)
	h.Add("e", 16)
	h.Add("f", 7)
	h.Add("g", 11)
	h.Add("j", 15)
	h.Add("k", 17)
	h.Add("l", 20)
	h.Add("m", 17)
	h.Add("n", 15)
	h.Add("o", 8)
	h.Add("p", 16)
	h.Add("q", 21)
	fmt.Println(h.String())
	assert.Equal(t, h.IsCorrupt(), false)
	h.Pop(2)
	fmt.Println(h.String())
}

func TestPopHeapifyDownLast(t *testing.T) {
	h := NewHeap[string](15)
	h.Add("a", 4)
	h.Add("b", 7)
	h.Add("c", 7)
	h.Add("d", 10)
	h.Add("e", 16)
	h.Add("f", 7)
	h.Add("g", 11)
	fmt.Println(h.String())
	assert.Equal(t, h.IsCorrupt(), false)
	h.Pop(14)
	fmt.Println(h.String())
}

func TestChangeKey(t *testing.T) {
	h := NewHeap[string](10)
	h.Add("e", 7)
	h.Add("a", 3)
	h.Add("b", 2)
	h.Add("c", 5)
	h.Add("d", 1)
	fmt.Println(h.String())
	h.ChangeKey("d", 8)
	fmt.Println(h.String())
	assert.Equal(t, "b", h.positions[0])
	assert.Equal(t, false, h.IsCorrupt())
}
