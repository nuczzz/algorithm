package qs

import (
	"testing"
	"algorithm/queue"
	"algorithm/stack"
)

func TestStackQueue(t *testing.T) {
	queue, _ := queue.NewQueue("stack_queue",3)

	t.Log(">>push data, and stack_queue element list:")
	for i := 0; i < 3; i++ {
		queue.Push(i)
	}

	list := queue.List()
	for _, ele := range list {
		t.Log(ele)
	}

	// pop test
	t.Log(">>pop test:")
	for i := 0; i < 3; i++ {
		ele, _ := queue.Pop()
		t.Log(ele)
	}
}

func TestQueueStack(t *testing.T) {
	stack, _ := stack.NewStack("queue_stack",3)

	t.Log(">>push data, and queue_stack element list:")
	for i := 0; i < 3; i++ {
		stack.Push(i)
	}

	list := stack.List()
	for _, ele := range list {
		t.Log(ele)
	}

	// pop test
	t.Log(">>pop test:")
	for i := 0; i < 3; i++ {
		ele, _ := stack.Pop()
		t.Log(ele)
	}
}
