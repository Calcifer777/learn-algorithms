package graphs

import (
	"testing"

	"github.com/calcifer777/learn-algorithms/data"
	"github.com/stretchr/testify/assert"
)

func TestBFSFound(t *testing.T) {
	g := data.GraphFromEdges(
		[]data.Edge[int]{
			data.NewEdge(1, 2, 1),
			data.NewEdge(1, 3, 1),
			data.NewEdge(2, 4, 1),
			data.NewEdge(2, 5, 1),
			data.NewEdge(3, 6, 1),
			data.NewEdge(3, 7, 1),
			data.NewEdge(7, 8, 1),
		},
	)
	dist := Bfs(g, 1, 8)
	assert.Equal(t, 3, dist)
}

func TestBFSQueueFound(t *testing.T) {
	g := data.GraphFromEdges(
		[]data.Edge[int]{
			data.NewEdge(1, 2, 1),
			data.NewEdge(1, 3, 1),
			data.NewEdge(2, 4, 1),
			data.NewEdge(2, 5, 1),
			data.NewEdge(3, 6, 1),
			data.NewEdge(3, 7, 1),
			data.NewEdge(7, 8, 1),
		},
	)
	steps := BfsQueue(g, 1, 8)
	assert.Equal(t, 7, steps)
}

func TestDFSQueueFound(t *testing.T) {
	g := data.GraphFromEdges(
		[]data.Edge[int]{
			data.NewEdge(1, 2, 1),
			data.NewEdge(1, 3, 1),
			data.NewEdge(2, 4, 1),
			data.NewEdge(2, 5, 1),
			data.NewEdge(3, 6, 1),
			data.NewEdge(3, 7, 1),
			data.NewEdge(4, 8, 1),
		},
	)
	steps := DfsStack(g, 1, 8)
	assert.Equal(t, 3, steps)
}
