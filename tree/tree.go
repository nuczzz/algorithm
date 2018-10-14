package tree

type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

func NewBinaryTreeNode(data interface{}) *BinaryTree {
	return &BinaryTree{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (root *BinaryTree) FirstOrderTraversal() []interface{}{
	if root == nil {
		return nil
	}

	ret := make([]interface{}, 0)
	ret1 := []interface{}{root.Data}
	var ret2, ret3 []interface{}
	if root.Left != nil {
		ret2 = root.Left.FirstOrderTraversal()
	}
	if root.Right != nil {
		ret3 = root.Right.FirstOrderTraversal()
	}

	ret = append(ret1, ret2...)
	ret = append(ret, ret3...)
	return ret
}

func (root *BinaryTree) SequentialTraversal() []interface{} {
	if root == nil {
		return nil
	}

	ret := make([]interface{}, 0)
	ret1 := []interface{}{root.Data}
	var ret2, ret3 []interface{}
	if root.Left != nil {
		ret2 = root.Left.SequentialTraversal()
	}
	if root.Right != nil {
		ret3 = root.Right.SequentialTraversal()
	}

	ret = append(ret2, ret1...)
	ret = append(ret, ret3...)
	return ret
}

func (root *BinaryTree) PostOrderTraversal() []interface{} {
	if root == nil {
		return nil
	}

	ret := make([]interface{}, 0)
	ret1 := root.Data
	var ret2, ret3 []interface{}
	if root.Left != nil {
		ret2 = root.Left.PostOrderTraversal()
	}
	if root.Right != nil {
		ret3 = root.Right.PostOrderTraversal()
	}

	ret = append(ret2, ret3...)
	ret = append(ret, ret1)
	return ret
}
