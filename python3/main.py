from collections import defaultdict

class DirectedGraph:
    def __init__(self):
        # The graph is represented as an adjacency list (a dictionary of sets)
        self.graph = defaultdict(set)

    def add_node(self, value):
        """Add a node to the graph if it doesn't already exist."""
        if value not in self.graph:
            self.graph[value]  # This initializes the node with an empty set of neighbors
        return value

    def add_edge(self, from_node, to_node):
        """Add a directed edge from `from_node` to `to_node`."""
        self.graph[from_node].add(to_node)

    def get_neighbors(self, node):
        """Get the neighbors of a node as a list."""
        return list(self.graph.get(node, []))

    def print_graph(self):
        """Print the graph as an adjacency list."""
        for node, neighbors in self.graph.items():
            neighbor_list = ", ".join(map(str, neighbors))
            print(f"{node} -> {neighbor_list}")

# Example usage
if __name__ == "__main__":
    # Create a new graph
    graph = DirectedGraph()

    # Add nodes
    alice = graph.add_node("Alice")
    bob = graph.add_node("Bob")
    charlie = graph.add_node("Charlie")

    # Add edges
    graph.add_edge(alice, bob)
    graph.add_edge(bob, charlie)
    graph.add_edge(alice, charlie)

    # Get neighbors of a node
    alice_neighbors = graph.get_neighbors(alice)
    print(f"Neighbors of Alice: {alice_neighbors}")

    # Print the whole graph
    graph.print_graph()

