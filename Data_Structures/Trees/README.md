# Trees

### Binary tree
- Each node has only either 0, 1 or 2 nodes
- Each child has only 1 parent

### Binary search tree (BST)
BST Rules
- Each node of a BST may only have exactly 0, 1, or 2 children.
- For each node X, all nodes in its left subtree have a value less than X. Thus, all nodes in its right subtree have a value greater than X. The inverse is also true.
- Both the left and right subtrees of X must also be BSTs.

- Lookup: O(log n)
- Insert: O(log n)
- Delete: O(log n)

### Balanced vs Unbalanced BST
- A BST is balanced if any two sibling subtrees do not differ in height by more than one level.
- Unbalanced BST turns into a linked list:
+ Lookup: O(n)
+ Insert: O(n)
+ Delete: O(n)

### BST Pros and Cons:
Pros:
- Better than O(n)
- Ordered
- Flexible size
Cons:
- No O(1) operations

### AVL and Red-Black tree
- An AVL tree is a self-balancing BST. In an AVL tree, the heights of the two child subtrees of any node differ by at most one; if at any time they differ by more than one, rebalancing is done to restore this property.

How it works: [https://medium.com/basecs/the-little-avl-tree-that-could-86a3cae410c7](https://medium.com/basecs/the-little-avl-tree-that-could-86a3cae410c7)

- A red-black tree is a type of self-balancing binary search tree, very similar to other self-balancing trees, such as AVL trees. However, a red-black tree is a structure that has to adhere to a very strict set of rules in order to make sure that it stays balanced; the rules of a red-black tree are exactly what enables it to maintain a guaranteed logarithmic time complexity.

Generally speaking, the four rules of a red-black tree are always presented in the same order, as follows:
1. Every single node in the tree must be either red or black.
2. The root node of the tree must **always** be black.
3. Two red nodes can never appear consecutively, one after another; a red node must always be preceded by a black node (it must have a black parent node), and a red node must always have black children nodes.
4. Every branch path — the path from a root node to an empty (null) leaf node — must pass through the exact same number of black nodes. A branch path from the root to an empty leaf node is also known as an unsuccessful search path, since it represents the path we would take if we were to search for a node that didn’t exist within the tree.

How it works: [https://medium.com/basecs/painting-nodes-black-with-red-black-trees-60eacb2be9a5](https://medium.com/basecs/painting-nodes-black-with-red-black-trees-60eacb2be9a5)

### Binary Heap
A Binary Heap is a Binary Tree with following properties.
1) It’s a complete tree (All levels are completely filled except possibly the last level and the last level has all keys as left as possible). This property of Binary Heap makes them suitable to be stored in an array.

2) A Binary Heap is either Min Heap or Max Heap. In a Min Binary Heap, the key at root must be minimum among all keys present in Binary Heap. The same property must be recursively true for all nodes in Binary Tree. Max Binary Heap is similar to MinHeap.
- Lookup: O(n)
- Insert: O(log n)
- Delete: O(log n)

Pros:
- Priority
- Flexible size
- Fast insert

Cons:
- Slow lookup

### Priority Queue
- Because we do left-to-right insertion, there is no concept of an unbalanced Binary Heap like the BST
- The only guarantee that Binary Heap gives us is the parent if always greater (or smaller) than the children
- Priority Queue is type of of data where each element has a priority, and elements with higher priority are served before elements with lower priority

### Trie
- Is a specialized tree used in searching most often with text
- Trie allows you to know if a word or part of a word exists in a body of text
- Lookup: O(length of the word)