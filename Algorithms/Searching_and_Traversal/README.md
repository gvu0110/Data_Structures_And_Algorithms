# Searching_and_Traversal

### Linear search
- Sequentially checks each element of the list for the target value until a match is found or until all the elements have been searched.

### Binary search
- If the data is sorted, we can do better than linear time with binary search tree
- Time complexity: O(log n)

### Graph + Tree traversal
- We have to visit every single node
- Time complexity: O(n)
- Techniques:
    + Breadth First Search (BFS)
    + Depth First Search (DFS)

### Breadth First Search
- Start with the root node and move left to right across the 2nd level
- Then move left to right across the 3rd level. And so on and so forth
- BFS uses additional memory because it is necessary to track the child nodes of all the nodes on a given level while searching that level.
```
      9
  4       20
1   6  15   170

BFS[9, 4, 20, 1, 6, 15, 170]
```

### Depth First Search
- Follow one branch of the tree down as many levels as possible until the target notice found or the end is reached
- When the search can't go on any further, it continues at the nearest ancestor with an unexplored child
- DFS has a lower memory requirement than BFS because it's not necessary to store all the child pointers at each level
```
      9
  4       20
1   6  15   170

DFS[9, 4, 1, 6, 20, 15, 170]
```

### BFS vs DFS
BFS
- Pros:
    + Shortest path
    + Closer nodes
- Cons:
    + More memory

DFS
- Pros:
    + Less memory
    + Does path exist?
- Cons:
    + Can get slow
```
//If you know a solution is not far from the root of the tree: BFS

//If the tree is very deep and solutions are rare: BFS, because the tree is very deep, DFS will take long

//If the tree is very wide: DFS, because the tree is very wide, BFS might take more memory

//If solutions are frequent but located deep in the tree: DFS

//Determining whether a path exists between two nodes: DFS

//Finding the shortest path: BFS
```

### PreOrder, InOrder, PostOrder
```
      9
  4       20
1   6  15   170

InOrder - Left, Root, Right: [1, 4, 6, 9, 15, 20, 170]
PreOrder - Root, Left, Right: [9, 4, 1, 6, 20, 15, 170]
PostOrder - Left, Right, Root: [1, 6, 4, 15, 170, 20, 9]
```

### Graph Traversal
- BFS: shortest path
    + Determining the shortest path between any note
    + Used in recommendation engines, peer to peer networks, Facebook friend requests, Instagram recommendations
- DFS: check to see if it exists
    + Like solving a maze
    + If you hit a roadblock and you can't go anymore, you backtrack and find a different route
    + Then you keep backtracking until you find the desired node or you exit the maze

### Dijkstra + Bellman-Ford algorithms
- Used for finding the shortest path between 2 nodes
- BFS is great for the shortest path problem, but it is assumes that each path has the same weight
- Bellman-Ford is very effective at solving the shortest path over Dijkstra's algorithm because it can accommodate negative weights, which Dijkstra can't
- When all weights are positive, Dijkstra is faster and more efficient than Bellman-Ford