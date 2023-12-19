package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromEdges(t *testing.T) {
	g := GraphFromEdges(
		[]Edge[int]{
			NewEdge(1, 2),
			NewEdge(1, 3),
			NewEdge(2, 4),
			NewEdge(2, 5),
		},
	)
	assert.ElementsMatch(t, g.Nodes(), []int{1, 2, 3, 4, 5})
	assert.Len(t, g.Edges(1), 2)
	assert.Len(t, g.Edges(2), 3)
	assert.Len(t, g.Edges(5), 1)
}
