package data

type Graph[T comparable] struct {
	nodes    []T
	adjList  [][]Edge[T]
	nodeToId map[T]int
	directed bool
}

func NewGraph[T comparable](directed bool) Graph[T] {
	return Graph[T]{
		make([]T, 0),
		make([][]Edge[T], 0),
		make(map[T]int),
		directed,
	}
}

func (g *Graph[T]) AddNode(t T) {
	if _, ok := g.nodeToId[t]; !ok {
		g.nodeToId[t] = len(g.nodes)
		g.nodes = append(g.nodes, t)
		g.adjList = append(g.adjList, make([]Edge[T], 0))
	}
}

func (g *Graph[T]) AddEdge(e Edge[T]) {
	if _, ok := g.nodeToId[e.f]; !ok {
		g.AddNode(e.f)
	}
	if _, ok := g.nodeToId[e.t]; !ok {
		g.AddNode(e.t)
	}
	if !contains(g.adjList[g.nodeToId[e.f]], e) {
		g.adjList[g.nodeToId[e.f]] = append(g.adjList[g.nodeToId[e.f]], e)
	}
	if !g.directed {
		edgeRev := NewEdge(e.t, e.f, e.d)
		g.adjList[g.nodeToId[e.t]] = append(g.adjList[g.nodeToId[e.t]], edgeRev)
	}
}

func (g *Graph[T]) Nodes() []T {
	return g.nodes
}

func (g *Graph[T]) GetEdges() []Edge[T] {
	allEdges := make([]Edge[T], 0)
	for _, edges := range g.adjList {
		allEdges = append(allEdges, edges...)
	}
	return allEdges
}

func (g *Graph[T]) Edges(t1 T) []Edge[T] {
	return g.adjList[g.nodeToId[t1]]
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

func contains[T comparable](arr []T, needle T) bool {
	for _, t := range arr {
		if t == needle {
			return true
		}
	}
	return false
}
