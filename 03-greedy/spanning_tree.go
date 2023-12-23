package greedy

import (
	"math"

	"github.com/calcifer777/learn-algorithms/data"
)

func Prim[T comparable](g data.Graph[T]) data.Graph[T] {
	nodes := g.Nodes()
	root := nodes[0]
	pq := data.NewPriorityQueue[T](20)
	for _, n := range g.Nodes() {
		if n == root {
			pq.Push(n, 0)
		} else {
			pq.Push(n, math.MaxInt16)
		}
	}
	treeEdges := make(map[T]data.Edge[T])
	for {
		node, dist, ok := pq.Pop()
		if !ok {
			break
		}
		for _, edge := range g.Edges(*node) {
			linkedDist, _ := pq.Value(edge.To())
			if linkedDist > dist+edge.Dist() {
				pq.Change(edge.To(), dist+edge.Dist())
				treeEdges[edge.To()] = edge
			}
		}

	}
	edges := make([]data.Edge[T], 0)
	for _, e := range treeEdges {
		edges = append(edges, e)
	}
	return data.GraphFromEdges(edges, true)
}
