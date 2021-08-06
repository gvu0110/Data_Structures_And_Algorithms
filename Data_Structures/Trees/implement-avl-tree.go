package main

import (
	"fmt"
)

/*
Left rotation:
- Set temp = grandparent's right child
- Point grandparent's right child to temp's left child
- Point temp's left child to grandparent
- Use temp instead of grandparent

Right rotation:
- Set temp = grandparent's left child
- Point grandparent's left child to temp's right child
- Point temp's right child to grandparent
- Use temp instead of grandparent

Right left rotation:
- Right rotate on the parent
- Left rotate on the grandparent

Left right rotation:
- Left rotate on the parent
- Right rotate on the grandparent
*/

type Node struct {
	parent *Node
	right  *Node
	left   *Node
	value  int
}

type Tree struct {
	root *Node
	size int
}

func newTree() *Tree {
	return &Tree{}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// leftRotate and rightRotate function input is the grandparent node and return the root of the balanced subtree
func leftRotate(node *Node) *Node {
	tempNode := node.right
	node.right = tempNode.left
	tempNode.left = node
	return tempNode
}

func rightRotate(node *Node) *Node {
	tempNode := node.left
	node.left = tempNode.right
	tempNode.right = node
	return tempNode
}

func rightLeftRotate(node *Node) *Node {
	tempNode := node.right
	node.right = rightRotate(tempNode)
	node.right.parent = node
	tempNode.parent = node.right
	return leftRotate(node)
}

func leftRightRotate(node *Node) *Node {
	tempNode := node.left
	node.left = leftRotate(tempNode)
	node.left.parent = node
	tempNode.parent = node.left
	return rightRotate(node)
}

func (t *Tree) CheckBalance(node *Node) {
	//fmt.Printf("Checking node: %v\n", node)
	currentParent := node.parent
	if (height(node.left)-height(node.right) > 1) || (height(node.left)-height(node.right) < -1) {
		t.Rebalance(node)
	}
	if currentParent != nil {
		t.CheckBalance(currentParent)
	}
}

func (t *Tree) Rebalance(node *Node) {
	var newNode *Node
	if height(node.left)-height(node.right) > 1 {
		if height(node.left.left) > height(node.left.right) {
			newNode = rightRotate(node)
		} else {
			newNode = leftRightRotate(node)
		}
	} else if height(node.left)-height(node.right) < -1 {
		if height(node.right.right) > height(node.right.left) {
			newNode = leftRotate(node)
		} else {
			newNode = rightLeftRotate(node)
		}
	}

	if node.parent == nil {
		t.root = newNode
		//fmt.Printf("Address: %p, Parent: %p, Left: %p, Right: %p, Value: %d\n", newNode, newNode.parent, newNode.left, newNode.right, newNode.value)
	} else {
		if node.parent.left == node {
			node.parent.left = newNode
		}
		if node.parent.right == node {
			node.parent.right = newNode
		}
		newNode.parent = node.parent
		node.parent = newNode
		//fmt.Printf("Address: %p, Parent: %p, Left: %p, Right: %p, Value: %d\n", newNode, newNode.parent, newNode.left, newNode.right, newNode.value)
	}
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	heightLeft := 1
	heightRight := 1
	if node.left != nil {
		heightLeft = height(node.left) + 1
	}
	if node.right != nil {
		heightRight = height(node.right) + 1
	}
	return max(heightLeft, heightRight)
}

func insertNode(node, newNode *Node) *Node {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
			newNode.parent = node
			return newNode
		}
		return insertNode(node.left, newNode)
	} else {
		if node.right == nil {
			node.right = newNode
			newNode.parent = node
			return newNode
		}
		return insertNode(node.right, newNode)
	}
}

func (t *Tree) Insert(value int) {
	newNode := &Node{
		value: value,
	}
	if t.root == nil {
		t.root = newNode
		t.size++
	} else {
		newAddedNode := insertNode(t.root, newNode)
		t.size++
		t.CheckBalance(newAddedNode)
	}

}

func printTree(node *Node, level int) {
	if node != nil {
		//fmt.Printf("Address: %p, Parent: %p, Left: %p, Right: %p, Value: %d\n", node, node.parent, node.left, node.right, node.value)
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
	fmt.Println("--------------------------------")
}

func main() {
	t := newTree()
	t.Insert(8)
	t.Insert(9)
	t.Insert(6)
	t.Insert(4)
	t.Display()
	t.Insert(3)
	t.Display()
	t.Insert(7)
	t.Display()
	t.Insert(2)
	t.Display()
	t.Insert(1)
	t.Display()
	t.Insert(0)
	t.Display()
}
