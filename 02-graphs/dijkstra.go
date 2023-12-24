package graphs

import (
	"math"

	"github.com/calcifer777/learn-algorithms/data"
)

func Dijkstra[T comparable](g data.Graph[T], start T, dest T) int {
	nodes := g.Nodes()
	pq := data.NewPriorityQueue[int](20)
	startId := g.GetNodeId(start)
	destId := g.GetNodeId(dest)
	for _, n := range nodes {
		priority := math.MaxInt16
		if n == startId {
			priority = 0
		}
		pq.Push(n, priority)
	}
	for {
		n, d, ok := pq.Pop()
		if !ok {
			break
		}
		if *n == destId {
			return d
		}
		for _, edge := range g.Edges(*n) {
			linkedDist, _ := pq.Value(edge.To())
			if linkedDist > d+edge.Dist() {
				pq.Change(edge.To(), d+edge.Dist())
			}
		}
	}
	return -1
}
