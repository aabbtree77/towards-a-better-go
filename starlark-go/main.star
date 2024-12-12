# Install Go.
# Install Starlark-Go:
# go install go.starlark.net/cmd/starlark@latest
# ~/go/bin/starlark main.star

def create_graph():
    """Create a new empty directed graph."""
    return {}

def add_node(graph, node):
    """
    Add a node to the graph.
    Returns the updated graph.
    """
    if not graph.get(node):
        graph[node] = {}  # Initialize the node with an empty dict as its "set"
    return graph

def add_edge(graph, from_node, to_node):
    """
    Add a directed edge from `from_node` to `to_node`.
    Returns the updated graph.
    """
    if not graph.get(from_node):
        graph = add_node(graph, from_node)
    if not graph.get(to_node):
        graph = add_node(graph, to_node)
    graph[from_node][to_node] = None
    return graph

def get_neighbors(graph, node):
    """
    Get the neighbors of a node as a list.
    If the node doesn't exist, return an empty list.
    """
    neighbors = graph.get(node, {})
    return list(neighbors.keys())

def print_graph(graph):
    """
    Print the graph as an adjacency list.
    """
    for node, neighbors in graph.items():
        neighbor_list = ", ".join(neighbors.keys())
        print("%s -> %s" % (node, neighbor_list))
        
graph = create_graph()

# Add nodes
graph2 = add_node(graph, "Alice")
graph3 = add_node(graph2, "Bob")
graph4 = add_node(graph3, "Charlie")

# Add edges
graph5 = add_edge(graph4, "Alice", "Bob")
graph6 = add_edge(graph5, "Bob", "Charlie")
graph7 = add_edge(graph6, "Alice", "Charlie")

# Get neighbors
alice_neighbors = get_neighbors(graph7, "Alice")
print("Neighbors of Alice:", alice_neighbors)

# Print the entire graph
print_graph(graph7)
        
# Updating graph directly, e.g. graph = add_node(graph, "Alice"), would give this error:
# ~/go/bin/starlark test_generic_directed_graph.star
# test_generic_directed_graph.star:51:1: cannot reassign global graph declared at test_generic_directed_graph.star:42:1
# Solution: introduce new variables after each graph update as shown in the code above.

# Output:
# Neighbors of Alice: ["Bob", "Charlie"]
# Alice -> Bob, Charlie
# Bob -> Charlie
# Charlie -> 

