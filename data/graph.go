package data

type Graph[T comparable] struct {
	nodes     []T
	adjList   [][]Edge[T]
	Positions map[T]int
	directed  bool
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
	if _, ok := g.Positions[t]; !ok {
		g.Positions[t] = len(g.nodes)
		g.nodes = append(g.nodes, t)
		g.adjList = append(g.adjList, make([]Edge[T], 0))
	}
}

func (g *Graph[T]) AddEdge(e Edge[T]) {
	if _, ok := g.Positions[e.f]; !ok {
		g.AddNode(e.f)
	}
	if _, ok := g.Positions[e.t]; !ok {
		g.AddNode(e.t)
	}
	if !contains(g.adjList[g.Positions[e.f]], e) {
		g.adjList[g.Positions[e.f]] = append(g.adjList[g.Positions[e.f]], e)
	}
	if !g.directed {
		edgeRev := NewEdge(e.f, e.t, e.d)
		g.adjList[g.Positions[e.t]] = append(g.adjList[g.Positions[e.t]], edgeRev)
	}
}

func (g *Graph[T]) Nodes() []T {
	return g.nodes
}

func (g *Graph[T]) Edges(t1 T) []Edge[T] {
	return g.adjList[g.Positions[t1]]
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
