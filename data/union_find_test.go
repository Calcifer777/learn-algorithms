package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFindUnion(t *testing.T) {
	nodes := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	uf := NewUnionFind(nodes)
	uf.Union(0, 9)
	assert.Equal(t, 0, uf.Find(9))
}

func TestUnionFindSingle(t *testing.T) {
	nodes := []int{0, 1, 2, 3, 4, 5}
	uf := NewUnionFind(nodes)
	uf.Union(0, 5)
	uf.Union(1, 4)
	uf.Union(2, 3)
	uf.Union(1, 5)
	uf.Union(2, 4)
	for _, n := range nodes {
		assert.Equal(t, 0, uf.Find(n))
	}
}
