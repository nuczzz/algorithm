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

func (root *BSNode) SetLeft(left *BSNode) {
	root.left = left
	left.parent = root
}

func (root *BSNode) SetRight(right *BSNode) {
	root.right = right
	right.parent = root
}

func (root *BSNode)InOrder() {
	if root != nil {
		root.left.InOrder()
		fmt.Println(root.data)
		root.right.InOrder()
	}
}

func (root *BSNode) Search(data int) *BSNode {
	if root == nil || root.data == data {
		return root
	}
	if root.data < data {
		return root.right.Search(data)
	}
	return root.left.Search(data)
}

func (root *BSNode)Minimum() *BSNode {
	if root == nil {
		return nil
	}
	if root.left != nil {
		return root.left.Minimum()
	}
	return root
}

func (root *BSNode)Maximum() *BSNode {
	if root == nil {
		return nil
	}
	if root.right != nil {
		return root.right.Maximum()
	}
	return root
}

func (root *BSNode)Precursor() *BSNode {
	if root.left != nil {
		return root.left.Maximum()
	}
	child := root
	parent := root.parent
	for parent != nil && child == parent.left {
		child = parent
		parent = parent.parent
	}
	return parent
}

func (root *BSNode)Successor() *BSNode {
	if root.right != nil {
		return root.right.Minimum()
	}
	child := root
	parent := root.parent
	for parent != nil && child == parent.right {
		child = parent
		parent = parent.parent
	}
	return parent
}

//root can not be nil
func (root *BSNode) Insert(node *BSNode) {
	parent := root
	var aim *BSNode
	for parent != nil {
		aim = parent
		if node.data < parent.data {
			parent = parent.left
		} else {
			parent = parent.right
		}
	}
	node.parent = aim
	if node.data < aim.data {
		aim.left = node
	} else {
		aim.right = node
	}
}

func (root *BSNode) Delete(node *BSNode) {
	//todo
}
