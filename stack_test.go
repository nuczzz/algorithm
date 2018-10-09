package algorithm

import "testing"

const stackLength = 3

func TestNewDefaultStack(t *testing.T) {
	stack := NewDefaultStack()
	_, err := stack.Pop()
	if err == nil {
		t.Fatal("should underflow")
	}
	t.Log(err)

	for i := 0; i < stackLength; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}

	list := stack.List()
	for ele := range list {
		t.Log(ele)
	}

	err = stack.Push(11)
	if err == nil {
		t.Fatal("should overflow")
	}
	t.Log(err)

	t.Log("TestNewDefaultStack success")
}
