package data

type Graph[T comparable] struct {
	nodes     []T
	adjList   [][]*T
	positions map[T]int
}

func NewGraph[T comparable]() Graph[T] {
	return Graph[T]{
		make([]T, 0),
		make([][]*T, 0),
		make(map[T]int),
	}
}

func (g *Graph[T]) AddNode(t T) {
	if _, ok := g.positions[t]; !ok {
		g.positions[t] = len(g.nodes)
		g.nodes = append(g.nodes, t)
		g.adjList = append(g.adjList, make([]*T, 0))
	}
}

func (g *Graph[T]) AddEdge(t1, t2 T) {
	if _, ok := g.positions[t1]; !ok {
		g.AddNode(t1)
	}
	if _, ok := g.positions[t2]; !ok {
		g.AddNode(t2)
	}
	if !contains(g.adjList[g.positions[t1]], &t2) {
		g.adjList[g.positions[t1]] = append(g.adjList[g.positions[t1]], &t2)
		g.adjList[g.positions[t2]] = append(g.adjList[g.positions[t2]], &t1)

	}
}

func (g *Graph[T]) Nodes() []T {
	return g.nodes
}

func (g *Graph[T]) Edges(t1 T) []*T {
	return g.adjList[g.positions[t1]]
}

func GraphFromEdges[T comparable](edges []Edge[T]) Graph[T] {
	g := NewGraph[T]()
	for _, e := range edges {
		g.AddEdge(e.f, e.t)
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
