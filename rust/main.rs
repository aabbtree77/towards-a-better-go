use std::collections::{HashMap, HashSet};
use std::fmt::Debug;

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Node<T> {
    id: usize,
    data: T,
}

impl<T> Node<T> {
    fn new(id: usize, data: T) -> Self {
        Node { id, data }
    }
}

struct Graph<T> {
    nodes: HashMap<usize, Node<T>>, // Store nodes by their IDs
    edges: HashMap<usize, HashSet<usize>>, // Adjacency list representation
    next_id: usize, // Auto-incrementing ID for new nodes
}

impl<T> Graph<T>
where
    T: Debug + Clone + PartialEq + Eq,
{
    fn new() -> Self {
        Graph {
            nodes: HashMap::new(),
            edges: HashMap::new(),
            next_id: 0,
        }
    }

    // Add a new node and return a reference to it
    fn add_node(&mut self, data: T) -> usize {
        let id = self.next_id;
        self.next_id += 1;
        self.nodes.insert(id, Node::new(id, data));
        self.edges.entry(id).or_insert_with(HashSet::new);
        id
    }

    // Add a directed edge from `from` to `to`
    fn add_edge(&mut self, from: usize, to: usize) {
        if self.nodes.contains_key(&from) && self.nodes.contains_key(&to) {
            self.edges.entry(from).or_insert_with(HashSet::new).insert(to);
        }
    }

    // Get neighbors of a node by its ID
    fn get_neighbors(&self, id: usize) -> Vec<&Node<T>> {
        if let Some(neighbors) = self.edges.get(&id) {
            neighbors
                .iter()
                .filter_map(|&neighbor_id| self.nodes.get(&neighbor_id))
                .collect()
        } else {
            vec![]
        }
    }

    // Print the graph
    fn print_graph(&self) {
        for (id, node) in &self.nodes {
            print!("Node {:?}: ", node.data);
            if let Some(neighbors) = self.edges.get(id) {
                let neighbor_data: Vec<_> = neighbors
                    .iter()
                    .filter_map(|&neighbor_id| self.nodes.get(&neighbor_id))
                    .map(|node| &node.data)
                    .collect();
                println!("{:?}", neighbor_data);
            } else {
                println!("[]");
            }
        }
    }
}

fn main() {
    let mut graph = Graph::new();

    // Add nodes
    let alice = graph.add_node("Alice".to_string());
    let bob = graph.add_node("Bob".to_string());
    let charlie = graph.add_node("Charlie".to_string());

    // Add edges
    graph.add_edge(alice, bob);       // Alice -> Bob
    graph.add_edge(bob, charlie);    // Bob -> Charlie
    graph.add_edge(alice, charlie);  // Alice -> Charlie

    // Get neighbors
    let neighbors_of_alice = graph.get_neighbors(alice);
    println!("Neighbors of Alice: {:?}", neighbors_of_alice.iter().map(|n| &n.data).collect::<Vec<_>>());

    // Print the graph
    graph.print_graph();
}

