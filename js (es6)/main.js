/**
 * Class representing a directed graph.
 */
class Graph {
  constructor() {
    /**
     * @type {Set<Object>} Set of nodes in the graph.
     */
    this.nodes = new Set();

    /**
     * @type {Map<Object, Object[]>} Map of nodes to their neighbors.
     */
    this.edges = new Map();
  }

  /**
   * Adds a node to the graph.
   * @param {*} data - The data to store in the node.
   * @returns {Object} The newly created node.
   */
  addNode(data) {
    const node = { data };
    this.nodes.add(node);
    this.edges.set(node, []);
    return node;
  }

  /**
   * Adds a directed edge between two nodes.
   * @param {Object} fromNode - The starting node.
   * @param {Object} toNode - The ending node.
   */
  addEdge(fromNode, toNode) {
    if (!this.edges.has(fromNode) || !this.edges.has(toNode)) {
      throw new Error("Both nodes must exist in the graph.");
    }
    this.edges.get(fromNode).push(toNode);
  }

  /**
   * Gets the neighbors of a given node.
   * @param {Object} node - The node to find neighbors for.
   * @returns {Object[]} An array of neighboring nodes.
   */
  getNeighbors(node) {
    if (!this.edges.has(node)) {
      throw new Error("Node does not exist in the graph.");
    }
    return this.edges.get(node);
  }

  /**
   * Prints the graph structure.
   */
  printGraph() {
    for (const node of this.nodes) {
      const neighborData = this.edges.get(node).map(neighbor => neighbor.data);
      console.log(`${node.data} -> [${neighborData.join(", ")}]`);
    }
  }
}

// Example Usage
const graph = new Graph();

// Add nodes
const alice = graph.addNode("Alice");
const bob = graph.addNode("Bob");
const charlie = graph.addNode("Charlie");

// Add edges
graph.addEdge(alice, bob);       // Alice -> Bob
graph.addEdge(bob, charlie);    // Bob -> Charlie
graph.addEdge(alice, charlie);  // Alice -> Charlie

// Get neighbors
console.log("Neighbors of Alice:", graph.getNeighbors(alice).map(n => n.data));

// Print the graph
graph.printGraph();

