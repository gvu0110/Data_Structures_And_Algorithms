package main

import "fmt"

type Node struct {
	key int
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

func newGraph() *Graph {
	return &Graph{
		nodes: make([]*Node, 0),
		edges: make(map[Node][]*Node),
	}
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(node1, node2 *Node) {
	g.edges[*node1] = append(g.edges[*node1], node2)
	g.edges[*node2] = append(g.edges[*node2], node1)
}

func (g *Graph) ShowConnections() {
	for _, node := range g.nodes {
		fmt.Printf("%d --> ", node.key)
		if len(g.edges) != 0 {
			for _, edge := range g.edges[*node] {
				fmt.Printf("%d ", edge.key)
			}
		}
		fmt.Println()
	}
}

func main() {
	g := newGraph()
	node0 := Node{key: 0}
	node1 := Node{key: 1}
	node2 := Node{key: 2}
	node3 := Node{key: 3}
	node4 := Node{key: 4}
	node5 := Node{key: 5}
	node6 := Node{key: 6}
	g.AddNode(&node0)
	g.AddNode(&node1)
	g.AddNode(&node2)
	g.AddNode(&node3)
	g.AddNode(&node4)
	g.AddNode(&node5)
	g.AddNode(&node6)
	g.ShowConnections()
	g.AddEdge(&node0, &node1)
	g.AddEdge(&node0, &node2)
	g.AddEdge(&node1, &node2)
	g.AddEdge(&node1, &node3)
	g.AddEdge(&node2, &node4)
	g.AddEdge(&node3, &node4)
	g.AddEdge(&node4, &node5)
	g.AddEdge(&node5, &node6)
	g.ShowConnections()
}
