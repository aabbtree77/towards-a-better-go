package main

import (
	"fmt"
)

// Node represents a single vertex in the graph
type Node struct {
	data string
	id   int
}

// Graph represents a directed graph with string data in nodes
type Graph struct {
	nodes map[int]*Node
	edges map[int][]*Node
	idSeq int
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
		edges: make(map[int][]*Node),
		idSeq: 0,
	}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(data string) *Node {
	node := &Node{
		data: data,
		id:   g.idSeq,
	}
	g.nodes[g.idSeq] = node
	g.idSeq++
	return node
}

// AddEdge creates a directed edge from one node to another
func (g *Graph) AddEdge(from, to *Node) {
	g.edges[from.id] = append(g.edges[from.id], to)
}

// GetNeighbors returns the neighbors of a given node
func (g Graph) GetNeighbors(node *Node) []*Node {
	return g.edges[node.id]
}

// PrintGraph prints the graph in a readable format
func (g Graph) PrintGraph() {
	for id, node := range g.nodes {
		fmt.Printf("Node %d (%v):\n", id, node.data)
		for _, neighbor := range g.edges[id] {
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

