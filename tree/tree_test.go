package tree

import (
	"testing"
)

func InitTestTree() *BinaryTree {
	root := NewBinaryTreeNode(1)
	root.Left = NewBinaryTreeNode(2)
	root.Right = NewBinaryTreeNode(3)

	root.Left.Left = NewBinaryTreeNode(4)
	root.Left.Right = NewBinaryTreeNode(5)

	root.Right.Left = NewBinaryTreeNode(6)
	root.Right.Right = NewBinaryTreeNode(7)

	return root
}

func TestTreeTraversal(t *testing.T) {
	root := InitTestTree()

	t.Logf(">>FirstOrderTraversal:")
	list := root.FirstOrderTraversal()
	for _, element := range list {
		t.Log(element)
	}

	t.Logf(">>SequentialTraversal:")
	list = root.SequentialTraversal()
	for _, element := range list {
		t.Log(element)
	}

	t.Logf(">>PostOrderTraversal:")
	list = root.PostOrderTraversal()
	for _, element := range list {
		t.Log(element)
	}
}
