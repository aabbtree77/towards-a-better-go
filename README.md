## Introduction

How good is your programming language w.r.t. ChatGPT?

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

## Results

### TypeScript, Python3, Go, Rust

ChatGPT writes correct, human-readable code. 

Js/Es6 with JSDoc are not as good as TypeScript. We only get `*` (Any), `Object`, and `Object[]` in JSDoc.

TypeScript >=v2/0 allows to make everything non-nullable by default with the type `never` providing exhaustive switch. ChatGPT did not use any of these concepts, opting for a questionable no-null assertion with `!`: 

```
this.nodes.get(from)!.add(to);
```

If asked specifically about the type safety here, it does suggest rewriting to either this:

```
if (this.nodes.has(from)) {
    this.nodes.get(from)!.add(to);
} else {
    this.nodes.set(from, new Set([to]));
}
```

or that:

```
const targetSet = this.nodes.get(from) ?? new Set<Node<T>>();
targetSet.add(to);
this.nodes.set(from, targetSet);
```

It starts to care about safety only after you ask it. Null-safety may exist in the language, but it is not guaranteed to appear in the code. Python3 and Go are even worse safety-wise.

Note that ChatGPT groks pointers and the borrow checker despite that these are so fragile syntactically. There must be a lot of Go and Rust codes out there to learn from.

### Clojure, Starlark-Go, F#

ChatGPT does not produce correct code on the first run, but it self-corrects from a compiler/runtime feedback.

The languages listed here emphasize immutability in different ways, but this is already tricky. It does not help that there is likely 1000x less code written in these languages than say in Js. ChatGPT already begins to struggle. 

Clojure: 

- Ironically, ChatGPT has missed one obvious parenthesis, so this is fragile. 

- Installing [Java SDK](https://sdkman.io/), VS Code [Calva](https://calva.io/get-started-with-clojure/) and connecting Clojure to its famous power REPL went smoothly. Leiningen vs deps.edn are both fine, but all this feels... complecting. 

Starlark-Go:

- ChatGPT has missed that the global scope variables in Starlark are immutable. Self-corrected after prompting it.

- [Starlark-Go](https://github.com/google/starlark-go) is what Python should have been, a modern SETL derivative with no nonsense. However, it has no batteries... It is also incredible that such a small language without a build system and package manager still gets several hundred github issues. Coding is tricky...

F#:

- `Program.fs`: Errored on a list which had to be `ResizeArray`. Self-corrected after a single feedback.

- `Program_idiomatic.fs`. When asked to write a more idiomatic F#, ChatGPT got confused with immutability, but self-corrected after a single feedback by renaming the graph updates.

- `Program_evenmore_idiomatic.fs`. I have asked it not to rely on any maps and lists, use recursive data structures, recursive functions, and immutability. This took about five cycles of self-correction. I thought this was hopeless, but it did succeed.

### Roc and Gleam

ChatGPT fails here. It totally hallucinates, does not even warn a user that these languages have changed since the last time ChatGPT saw them. It does not browse and self-update, in my experience.

One could mention here [Borgo](https://github.com/borgo-lang/borgo) and [Inko](https://github.com/inko-lang/inko) (both written in Rust), as well as [Grain](https://github.com/grain-lang/grain) (written in Reason), but I fear the results would be similar. However, the idea behind these languages is very clear and makes a lot of sense: Remove Rust's compile time complexity, keep only the good bits. To be revisited.

 
