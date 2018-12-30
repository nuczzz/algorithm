package bstree

import "fmt"

type BSNode struct {
	data int

	left   *BSNode
	right  *BSNode
	parent *BSNode
}

func NewBSNode(data int) *BSNode {
	return &BSNode{data: data}
}

func (bs *BSNode) SetLeft(left *BSNode) {
	bs.left = left
	left.parent = bs
}

func (bs *BSNode) SetRight(right *BSNode) {
	bs.right = right
	right.parent = bs
}

func InOrder(root *BSNode) {
	if root != nil {
		InOrder(root.left)
		fmt.Println(root.data)
		InOrder(root.right)
	}
}

func Search(root *BSNode, data int) *BSNode {
	if root == nil || root.data == data {
		return root
	}
	if root.data < data {
		return Search(root.right, data)
	}
	return Search(root.left, data)
}
