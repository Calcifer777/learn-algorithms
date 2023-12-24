package graphs

import (
	"slices"

	"github.com/calcifer777/learn-algorithms/data"
)

func Bfs[T comparable](g data.Graph[T], f T, t T) int {
	visited := make(map[int]bool)
	fromId := g.GetNodeId(f)
	toId := g.GetNodeId(t)
	visited[fromId] = true
	bfsLayers := make([][]int, 0)
	bfsLayers = append(bfsLayers, []int{fromId})
	layerIdx := 0
	for len(bfsLayers[layerIdx]) > 0 {
		nextLayer := make([]int, 0)
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
		if visited[toId] {
			return layerIdx
		}
	}
	return -1
}

func BfsQueue[T comparable](g data.Graph[T], f T, t T) int {
	fromId := g.GetNodeId(f)
	toId := g.GetNodeId(t)
	visited := make(map[int]bool)
	visited[fromId] = true
	queue := data.NewQueue[int]()
	queue.Push(fromId)
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
		if visited[toId] {
			return steps
		}
	}
}

func DfsStack[T comparable](g data.Graph[T], f T, t T) int {
	fromId := g.GetNodeId(f)
	toId := g.GetNodeId(t)
	explored := make(map[int]bool)
	stack := data.NewStack[int]()
	stack.Push(fromId)
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
				if edge.To() == toId {
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
