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
			for _, linked := range g.Edges(n) {
				if !visited[*linked] {
					nextLayer = append(nextLayer, *linked)
					visited[*linked] = true
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
		for _, linked := range g.Edges(*next) {
			if !visited[*linked] {
				visited[*linked] = true
				queue.Push(*linked)
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
			for _, linked := range toVisit {
				stack.Push(*linked)
				if *linked == t {
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
