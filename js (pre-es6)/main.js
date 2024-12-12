/**
 * Creates a directed graph.
 * @returns {Object} The graph object with methods to add nodes, add edges, find neighbors, and print the graph.
 */
function createGraph() {
  var nodes = [];
  var edges = new Map();

  /**
   * Adds a node to the graph.
   * @param {*} data - The data to store in the node.
   * @returns {Object} The newly created node.
   */
  function addNode(data) {
    var node = { data: data };
    nodes.push(node);
    edges.set(node, []);
    return node;
  }

  /**
   * Adds a directed edge between two nodes.
   * @param {Object} fromNode - The starting node.
   * @param {Object} toNode - The ending node.
   */
  function addEdge(fromNode, toNode) {
    if (!edges.has(fromNode) || !edges.has(toNode)) {
      throw new Error("Both nodes must exist in the graph.");
    }
    edges.get(fromNode).push(toNode);
  }

  /**
   * Gets the neighbors of a given node.
   * @param {Object} node - The node to find neighbors for.
   * @returns {Object[]} An array of neighboring nodes.
   */
  function getNeighbors(node) {
    if (!edges.has(node)) {
      throw new Error("Node does not exist in the graph.");
    }
    return edges.get(node);
  }

  /**
   * Prints the graph structure.
   */
  function printGraph() {
    nodes.forEach(function (node) {
      var neighborData = edges.get(node).map(function (neighbor) {
        return neighbor.data;
      });
      console.log(node.data + " -> [" + neighborData.join(", ") + "]");
    });
  }

  // Expose the methods of the graph
  return {
    addNode: addNode,
    addEdge: addEdge,
    getNeighbors: getNeighbors,
    printGraph: printGraph
  };
}

// Example Usage
var graph = createGraph();

// Add nodes
var alice = graph.addNode("Alice");
var bob = graph.addNode("Bob");
var charlie = graph.addNode("Charlie");

// Add edges
graph.addEdge(alice, bob);       // Alice -> Bob
graph.addEdge(bob, charlie);    // Bob -> Charlie
graph.addEdge(alice, charlie);  // Alice -> Charlie

// Get neighbors
console.log("Neighbors of Alice:", graph.getNeighbors(alice).map(function (n) { return n.data; }));

// Print the graph
graph.printGraph();

