// A generic Node type to hold data of any type.
class Node<T> {
    constructor(public data: T) {}
}

// The Graph class is generic over the type of data contained in its nodes.
class Graph<T> {
    private nodes: Map<Node<T>, Set<Node<T>>> = new Map();

    // Add a node to the graph
    addNode(data: T): Node<T> {
        const newNode = new Node(data);
        if (!this.nodes.has(newNode)) {
            this.nodes.set(newNode, new Set());
        }
        return newNode;
    }

    // Add a directed edge from one node to another
    addEdge(from: Node<T>, to: Node<T>): void {
        if (!this.nodes.has(from) || !this.nodes.has(to)) {
            throw new Error("Both nodes must exist in the graph.");
        }
        this.nodes.get(from)!.add(to);
    }

    // Get neighbors of a node
    getNeighbors(node: Node<T>): Node<T>[] {
        if (!this.nodes.has(node)) {
            throw new Error("Node does not exist in the graph.");
        }
        return Array.from(this.nodes.get(node)!);
    }

    // Print the graph
    printGraph(): void {
        for (const [node, neighbors] of this.nodes) {
            const neighborData = Array.from(neighbors).map((n) => n.data);
            console.log(`${node.data} -> ${neighborData.join(", ")}`);
        }
    }
}

// Example Usage:
const graph = new Graph<string>();

// Add nodes
const alice = graph.addNode("Alice");
const bob = graph.addNode("Bob");
const charlie = graph.addNode("Charlie");

// Add edges
graph.addEdge(alice, bob);       // Alice -> Bob
graph.addEdge(bob, charlie);    // Bob -> Charlie
graph.addEdge(alice, charlie);  // Alice -> Charlie

// Get neighbors
console.log("Neighbors of Alice:", graph.getNeighbors(alice).map((n) => n.data));

// Print the graph
graph.printGraph();

