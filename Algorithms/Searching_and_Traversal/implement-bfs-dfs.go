package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	right *Node
	left  *Node
	value int
}

type Tree struct {
	root *Node
	size int
}

type Queue struct {
	Elements *list.List
}

// newQueue initializes and returns a new queue.
func newQueue() *Queue {
	l := list.New()
	return &Queue{Elements: l}
}

// Size returns the total number of elements in the queue.
func (q *Queue) Size() int {
	return q.Elements.Len()
}

// Enqueue adds a element in the back of the queue.
func (q *Queue) Enqueue(v interface{}) {
	_ = q.Elements.PushBack(v)
}

// Dequeue removes and returns the element in the front of the queue.
func (q *Queue) Dequeue() interface{} {
	return q.Elements.Remove(q.Elements.Front())
}

func newTree() *Tree {
	return &Tree{}
}

func insertNode(node, newNode *Node) {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}

	}
}

func (t *Tree) Insert(value int) {
	newNode := &Node{
		value: value,
	}
	if t.root == nil {
		t.root = newNode
	} else {
		insertNode(t.root, newNode)
	}
	t.size++
}

func printTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printTree(node.left, level)
		fmt.Printf(format+"%d\n", node.value)
		printTree(node.right, level)
	}
}

func (t *Tree) Display() {
	printTree(t.root, 0)
}

func (t *Tree) BreadthFirstSearch() {
	currentNode := t.root
	bfsList := make([]int, 0)
	// nodeQueue keeps track of the tree level and used to access the children
	nodeQueue := newQueue()

	nodeQueue.Enqueue(currentNode)

	for nodeQueue.Size() > 0 {
		currentNode = nodeQueue.Dequeue().(*Node)
		bfsList = append(bfsList, currentNode.value)
		if currentNode.left != nil {
			nodeQueue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			nodeQueue.Enqueue(currentNode.right)
		}
	}
	fmt.Printf("%v\n", bfsList)
}

func (t *Tree) BreadthFirstSearchRecursive(nodeQueue *Queue, bfsList []int) []int {
	if nodeQueue.Size() == 0 {
		return bfsList
	}
	currentNode := nodeQueue.Dequeue().(*Node)
	bfsList = append(bfsList, currentNode.value)
	if currentNode.left != nil {
		nodeQueue.Enqueue(currentNode.left)
	}
	if currentNode.right != nil {
		nodeQueue.Enqueue(currentNode.right)
	}
	return t.BreadthFirstSearchRecursive(nodeQueue, bfsList)
}

func traverseInOrder(node *Node, dfsList []int) []int {
	if node.left != nil {
		dfsList = traverseInOrder(node.left, dfsList)
	}
	dfsList = append(dfsList, node.value)
	if node.right != nil {
		dfsList = traverseInOrder(node.right, dfsList)
	}
	return dfsList
}

func (t *Tree) DepthFirstSearchInOrder() {
	dfsList := make([]int, 0)
	fmt.Printf("%v\n", traverseInOrder(t.root, dfsList))
}

func traversePreOrder(node *Node, dfsList []int) []int {
	dfsList = append(dfsList, node.value)
	if node.left != nil {
		dfsList = traversePreOrder(node.left, dfsList)
	}
	if node.right != nil {
		dfsList = traversePreOrder(node.right, dfsList)
	}
	return dfsList
}

func (t *Tree) DepthFirstSearchPreOrder() {
	dfsList := make([]int, 0)
	fmt.Printf("%v\n", traversePreOrder(t.root, dfsList))
}

func traversePostOrder(node *Node, dfsList []int) []int {
	if node.left != nil {
		dfsList = traversePostOrder(node.left, dfsList)
	}
	if node.right != nil {
		dfsList = traversePostOrder(node.right, dfsList)
	}
	dfsList = append(dfsList, node.value)
	return dfsList
}

func (t *Tree) DepthFirstSearchPostOrder() {
	dfsList := make([]int, 0)
	fmt.Printf("%v\n", traversePostOrder(t.root, dfsList))
}

func main() {
	t := newTree()
	t.Insert(8)
	t.Insert(4)
	t.Insert(10)
	t.Insert(2)
	t.Insert(6)
	t.Insert(1)
	t.Insert(3)
	t.Insert(5)
	t.Insert(7)
	t.Insert(9)
	t.Display()
	t.BreadthFirstSearch()

	nodeQueue := newQueue()
	nodeQueue.Enqueue(t.root)
	bfsList := make([]int, 0)
	fmt.Printf("%v\n", t.BreadthFirstSearchRecursive(nodeQueue, bfsList))

	fmt.Println("Depth First Search InOrder:")
	t.DepthFirstSearchInOrder()

	fmt.Println("Depth First Search PreOrder:")
	t.DepthFirstSearchPreOrder()

	fmt.Println("Depth First Search PostOrder:")
	t.DepthFirstSearchPostOrder()
}
