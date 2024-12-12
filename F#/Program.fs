open System
open System.Collections.Generic

// Define a Node type
type Node<'T>(data: 'T) =
    member val Data = data with get, set
    override this.ToString() = sprintf "%A" this.Data

// Define a Graph type
type Graph<'T>() =
    // A dictionary to store adjacency lists
    let adjacencyList = Dictionary<Node<'T>, ResizeArray<Node<'T>>>()

    // Add a new node to the graph
    member this.AddNode(data: 'T) =
        let node = Node(data)

        if not (adjacencyList.ContainsKey(node)) then
            adjacencyList.[node] <- ResizeArray<Node<'T>>()

        node

    // Add a directed edge between two nodes
    member this.AddEdge(fromNode: Node<'T>, toNode: Node<'T>) =
        if adjacencyList.ContainsKey(fromNode) then
            adjacencyList.[fromNode].Add(toNode)
        else
            failwith "Source node not found in the graph."

    // Get neighbors of a given node
    member this.GetNeighbors(node: Node<'T>) =
        if adjacencyList.ContainsKey(node) then
            adjacencyList.[node]
        else
            failwith "Node not found in the graph."

    // Print the graph for visualization
    member this.PrintGraph() =
        for KeyValue(node, neighbors) in adjacencyList do
            let neighborsData =
                neighbors |> Seq.map (fun n -> n.Data.ToString()) |> String.concat ", "

            printfn "%A -> [%s]" node.Data neighborsData

// Example usage
let graph = Graph<string>()

// Add nodes
let alice = graph.AddNode("Alice")
let bob = graph.AddNode("Bob")
let charlie = graph.AddNode("Charlie")

// Add edges
graph.AddEdge(alice, bob) // Alice -> Bob
graph.AddEdge(bob, charlie) // Bob -> Charlie
graph.AddEdge(alice, charlie) // Alice -> Charlie

// Get neighbors
printfn "Neighbors of Alice: %A" (graph.GetNeighbors(alice) |> Seq.map (fun n -> n.Data) |> Seq.toList)

// Print the graph
graph.PrintGraph()
