type Node<'T> =
    { Id: int; Data: 'T }

// Define an edge as a directed connection between two nodes
type Edge<'T> =
    { From: Node<'T>; To: Node<'T> }

// Define the graph with a list of nodes and edges
type Graph<'T> =
    | Empty
    | Graph of nodes: Node<'T> list * edges: Edge<'T> list

module Graph =

    // Helper to create unique IDs for nodes
    let private generateId =
        let counter = ref 0
        fun () ->
            let id = !counter
            counter := id + 1
            id

    // Add a node to the graph
    let addNode (data: 'T) (graph: Graph<'T>) : Graph<'T> * Node<'T> =
        let newNode = { Id = generateId (); Data = data }
        match graph with
        | Empty -> (Graph ([newNode], []), newNode)
        | Graph (nodes, edges) -> (Graph (newNode :: nodes, edges), newNode)

    // Add a directed edge between two nodes
    let addEdge (fromNode: Node<'T>) (toNode: Node<'T>) (graph: Graph<'T>) : Graph<'T> =
        match graph with
        | Empty -> graph // Cannot add edges to an empty graph
        | Graph (nodes, edges) ->
            let newEdge = { From = fromNode; To = toNode }
            Graph (nodes, newEdge :: edges)

    // Get neighbors of a node
    let getNeighbors (node: Node<'T>) (graph: Graph<'T>) : Node<'T> list =
        match graph with
        | Empty -> []
        | Graph (_, edges) ->
            edges
            |> List.filter (fun edge -> edge.From = node)
            |> List.map (fun edge -> edge.To)

    // Print the graph
    let printGraph (graph: Graph<'T>) : unit =
        match graph with
        | Empty ->
            printfn "Graph is empty."
        | Graph (nodes, edges) ->
            printfn "Nodes:"
            nodes |> List.iter (fun node -> printfn "  Node %d: %A" node.Id node.Data)
            printfn "Edges:"
            edges |> List.iter (fun edge -> printfn "  %A -> %A" edge.From.Data edge.To.Data)

// Example usage
let initialGraph = Empty

// Add nodes
let graphWithAlice, alice = Graph.addNode "Alice" initialGraph
let graphWithBob, bob = Graph.addNode "Bob" graphWithAlice
let graphWithCharlie, charlie = Graph.addNode "Charlie" graphWithBob

// Add edges
let graphWithEdge1 = Graph.addEdge alice bob graphWithCharlie
let graphWithEdge2 = Graph.addEdge bob charlie graphWithEdge1
let finalGraph = Graph.addEdge alice charlie graphWithEdge2

// Get neighbors of Alice
let neighbors = Graph.getNeighbors alice finalGraph
printfn "Neighbors of Alice: %A" (neighbors |> List.map (fun n -> n.Data))

// Print the graph
Graph.printGraph finalGraph

