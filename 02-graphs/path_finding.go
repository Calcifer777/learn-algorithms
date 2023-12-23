package graphs

import (
	"slices"

	"github.com/calcifer777/learn-algorithms/data"
)

func Bfs[T comparable](g data.Graph[T], f T, t T) int {
	visited := make(map[T]bool)
	visited[f] = true
	bfsLayers := make([][]T, 0)
	bfsLayers = append(bfsLayers, []T{f})
	layerIdx := 0
	for len(bfsLayers[layerIdx]) > 0 {
		nextLayer := make([]T, 0)
		for _, n := range bfsLayers[layerIdx] {
			for _, edge := range g.Edges(n) {
				if !visited[edge.To()] {
					nextLayer = append(nextLayer, edge.To())
					visited[edge.To()] = true
				}
			}
		}
		bfsLayers = append(bfsLayers, nextLayer)
		layerIdx += 1
		if visited[t] {
			return layerIdx
		}
	}
	return -1
}

func BfsQueue[T comparable](g data.Graph[T], f T, t T) int {
	visited := make(map[T]bool)
	visited[f] = true
	queue := data.NewQueue[T]()
	queue.Push(f)
	steps := 0
	for {
		next, ok := queue.Pop()
		if !ok {
			return -1
		}
		for _, edge := range g.Edges(*next) {
			if !visited[edge.To()] {
				visited[edge.To()] = true
				queue.Push(edge.To())
			}
		}
		steps += 1
		if visited[t] {
			return steps
		}
	}
}

func DfsStack[T comparable](g data.Graph[T], f T, t T) int {
	explored := make(map[T]bool)
	stack := data.NewStack[T]()
	stack.Push(f)
	steps := 0
	found := false
	for {
		next, ok := stack.Pop()
		if !ok {
			return -1
		}
		if !explored[*next] {
			explored[*next] = true
			toVisit := g.Edges(*next)
			slices.Reverse(toVisit)
			for _, edge := range toVisit {
				stack.Push(edge.To())
				if edge.To() == t {
					found = true
				}
			}
		}
		if found {
			return steps
		}
		steps += 1
	}
}
