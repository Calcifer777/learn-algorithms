package greedy

import (
	"math"

	"github.com/calcifer777/learn-algorithms/data"
)

func Prim[T comparable](g data.Graph[T]) data.Graph[T] {
	nodes := g.Nodes()
	root := nodes[0]
	pq := data.NewPriorityQueue[int](20)
	for _, n := range g.Nodes() {
		if n == root {
			pq.Push(n, 0)
		} else {
			pq.Push(n, math.MaxInt16)
		}
	}
	treeEdges := make(map[int]data.Edge[int])
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
		edgeLabel := data.NewEdge(
			g.GetNodeLabel(e.From()),
			g.GetNodeLabel(e.To()),
			e.Dist(),
		)
		edges = append(edges, edgeLabel)
	}
	return data.GraphFromEdges(edges, true)
}

// func Kruskal[T comparable](g data.Graph[T]) data.Graph[T] {
// 	edges := g.GetEdges()
// 	slices.SortFunc(edges, func(e1, e2 data.Edge[T]) int {
// 		return e1.Dist() - e2.Dist()
// 	})
// 	// probably can be skipped if graph holds node ids (int)
// 	// instead of values (T)
// 	edgeIds := make([]int, len(edges))
// 	for i, _ := range edges {
// 		edgeIds[i] = i
// 	}
// 	unionFind := data.NewUnionFind(edgeIds)
// 	for i := range edges {
// 		if unionFind.Find()
// 	}
//
// 	treeEdges := make([]data.Edge[T], 0)
// 	// return data.GraphFromEdges(edges, true)
// }
