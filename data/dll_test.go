package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	dll := NewDLL[int]()
	assert.Empty(t, dll.first)
	assert.Empty(t, dll.last)
}

func TestPushOne(t *testing.T) {
	dll := NewDLL[int]()
	dll.Push(3, true)
	assert.Equal(t, 3, dll.first.v)
	assert.Equal(t, 3, dll.last.v)
}

func TestPushMany(t *testing.T) {
	dll := NewDLL[int]()
	dll.Push(3, true)
	dll.Push(2, true)
	dll.Push(1, true)
	assert.Equal(t, 1, dll.first.v)
	assert.Equal(t, 3, dll.last.v)
}

func TestPopEmpty(t *testing.T) {
	dll := NewDLL[int]()
	ok := dll.Pop(true)
	assert.Equal(t, false, ok)
}

func TestPopOne(t *testing.T) {
	dll := NewDLL[int]()
	dll.Push(3, true)
	dll.Push(2, true)
	dll.Push(1, true)
	dll.Pop(true)
	assert.Equal(t, 2, dll.first.v)
	assert.Equal(t, 3, dll.last.v)
}

func TestPopSingle(t *testing.T) {
	dll := NewDLL[int]()
	dll.Push(1, true)
	dll.Pop(true)
	assert.Empty(t, dll.first)
	assert.Empty(t, dll.last)
}

func TestPopMany(t *testing.T) {
	dll := NewDLL[int]()
	dll.Push(1, true)
	dll.Push(2, true)
	dll.Pop(true)
	dll.Pop(true)
	assert.Empty(t, dll.first)
	assert.Empty(t, dll.last)
}
