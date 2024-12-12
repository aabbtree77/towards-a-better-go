package main

import (
	"fmt"
)

// Node represents a single vertex in the graph and its edges
type Node struct {
	data     string
	neighbors []Node // Store actual neighbor nodes
}

// Graph represents a directed graph with string data in nodes
type Graph struct {
	nodes []Node
}

// NewGraph creates a new empty graph
func NewGraph() Graph {
	return Graph{
		nodes: []Node{},
	}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(data string) int {
	node := Node{
		data:     data,
		neighbors: []Node{},
	}
	g.nodes = append(g.nodes, node)
	return len(g.nodes) - 1
}

// AddEdge creates a directed edge from one node to another
func (g *Graph) AddEdge(from, to int) {
	g.nodes[from].neighbors = append(g.nodes[from].neighbors, g.nodes[to])
}

// GetNeighbors returns the neighbors of a given node
func (g *Graph) GetNeighbors(index int) []Node {
	return g.nodes[index].neighbors
}

// PrintGraph prints the graph in a readable format
func (g *Graph) PrintGraph() {
	for i, node := range g.nodes {
		fmt.Printf("Node %d (%v):\n", i, node.data)
		for _, neighbor := range node.neighbors {
			fmt.Printf("  -> %v\n", neighbor.data)
		}
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add nodes
	alice := graph.AddNode("Alice")
	bob := graph.AddNode("Bob")
	charlie := graph.AddNode("Charlie")

	// Add edges
	graph.AddEdge(alice, bob)       // Alice -> Bob
	graph.AddEdge(bob, charlie)    // Bob -> Charlie
	graph.AddEdge(alice, charlie)  // Alice -> Charlie

	// Get neighbors
	neighbors := graph.GetNeighbors(alice)
	fmt.Printf("Neighbors of Alice: ")
	for _, neighbor := range neighbors {
		fmt.Printf("%v ", neighbor.data)
	}
	fmt.Println()

	// Print the graph
	graph.PrintGraph()
}

