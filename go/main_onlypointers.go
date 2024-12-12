package main

import (
	"fmt"
)

// Node represents a single vertex in the graph and its edges
type Node struct {
	data     string
	next     *Node // Next node in the graph
	neighbors *Edge // Linked list of neighbors
}

// Edge represents an edge in the graph
// Each edge points to a neighbor node and the next edge

type Edge struct {
	target *Node
	next   *Edge
}

// Graph represents a directed graph with linked list structure
type Graph struct {
	head *Node // Head of the list of nodes
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(data string) *Node {
	newNode := &Node{
		data:     data,
		neighbors: nil,
		next:     g.head,
	}
	g.head = newNode
	return newNode
}

// AddEdge creates a directed edge from one node to another
func (g *Graph) AddEdge(from, to *Node) {
	newEdge := &Edge{
		target: to,
		next:   from.neighbors,
	}
	from.neighbors = newEdge
}

// GetNeighbors returns the neighbors of a given node as a slice of pointers
func (g *Graph) GetNeighbors(node *Node) []*Node {
	var neighbors []*Node
	edge := node.neighbors
	for edge != nil {
		neighbors = append(neighbors, edge.target)
		edge = edge.next
	}
	return neighbors
}

// PrintGraph prints the graph in a readable format
func (g *Graph) PrintGraph() {
	node := g.head
	for node != nil {
		fmt.Printf("Node (%v):\n", node.data)
		edge := node.neighbors
		for edge != nil {
			fmt.Printf("  -> %v\n", edge.target.data)
			edge = edge.next
		}
		node = node.next
	}
}

func main() {
	// Create a new graph
	graph := &Graph{}

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

