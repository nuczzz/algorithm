package algorithm

import "testing"

const stackLength = 3

func TestNewDefaultStack(t *testing.T) {
	// create a new stack
	stack := NewDefaultStack()

	// underflow test
	_, err := stack.Pop()
	if err == nil {
		t.Fatal("should underflow")
	}
	t.Logf(">>underflow test success")

	// push elements into stack
	for i := 0; i < stackLength; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	// get stack elements list
	list := stack.List()
	for ele := range list {
		t.Logf("push element: %v\n", ele)
	}
	t.Logf(">>push element test success")

	// overflow test
	err = stack.Push(11)
	if err == nil {
		t.Fatal("should overflow")
	}
	t.Logf(">>overflow test success")

	// pop element
	for i := 0; i < stackLength; i++ {
		ele, err := stack.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("pop element: %v\n", ele)
	}
	t.Logf(">>pop element test success")

	t.Log("TestNewDefaultStack success")
}
