package link

import (
	"testing"
)

type TestingLog struct {
	t *testing.T
}

func NewTestingLog(t *testing.T) LinkLog {
	return &TestingLog{
		t: t,
	}
}

func (tl *TestingLog) Printf(format string, data...interface{}){
	tl.t.Logf(format, data)
}

func TestLink(t *testing.T) {
	log := NewTestingLog(t)

	head := InitTestLink(3)
	TraversalLink(head, log)
}

func TestReverseLink(t *testing.T) {
	log := NewTestingLog(t)

	head := InitTestLink(0)
	t.Log("zero node link traversal:")
	TraversalLink(head, log)
	t.Log("zero node link reverse:")
	head = ReverseLink(head)
	TraversalLink(head, log)

	head = InitTestLink(1)
	t.Log("one node link traversal:")
	TraversalLink(head, log)
	t.Log("one node link reverse:")
	head = ReverseLink(head)
	TraversalLink(head, log)

	head = InitTestLink(2)
	t.Log("two node link traversal:")
	TraversalLink(head, log)
	t.Log("two node link reverse:")
	head = ReverseLink(head)
	TraversalLink(head, log)

	head = InitTestLink(3)
	t.Log("three node link traversal:")
	TraversalLink(head, log)
	t.Log("three node link reverse:")
	head = ReverseLink(head)
	TraversalLink(head, log)

	head = InitTestLink(10)
	t.Log("ten node link traversal:")
	TraversalLink(head, log)
	t.Log("ten node link reverse:")
	head = ReverseLink(head)
	TraversalLink(head, log)
}
