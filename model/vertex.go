package model

type Vertex struct {
	Node1  *Node
	Node2  *Node
	Weight int
}

func NewVertex(node1 *Node, node2 *Node, weight int) *Vertex {
	return &Vertex{Node1: node1, Node2: node2, Weight: weight}
}
