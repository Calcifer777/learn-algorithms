package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromEdges(t *testing.T) {
	g := GraphFromEdges(
		[]Edge[int]{
			NewEdge(1, 2, 1),
			NewEdge(1, 3, 1),
			NewEdge(2, 4, 1),
			NewEdge(2, 5, 1),
		},
		false,
	)
	assert.ElementsMatch(t, g.Nodes(), []int{1, 2, 3, 4, 5})
	assert.Len(t, g.Edges(1), 2)
	assert.Len(t, g.Edges(2), 3)
	assert.Len(t, g.Edges(5), 1)
}
