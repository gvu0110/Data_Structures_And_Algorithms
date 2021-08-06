package main

import (
	"fmt"
)

/*
- Black aunt: rotate and end up with:
	black
	/	\
  red   red

- Red aunt: color flip and end up with:
	red
   /   \
black black
*/
type Node struct {
	parent *Node
	right *Node
	left  *Node
	isBlack bool
	isLeftChild bool
	value int
}

type Tree struct {
	root *Node
	size int
}

func newTree() *Tree {
	return &Tree{}
}

func correctTree(node *Node) {
	// Aunt is grandparent.right
	if node.parent.isLeftChild {
		// Black aunt, rotate
		if (node.parent.parent.right == nil) || node.parent.parent.right.isBlack {
			return rotate(node)
		}
		// Red aunt, color flip
		if !node.parent.parent.right.isBlack {
			node.parent.isBlack = false
			return
		}
	} else { // Aunt is grandparent.left

	}
}

func (t *tree) CheckColor(node *Node) {
	if node.parent == nil {
		return
	}

	if !node.isBlack && !node.parent.isBlack {
		correctTree(node)
	}
	t.CheckColor(node.parent)
}

func insertNode(node, newNode *Node) *Node {
	if newNode.value < node.value {
		if node.left == nil {
			node.left = newNode
			newNode.parent = node
			newNode.isLeftChild = true
			return newNode
		}
		return insertNode(node.left, newNode)
	} else {
		if node.right == nil {
			node.right = newNode
			newNode.parent = node
			newNode.isLeftChild = false
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
		t.root.isBlack = true
	} else {
		newAddNode := insertNode(t.root, newNode)
		t.CheckColor(newAddedNode)
	}
	t.size++
}