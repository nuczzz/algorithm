package link

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func NewLinkNode(data interface{}) *LinkNode {
	return &LinkNode{
		Data: data,
		Next: nil,
	}
}
