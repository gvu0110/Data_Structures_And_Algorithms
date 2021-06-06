package main

import (
	"fmt"
)

type Node struct {
	right *Node
	left  *Node
	value int
}

type Tree struct {
	root *Node
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
}

func searchNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value == node.value {
		return node
	} else if value < node.value {
		return searchNode(node.left, value)
	} else {
		return searchNode(node.right, value)
	}
}

func (t *Tree) Search(value int) *Node {
	return searchNode(t.root, value)
}

func (t *Tree) Delete(value int) {
	currentNode := t.root
	var parentNode *Node
	for currentNode != nil && currentNode.value != value {
		parentNode = currentNode
		if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// Case 1: Node to be deleted has no children (it's a leaf node): Remove the node from the tree
	// Case 2: Node to be deleted has only one child: Replace the node with it's child
	// Case 3: Node to be deleted has two children: Replace the inorder successor (node with the minimum value in the right subtree)
	// - Find inorder successor of the node
	// - Copy contents of the inorder successor to the node and delete the inorder successor
	// - Note that inorder predecessor can also be used.
	if currentNode.left == nil && currentNode.right == nil { // Case 1
		if parentNode == nil {
			t.root = nil
		} else {
			if parentNode.left == currentNode {
				parentNode.left = nil
			} else {
				parentNode.right = nil
			}
		}
	} else if currentNode.left == nil || currentNode.right == nil { // Case 2
		var childNode *Node
		if currentNode.left != nil {
			childNode = currentNode.left
		} else {
			childNode = currentNode.right
		}

		if parentNode == nil {
			t.root = childNode
		} else {
			if parentNode.left == currentNode {
				parentNode.left = childNode
			} else {
				parentNode.right = childNode
			}
		}
	} else { // Case 3:
		inOrderSuccessor, _ := t.FindInOrderSuccessorAndPredecessor(currentNode)
		t.Delete(inOrderSuccessor.value)
		currentNode.value = inOrderSuccessor.value

		// _, inOrderPredecessor := t.FindInOrderSuccessorAndPredecessor(currentNode)
		// t.Delete(inOrderPredecessor.value)
		// currentNode.value = inOrderPredecessor.value
	}
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

func inOrderSuccessorAndPredecessor(root, node *Node) (successor, predecessor *Node) {
	/*
			Input: root node, node
			Output: predecessor node, successor node

			1. If root is null
			return

			2. If node is found then
		    a. If its left subtree is not null
		        Then predecessor will be the right most
		        child of left subtree or left child itself.
		    b. If its right subtree is not null
		        The successor will be the left most child
		        of right subtree or right child itself.
		    return

			3. If node's value is smaller than root node's value
		        set the successor as root
		        search recursively into left subtree
		    else
		        set the predecessor as root
		        search recursively into right subtree
	*/
	if root != nil {
		if root.value == node.value {
			// go to the left most element in the right subtree, it will be the successor
			if root.right != nil {
				temp := root.right
				for temp.left != nil {
					temp = temp.left
				}
				successor = temp
			}
			// go to the right most node in the left subtree, it will be the predecessor
			if root.left != nil {
				temp := root.left
				for temp.right != nil {
					temp = temp.right
				}
				predecessor = temp
			}
			return
		} else if root.value > node.value {
			successor = root
			return inOrderSuccessorAndPredecessor(root.left, node)
		} else {
			predecessor = root
			return inOrderSuccessorAndPredecessor(root.right, node)
		}
	}
	return
}

func (t *Tree) FindInOrderSuccessorAndPredecessor(node *Node) (*Node, *Node) {
	return inOrderSuccessorAndPredecessor(t.root, node)
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.value)
		inOrder(node.right)
	}
}

func (t *Tree) TraverseInOrder() {
	inOrder(t.root)
	fmt.Println()
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
	fmt.Printf("%#v\n", t.Search(7))
	t.TraverseInOrder()
	inOrderSuccessor, inOrderPredecessor := t.FindInOrderSuccessorAndPredecessor(t.root.left)
	fmt.Printf("Successor: %#v, Predecessor: %#v\n", inOrderSuccessor, inOrderPredecessor)
	t.Delete(4)
	t.Display()
}
