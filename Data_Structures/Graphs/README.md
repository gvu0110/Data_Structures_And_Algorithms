# Graphs

- Graph is simply a set of values that are related in a pairwise fashion

### Types of graphs
- Directed: has edges with direction
- Undirected: has edges that do not have a direction

- Weighted: has weight on each edge (a **weight** is a numerical value attached to each individual edge in the graph)
- Unweighted: doesn't have weight on each edge

- Cyclic: has cycles, which means we can start from some node and follow a path such that you arrive at the same node we began
- Acyclic: has no cycles

### Graph data
```
   2 --- 0
 /   \
1 --- 3

// Edge List: show the connection, e.g. 0 is connected to 2, 2 is connected to 3
graph = [[0,2], [2,1], [2,3], [3,1]]

// Adjacent List: the index is the node and the value is the node neighbors
graph = [2, [2,3], [0,1,3], [1,2]]

// Adjacent Matricx: have 0s and 1s indicating whether the Node X is connected to Node Y
graph = [
    [0,0,1,0],
    [0,0,1,1],
    [1,1,0,1],
    [0,1,1,0]
]
```

### Graphs review
Pros:
- Relationships

Cons:
- Scaling is hard
