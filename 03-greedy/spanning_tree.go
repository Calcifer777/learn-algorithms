package greedy

import (
	"slices"

	"github.com/calcifer777/learn-algorithms/data"
)

type NodeEdge struct {
	n    int
	edge data.Edge[int]
}

func Prim[T comparable](g data.Graph[T]) data.Graph[T] {
	// not Kleinberg impl!
	nodes := g.Nodes()
	pq := data.NewPriorityQueue[NodeEdge](10)
	root := nodes[0]
	for _, e := range g.Edges(root) {
		pq.Push(NodeEdge{e.To(), e}, e.Dist())
	}
	exploredNodes := make(map[int]bool)
	treeEdges := make([]data.Edge[int], 0)
	for {
		nodeEdge, _, ok := pq.Pop()
		if !ok {
			break
		}
		_, explored := exploredNodes[nodeEdge.n]
		if !explored {
			exploredNodes[nodeEdge.n] = true
			treeEdges = append(treeEdges, nodeEdge.edge)
			for _, e := range g.Edges(nodeEdge.n) {
				pq.Push(NodeEdge{e.To(), e}, e.Dist())
			}
		}
	}
	//
	edges := make([]data.Edge[T], 0)
	for _, e := range treeEdges {
		edgeLabel := data.NewEdge(
			g.GetNodeLabel(e.From()),
			g.GetNodeLabel(e.To()),
			e.Dist(),
		)
		edges = append(edges, edgeLabel)
	}
	return data.GraphFromEdges(edges, g.IsDirected())

}

func Kruskal[T comparable](g data.Graph[T]) data.Graph[T] {
	edges := g.GetEdges()
	slices.SortFunc(edges, func(e1, e2 data.Edge[int]) int {
		return e1.Dist() - e2.Dist()
	})
	unionFind := data.NewUnionFind(g.Nodes())
	treeEdges := make([]data.Edge[T], 0)
	for _, e := range edges {
		componentFrom := unionFind.Find(e.From())
		componentTo := unionFind.Find(e.To())
		if componentFrom != componentTo {
			unionFind.Union(e.From(), e.To())
			treeEdges = append(
				treeEdges,
				data.NewEdge[T](
					g.GetNodeLabel(e.From()),
					g.GetNodeLabel(e.To()),
					e.Dist(),
				),
			)
		}
	}
	return data.GraphFromEdges(treeEdges, true)
}
