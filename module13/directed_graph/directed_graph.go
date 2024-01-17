package main

import (
	"fmt"
	"math"
)

// region Структура
type Directed_Graph struct {
	nodes []int
	edges map[int]map[int]int
}

//endregion

// region Конуструктор
func NewDirected_Graph() *Directed_Graph {
	return &Directed_Graph{nodes: []int{}, edges: make(map[int]map[int]int)}
}

//endregion

//region Методы

func (g *Directed_Graph) AddNode(node int) {
	g.nodes = append(g.nodes, node)
}

func (g *Directed_Graph) AddEdge(startNode, endNode, weight int) {
	if _, ok := g.edges[startNode]; !ok {
		g.edges[startNode] = make(map[int]int)
	}
	g.edges[startNode][endNode] = weight
}

// Dijkstra finds the shortest path relative to the weights of the graph nodes.
// Returns a map where the key is the endnode and the value is the shortest path.
func (g *Directed_Graph) dijkstra(start int) map[int]int {
	//Distances is the result of this method. It is map[int]int. Key - endnode, value - shortest way from start node.
	distances := g.getInfinite()
	distances[start] = 0

	//Visited map[int]bool. Key - string, number of node. Value - bool, if node is already checked - true.
	//If a node has already been checked, it does not need to be checked again.
	visited := make(map[int]bool)
	//The loop will run until all nodes are processed.
	for len(visited) < len(g.nodes) {
		node := g.minDistance(distances, visited)
		visited[node] = true

		//Loop goes through all neighbors of a node and updates the path to them.
		for neighbor, weight := range g.edges[node] {
			newDistance := distances[node] + weight
			if newDistance < distances[neighbor] {
				distances[neighbor] = newDistance
			}
		}
	}
	return distances
}

// GetInfinite returns map[int]int. Key - node, value - max int32 value.
func (g *Directed_Graph) getInfinite() map[int]int {
	distances := make(map[int]int)
	for _, node := range g.nodes {
		//Initially, all ways are equal to infinity. In this case, the maximum number of type int32.
		distances[node] = math.MaxInt32
	}
	return distances
}

// MinDistance returns min node.
func (g *Directed_Graph) minDistance(distances map[int]int, visited map[int]bool) int {
	min := math.MaxInt32
	var minNode int
	for node, distance := range distances {
		if distance < min && !visited[node] {
			min = distance
			minNode = node
		}
	}
	return minNode
}

//endregion

func main() {
	graph := NewDirected_Graph()

	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)
	graph.AddNode(5)

	graph.AddEdge(1, 2, 10)
	graph.AddEdge(1, 4, 30)
	graph.AddEdge(1, 5, 100)
	graph.AddEdge(2, 3, 50)
	graph.AddEdge(3, 5, 10)
	graph.AddEdge(4, 3, 20)
	graph.AddEdge(4, 5, 60)

	distances := graph.dijkstra(1)
	fmt.Println(distances)

}
