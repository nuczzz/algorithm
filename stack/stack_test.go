package stack

import "testing"

const stackLength = 3

func TestArrayStack(t *testing.T) {
	// create a new stack
	stack := NewArrayStack(0)

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
		t.Logf("push element: %v\n", i)
	}
	// get stack elements list
	list := stack.List()
	t.Logf(">>stack list: (bottom)%v", list)

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
