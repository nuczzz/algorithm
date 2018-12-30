package bstree

import (
	"fmt"
)

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

func Minimum(root *BSNode) *BSNode {
	if root == nil {
		return nil
	}
	if root.left != nil {
		return Minimum(root.left)
	}
	return root
}

func Maximum(root *BSNode) *BSNode {
	if root == nil {
		return nil
	}
	if root.right != nil {
		return Maximum(root.right)
	}
	return root
}

func Precursor(node *BSNode) *BSNode {
	if node.left != nil {
		return Maximum(node.left)
	}
	child := node
	parent := node.parent
	for parent != nil && child == parent.left {
		child = parent
		parent = parent.parent
	}
	return parent
}

func Successor(node *BSNode) *BSNode {
	if node.right != nil {
		return Minimum(node.right)
	}
	child := node
	parent := node.parent
	for parent != nil && child == parent.right {
		child = parent
		parent = parent.parent
	}
	return parent
}

func (bs *BSNode) Insert(node *BSNode) {
	//todo
}

func (bs *BSNode) Delete(node *BSNode) {
	//todo
}
