package main

import (
	"fmt"
)

// Node represents a single vertex in the graph
type Node[T any] struct {
	data T
	id   int
}

// Graph represents a directed graph with generic data in nodes
type Graph[T any] struct {
	nodes map[int]*Node[T]
	edges map[int][]*Node[T]
	idSeq int
}

// NewGraph creates a new empty graph
func NewGraph[T any]() *Graph[T] {
	return &Graph[T]{
		nodes: make(map[int]*Node[T]),
		edges: make(map[int][]*Node[T]),
		idSeq: 0,
	}
}

// AddNode adds a new node to the graph
func (g *Graph[T]) AddNode(data T) *Node[T] {
	node := &Node[T]{
		data: data,
		id:   g.idSeq,
	}
	g.nodes[g.idSeq] = node
	g.idSeq++
	return node
}

// AddEdge creates a directed edge from one node to another
func (g *Graph[T]) AddEdge(from, to *Node[T]) {
	g.edges[from.id] = append(g.edges[from.id], to)
}

// GetNeighbors returns the neighbors of a given node
func (g *Graph[T]) GetNeighbors(node *Node[T]) []*Node[T] {
	return g.edges[node.id]
}

// PrintGraph prints the graph in a readable format
func (g *Graph[T]) PrintGraph() {
	for id, node := range g.nodes {
		fmt.Printf("Node %d (%v):\n", id, node.data)
		for _, neighbor := range g.edges[id] {
			fmt.Printf("  -> %v\n", neighbor.data)
		}
	}
}

func main() {
	// Create a new graph
	graph := NewGraph[string]()

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

