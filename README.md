## ChatGPT Prompt

Create a directed graph in X. Nodes (vertices) can hold arbitrary data. Edges define relationships between nodes. The graph supports generic operations like adding nodes, adding edges, and finding neighbors. The example graph use in TypeScript:

```typescript
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
```

## TypeScript, Python3, Go, Rust

ChatGPT writes correct, human-readable code. There might be some deep null-safety issues, but they can be revealed only with a heavy testing. You have to ask it for specific patterns, idioms, safety. Iterations improve the results, just like in a human coding.

## Clojure, Starlark-Go, F#

ChatGPT does not produce correct code on the first run, but it self-corrects when given another prompt or two based on the compiler/runtime feedback. `Program_evenmore_idiomatic.fs` took about five iterations to get it right. I prompted it not to rely on any maps and lists, use recursive data structures, recursive functions, immutability. The latter is not easy at all, I fear we are all overreacting with it.

## Roc and Gleam

ChatGPT totally hallucinates here, does not even warn a user that these languages have changed since the last time ChatGPT saw them. It does not browse and self-update.

One could mention here [Borgo](https://github.com/borgo-lang/borgo) and [Inko](https://github.com/inko-lang/inko) (both written in Rust), as well as [Grain](https://github.com/grain-lang/grain) (written in Reason), but I fear the results would be similar. However, the idea behind these languages is very clear and makes a lot of sense: Remove Rust's compile time complexity, keep only the good bits. 
 
