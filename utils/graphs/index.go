package graph

import (
	"fmt"
	"strings"

	"github.com/ashinzekene/algorithms/utils/stacks"
)

type Graph[T comparable] struct {
	adjList map[T]map[T]bool
}

func New[T comparable]() *Graph[T] {
	return &Graph[T]{
		adjList: make(map[T]map[T]bool),
	}
}

func (g *Graph[T]) Add(node1 T, node2 T) {
	if g.adjList[node1] == nil {
		g.adjList[node1] = make(map[T]bool)
	}
	g.adjList[node1][node2] = true

	if g.adjList[node2] == nil {
		g.adjList[node2] = make(map[T]bool)
	}
	g.adjList[node2][node1] = true
}

func (g *Graph[T]) Remove(node T) {
	if g.adjList[node] != nil {
		delete(g.adjList, node)
	}

	for key := range g.adjList {
		if g.adjList[key][node] {
			delete(g.adjList[key], node)
		}
	}
}

func (g *Graph[T]) Str() string {
	result := ""
	for key := range g.adjList {
		keys := ""
		for k := range g.adjList[key] {
			keys += fmt.Sprintf("%v,", k)
		}
		
		result += fmt.Sprintf("%v => %s\n", fmt.Sprint(key), strings.TrimSuffix(keys, ","))
	}
	return result
}

func (g *Graph[T]) BFS() []T {
	result := make([]T, 0)
	visited := make(map[T]bool)
	for key := range g.adjList {
		if !visited[key] {
			result = append(result, key)
			visited[key] = true
		}
		for value := range g.adjList[key] {
			if !visited[value] {
				result = append(result, value)
				visited[value] = true
			}
		}
	}
	return result 
}

func (g *Graph[T]) DFS() []T {
	result := make([]T, 0)
	stack := stacks.NewStack[T]()
	visited := make(map[T]bool)

	for k:= range g.adjList {
		stack.Push(k)
		break
	}

	for stack.Length() > 0 {
		current := stack.Pop()
		if visited[current] {
			continue
		}
		result = append(result, current)
		visited[current] = true
		for key := range g.adjList[current] {
			if !visited[key] {
				stack.Push(key)
			}
		}
	}

	return result
}
