package graph

import "algorithm/queue"

const (
	ColorWhite = iota
	ColorGray
	ColorBlack
)

type Node struct {
	id       int
	distance int
	color    int
	next     *Node
}

func NewNode(id int) *Node {
	return &Node{id: id}
}

func (n *Node) SetNext(id int) *Node {
	n.next = NewNode(id)
	return n.next
}

type NeighborList struct {
	list []*Node
}

type NeighborMatrix struct {
	distance [][]int
}

type Graph struct {
	list      []int
	neighbors *NeighborList
}

// breadth first traversal
func BFTraversal(graph *Graph, start int) []*Node {
	list := make([]*Node, 0, len(graph.list))
	startNode := graph.neighbors.list[start]
	startNode.color = ColorGray

	q := queue.NewArrayQueue(len(graph.list))

}

// breadth first search
func BFS() {

}
