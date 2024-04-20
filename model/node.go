package model

type Node struct {
	Id       string
	Vertices []*Vertex
}

func (n *Node) AddVertex(v *Vertex) {
	n.Vertices = append(n.Vertices, v)
}
