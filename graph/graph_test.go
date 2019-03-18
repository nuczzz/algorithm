package graph

import "testing"

// 0-not use, so just 1~5
const NodeNum = 6

func initGraph() *Graph {
	g := &Graph{list: make([]int, 0, NodeNum)}
	for i := 0; i < NodeNum; i++ {
		g.list[i] = i
	}
	g.neighbors = initNeighborList()
	return g
}

func initNeighborList() *NeighborList {
	n := &NeighborList{list: make([]*Node, 0, NodeNum)}
	n.list[1] = NewNode(2).SetNext(5)
	n.list[2] = NewNode(1).SetNext(5).SetNext(3).SetNext(4)
	n.list[3] = NewNode(2).SetNext(4)
	n.list[4] = NewNode(2).SetNext(5).SetNext(3)
	n.list[5] = NewNode(4).SetNext(1).SetNext(2)
	return n
}

/*
	1 -- 2
    |   /| \
    |  / |  3
    | /  | /
    5 -- 4
*/
func TestBFS(t *testing.T) {
	graph := initGraph()

}
