package greedy

import (
	"testing"

	"github.com/calcifer777/learn-algorithms/data"
	"github.com/stretchr/testify/assert"
)

func TestPrim(t *testing.T) {
	g := data.GraphFromEdges([]data.Edge[string]{
		data.NewEdge("a", "b", 1),
		data.NewEdge("a", "c", 3),
		data.NewEdge("a", "d", 4),
		data.NewEdge("b", "d", 2),
		data.NewEdge("b", "e", 3),
		data.NewEdge("b", "f", 10),
		data.NewEdge("c", "d", 5),
		data.NewEdge("c", "e", 2),
		data.NewEdge("c", "f", 6),
		data.NewEdge("e", "f", 1),
	}, true)
	tree := Prim(g)
	expected := []data.Edge[string]{
		data.NewEdge("a", "b", 1),
		data.NewEdge("a", "c", 3),
		data.NewEdge("b", "d", 2),
		data.NewEdge("b", "e", 3),
		data.NewEdge("e", "f", 1),
	}
	assert.ElementsMatch(t, expected, tree.GetEdgesLabels())
}
