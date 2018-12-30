package bstree

import (
	"fmt"
	"testing"
)

/*
           15
         /    \
        6      18
      /   \    / \
     3     7  17  20
    / \     \
   2   4     13
            /
           9
*/

func initTestBSTree() *BSNode {
	n2 := NewBSNode(2)
	n3 := NewBSNode(3)
	n4 := NewBSNode(4)
	n6 := NewBSNode(6)
	n7 := NewBSNode(7)
	n13 := NewBSNode(13)
	n9 := NewBSNode(9)
	n15 := NewBSNode(15)
	n17 := NewBSNode(17)
	n18 := NewBSNode(18)
	n20 := NewBSNode(20)

	n15.SetLeft(n6)
	n15.SetRight(n18)
	n6.SetLeft(n3)
	n6.SetRight(n7)
	n3.SetLeft(n2)
	n3.SetRight(n4)
	n7.SetRight(n13)
	n13.SetLeft(n9)
	n18.SetLeft(n17)
	n18.SetRight(n20)

	return n15
}

func TestInOrder(t *testing.T) {
	root := initTestBSTree()
	InOrder(root)
}

func TestSearch(t *testing.T) {
	root := initTestBSTree()
	fmt.Println(Search(root, 9))
}

func TestMinimum(t *testing.T) {
	root := initTestBSTree()
	fmt.Println(Minimum(root))
}

func TestMaximum(t *testing.T) {
	root := initTestBSTree()
	fmt.Println(Maximum(root))
}

func TestPrecursor(t *testing.T) {
	root := initTestBSTree()
	fmt.Println(Precursor(root.left.left.left))//2
	fmt.Println(Precursor(root.left.left.right))//4
}

func TestSuccessor(t *testing.T) {
	root := initTestBSTree()
	fmt.Println(Successor(root.left.right.right)) //13
}
