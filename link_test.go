package algorithm

import "testing"

func TraversalLink(head *LinkNode, t *testing.T) {
	node := head
	for node != nil {
		t.Logf("%v", node.Data)
		node = node.Next
	}
}

func InitLink() *LinkNode {
	data := []int{2, 1, 4, 3, 5}

	var head, tail *LinkNode
	for i, d := range data {
		node := NewLinkNode(d)

		if i == 0 {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

func TestLink(t *testing.T) {
	head := InitLink()
	TraversalLink(head, t)
}
