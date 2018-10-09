package algorithm

import "testing"

const queueLength = 3

func TestNewDefaultQueue(t *testing.T) {
	// create a new queue
	queue := NewDefaultQueue()

	// underflow test
	_, err := queue.Pop()
	if err == nil {
		t.Fatal("should underflow")
	}
	t.Logf(">>underflow test success")

	// push elements into queue
	for i := 0; i < queueLength; i++ {
		err := queue.Push(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	// get queue elements list
	list := queue.List()
	for ele := range list {
		t.Logf("push element: %v\n", ele)
	}
	t.Logf(">>push element test success")

	// overflow test
	err = queue.Push(11)
	if err == nil {
		t.Fatal("should overflow")
	}
	t.Logf(">>overflow test success")

	// pop element
	for i := 0; i < queueLength; i++ {
		ele, err := queue.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("pop element: %v\n", ele)
	}
	t.Logf(">>pop element test success")

	t.Log("TestNewDefaultQueue success")
}
