package main

import (
	"fmt"
	"strings"
)

//type Node struct {
//	value    string
//	searched bool
//}

type Graph struct {
	graph map[string][]string
}

func NewGraph() Graph {
	return Graph{graph: make(map[string][]string)}
}

func main() {
	g := NewGraph()
	g.graph["A"] = []string{}
	g.graph["B"] = []string{}
	g.graph["C"] = []string{}
	g.graph["D"] = []string{}

	g.graph["A"] = append(g.graph["A"], "B")
	g.graph["A"] = append(g.graph["A"], "C")
	g.graph["C"] = append(g.graph["C"], "E")
	g.graph["B"] = append(g.graph["B"], "D")
	g.graph["A"] = append(g.graph["A"], "E")
	fmt.Println(g)
	g.BFS("A", "D")

}

func (g Graph) BFS(startNode string, target string) {
	var queue []string
	searchedNode := g.createSearched()
	var path []string
	queue = append(queue, startNode)
	searchedNode[startNode] = true

	for len(queue) > 0 {
		var current string
		current, queue = queue[0], queue[1:]

		path = append(path, current)
		edges := g.graph[current]
		if contains(edges, target) {
			path = append(path, target)
			fmt.Println(strings.Join(path, "->"))
		}
		for _, node := range g.graph[current] {
			if !searchedNode[node] {
				queue = append(queue, node)
				searchedNode[node] = true
			}
		}
	}
}

func (g Graph) createSearched() map[string]bool {
	searched := make(map[string]bool, len(g.graph))

	for key := range g.graph {
		searched[key] = false
	}
	return searched
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
