package main

import (
	"fmt"
	"strings"
)

// region Some idea for future modernization
//type Node struct {
//	value    string
//	searched bool
//}
//endregion

// region Структура
type Graph struct {
	graph map[string][]string
}

//endregion

// region Конструктор
func NewGraph() Graph {
	return Graph{graph: make(map[string][]string)}
}

//endregion

func main() {
	g := NewGraph()
	g.graph["A"] = []string{}
	g.graph["B"] = []string{}
	g.graph["C"] = []string{}
	g.graph["D"] = []string{}
	g.graph["E"] = []string{}

	g.graph["A"] = append(g.graph["A"], "B")
	g.graph["B"] = append(g.graph["B"], "A")
	g.graph["A"] = append(g.graph["A"], "C")
	g.graph["C"] = append(g.graph["C"], "A")
	g.graph["A"] = append(g.graph["A"], "E")
	g.graph["E"] = append(g.graph["E"], "A")
	g.graph["C"] = append(g.graph["C"], "E")
	g.graph["E"] = append(g.graph["E"], "C")
	g.graph["B"] = append(g.graph["B"], "D")
	g.graph["D"] = append(g.graph["D"], "B")

	fmt.Println(g)

	fmt.Println(g.BFS("A", "D"))
}

//region Методы

// BFS finds the shortest path, based on the number of nodes, from the starting node to the target node.
func (g Graph) BFS(startNode string, target string) (result string) {

	//Queue is a slice with all neighbors of a node. Works according to FIFO
	var queue []string

	//SearchedNode map that contains data about the node verification. SearchedNode map[string]bool.
	//Key - string, name of node. Value - bool, if node is already checked - true.
	//If a node has already been checked, it does not need to be checked again.
	searchedNode := g.createSearched()

	//Path is a path from start node to target node. The variable is written to the string result
	//and returned from the function.
	var path []string
	queue = append(queue, startNode)
	searchedNode[startNode] = true

	//This loop continues until the queue is not empty.
	//Nodes-neighbors of the current node are added to the queue, provided that they have not yet been encountered.
	for len(queue) > 0 {

		//Current - current node
		var current string
		current, queue = queue[0], queue[1:]
		path = append(path, current)

		//Edges are the edges of the current node. List of nodes that form a pair together with the current node.
		edges := g.graph[current]
		if contains(edges, target) {
			path = append(path, target)
			result = fmt.Sprint(strings.Join(path, "->"))
			return result
		}

		for _, node := range g.graph[current] {
			if !searchedNode[node] {
				queue = append(queue, node)
				searchedNode[node] = true
			}
		}
	}
	return result
}

// CreateSearched creates a map of searched nodes.
func (g Graph) createSearched() map[string]bool {
	searched := make(map[string]bool, len(g.graph))

	for key := range g.graph {
		searched[key] = false
	}
	return searched
}

//endregion

//region Функции

// Contains creates a set in which the keys of neighbor-nodes are paired with the current node.
// Using the set, it is checked whether the target node exists. Contains returns true, the target node
// is present among the set keys
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	fmt.Println(set)

	_, ok := set[item]
	return ok
}

//endregion
