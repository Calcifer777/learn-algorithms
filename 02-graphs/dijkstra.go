package graphs

import (
	"math"

	"github.com/calcifer777/learn-algorithms/data"
)

func Dijkstra[T comparable](g data.Graph[T], start T, dest T) int {
	nodes := g.Nodes()
	pq := data.NewPriorityQueue[T](20)
	for _, n := range nodes {
		if n == start {
			pq.Push(start, 0)
		} else {
			pq.Push(n, math.MaxInt16)
		}
	}
	for {
		n, d, ok := pq.Pop()
		if !ok {
			break
		}
		if *n == dest {
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
