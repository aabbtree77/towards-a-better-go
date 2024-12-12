type Node<'T> = { Id: int; Data: 'T; Edges: Edges<'T> }

and Edges<'T> =
    | NoEdges
    | Edge of Node<'T> * Edges<'T>

type Graph<'T> =
    | Empty
    | Graph of Node<'T>

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
        let newNode =
            { Id = generateId ()
              Data = data
              Edges = NoEdges }

        match graph with
        | Empty -> (Graph newNode, newNode)
        | Graph root -> (Graph root, newNode)

    // Add a directed edge between two nodes
    let rec addEdge (fromNode: Node<'T>) (toNode: Node<'T>) (graph: Graph<'T>) : Graph<'T> =
        let rec updateNode currentNode =
            if currentNode.Id = fromNode.Id then
                { currentNode with Edges = Edge(toNode, currentNode.Edges) }
            else
                match currentNode.Edges with
                | NoEdges -> currentNode
                | Edge (next, rest) -> { currentNode with Edges = Edge(updateNode next, rest) }

        match graph with
        | Empty -> graph
        | Graph root -> Graph(updateNode root)

    // Get neighbors of a node from the entire graph
    let rec getNeighborsFromGraph (node: Node<'T>) (graph: Graph<'T>) : Node<'T> list =
        let rec findNode currentNode =
            if currentNode.Id = node.Id then
                let rec collectNeighbors edges acc =
                    match edges with
                    | NoEdges -> acc
                    | Edge (next, rest) -> collectNeighbors rest (next :: acc)

                collectNeighbors currentNode.Edges []
            else
                match currentNode.Edges with
                | NoEdges -> []
                | Edge (next, rest) ->
                    match findNode next with
                    | [] -> findNode { currentNode with Edges = rest }
                    | result -> result

        match graph with
        | Empty -> []
        | Graph root -> findNode root



    // Print the graph
    let rec printNode (node: Node<'T>) (visited: Set<int>) : Set<int> =
        if Set.contains node.Id visited then
            visited
        else
            let newVisited = Set.add node.Id visited
            printfn "Node %d: %A" node.Id node.Data
            printfn "  Edges:"

            let rec printEdges edges =
                match edges with
                | NoEdges -> ()
                | Edge (n, r) ->
                    printfn "    -> Node %d" n.Id
                    printEdges r

            printEdges node.Edges

            let rec traverseEdges edges visited =
                match edges with
                | NoEdges -> visited
                | Edge (next, rest) ->
                    let updatedVisited = printNode next visited
                    traverseEdges rest updatedVisited

            traverseEdges node.Edges newVisited

    let printGraph (graph: Graph<'T>) : unit =
        match graph with
        | Empty -> printfn "Graph is empty."
        | Graph root -> ignore (printNode root Set.empty)

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
let neighbors = Graph.getNeighborsFromGraph alice finalGraph
printfn "Neighbors of Alice: %A" (neighbors |> List.map (fun n -> n.Data))



// Print the graph
Graph.printGraph finalGraph
