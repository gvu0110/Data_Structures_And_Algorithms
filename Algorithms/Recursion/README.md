# Recursion

### Recursion vs Iterative
- Anything you do with recursion CAN be done interatively (loop)

Pros:
- Dry
- Readability
Cons:
- Large stack consumes more memory

- Tail call optimization allows recursion to be called without increasing the call stack

### When to use recursion
- Every time you are using a tree or converting something into a tree, consider recusion.
1. Divided into a number of subproblems that are smaller instances of the same problem.
2. Each instance of the subproblem is identical in nature.
3. The solutions of each subproblem can be combined to solve the problem at hand.