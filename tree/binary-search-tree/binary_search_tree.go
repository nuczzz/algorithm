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

type BSTree struct {
	root *BSNode
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

func (root *BSNode)inOrder() {
	if root != nil {
		root.left.inOrder()
		fmt.Println(root.data)
		root.right.inOrder()
	}
}

func (tree *BSTree) InOrder() {
	tree.root.inOrder()
}

func (root *BSNode) search(data int) *BSNode {
	if root == nil || root.data == data {
		return root
	}
	if root.data < data {
		return root.right.search(data)
	}
	return root.left.search(data)
}

func (tree *BSTree) Search(data int) *BSNode {
	return tree.root.search(data)
}

func (root *BSNode)minimum() *BSNode {
	if root == nil {
		return nil
	}
	if root.left != nil {
		return root.left.minimum()
	}
	return root
}

func (tree *BSTree) Minimum() *BSNode {
	return tree.root.minimum()
}

func (root *BSNode)maximum() *BSNode {
	if root == nil {
		return nil
	}
	if root.right != nil {
		return root.right.maximum()
	}
	return root
}

func (tree *BSTree) Maximum() *BSNode {
	return tree.root.maximum()
}

func (root *BSNode)Precursor() *BSNode {
	if root.left != nil {
		return root.left.maximum()
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
		return root.right.minimum()
	}
	child := root
	parent := root.parent
	for parent != nil && child == parent.right {
		child = parent
		parent = parent.parent
	}
	return parent
}

func (tree *BSTree) Insert(node *BSNode) {
	child := tree.root
	var parent *BSNode
	for child != nil {
		parent = child
		if node.data < child.data {
			child = child.left
		} else {
			child = child.right
		}
	}
	node.parent = parent

	if parent == nil {
		tree.root = node
	} else if node.data < parent.data {
		parent.left = node
	} else {
		parent.right = node
	}
}

func (tree *BSTree) Delete(node *BSNode) {
	//todo
}
