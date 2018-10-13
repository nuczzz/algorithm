package link

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LinkLog interface {
	Printf(string, ...interface{})
}

func NewLinkNode(data interface{}) *LinkNode {
	return &LinkNode{
		Data: data,
		Next: nil,
	}
}

func TraversalLink(head *LinkNode, log LinkLog) {
	node := head
	for node != nil {
		log.Printf("%v\n", node.Data)
		node = node.Next
	}
}

func InitTestLink(num int) *LinkNode {
	var head, tail *LinkNode
	for i := 0; i < num; i++ {
		node := NewLinkNode(i)

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

func ReverseLink(head *LinkNode) *LinkNode {
	// null or one link node
	if head == nil || head.Next == nil {
		return head
	}

	// two or more link node
	var n1 *LinkNode
	n2 := head
	n3 := head.Next
	for n3 != nil{
		n2.Next = n1

		n1 = n2
		n2 = n3
		n3 = n3.Next
	}

	n2.Next = n1
	return n2
}
