package queue

import (
	"testing"
)

const queueLength = 3

func TestNewArrayQueue(t *testing.T) {
	// create a new queue
	queue := NewArrayQueue(queueLength)

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
		t.Logf("push element: %v\n", i)
	}

	// get queue elements list
	list := queue.List()
	t.Logf(">>queue list: (bottom)%v", list)

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

func TestNewQueue(t *testing.T) {
	queue, err := NewQueue("array_queue", 3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(queue)
}
