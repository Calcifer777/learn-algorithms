package data

import (
	"log/slog"
	"slices"
)

type Graph[T comparable] struct {
	nodes    []int
	labels   []T
	adjList  [][]Edge[int]
	nodeToId map[T]int
	directed bool
}

func NewGraph[T comparable](directed bool) Graph[T] {
	return Graph[T]{
		make([]int, 0),
		make([]T, 0),
		make([][]Edge[int], 0),
		make(map[T]int),
		directed,
	}
}

func (g *Graph[T]) AddNode(t T) int {
	nodeId, ok := g.nodeToId[t]
	if !ok {
		g.nodeToId[t] = len(g.nodes)
		g.nodes = append(g.nodes, len(g.nodes))
		g.labels = append(g.labels, t)
		g.adjList = append(g.adjList, make([]Edge[int], 0))
		nodeId = g.nodeToId[t]
	} else {
		slog.Info("Can't add node, `%v` already in graph!", t)
	}
	return nodeId
}

func (g *Graph[T]) AddEdge(e Edge[T]) {
	nodeFrom, ok := g.nodeToId[e.f]
	if !ok {
		nodeFrom = g.AddNode(e.f)
	}
	nodeTo, ok := g.nodeToId[e.t]
	if !ok {
		nodeTo = g.AddNode(e.t)
	}
	edge := NewEdge(nodeFrom, nodeTo, e.d)
	if slices.Index(g.adjList[nodeFrom], edge) < 0 {
		g.adjList[nodeFrom] = append(g.adjList[nodeFrom], edge)
	}
	if !g.directed {
		edgeRev := NewEdge(nodeTo, nodeFrom, e.d)
		if slices.Index(g.adjList[nodeTo], edgeRev) < 0 {
			g.adjList[nodeTo] = append(g.adjList[nodeTo], edgeRev)
		}
	}
}

func (g *Graph[T]) Nodes() []int {
	return g.nodes
}

func (g *Graph[T]) NodeLabels() []T {
	return g.labels
}

func (g *Graph[T]) GetEdgesLabels() []Edge[T] {
	allEdges := make([]Edge[T], 0)
	for _, edges := range g.adjList {
		for _, e := range edges {
			edgeLabel := NewEdge(
				g.GetNodeLabel(e.From()),
				g.GetNodeLabel(e.To()),
				e.Dist(),
			)
			allEdges = append(allEdges, edgeLabel)
		}
	}
	return allEdges
}

func (g *Graph[T]) GetNodeLabel(n int) T {
	return g.labels[n]
}

func (g *Graph[T]) GetNodeId(t T) int {
	return g.nodeToId[t]
}

func (g *Graph[T]) GetEdges() []Edge[int] {
	allEdges := make([]Edge[int], 0)
	for _, edges := range g.adjList {
		allEdges = append(allEdges, edges...)
	}
	return allEdges
}

func (g *Graph[T]) Edges(n int) []Edge[int] {
	return g.adjList[n]
}

func GraphFromEdges[T comparable](edges []Edge[T], directed bool) Graph[T] {
	g := NewGraph[T](directed)
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g

}

type Edge[T comparable] struct {
	f, t T
	d    int
}

func (e *Edge[T]) From() T { return e.f }

func (e *Edge[T]) To() T { return e.t }

func (e *Edge[T]) Dist() int { return e.d }

func NewEdge[T comparable](f, t T, d int) Edge[T] {
	return Edge[T]{f, t, d}
}
