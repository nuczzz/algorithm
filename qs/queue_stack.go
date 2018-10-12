package qs

import (
	"algorithm/queue"
	"algorithm/stack"
)

type QueueStack struct {
	queue1 queue.Queue
	queue2 queue.Queue
}

const defaultQueueStackLength = 3

func NewQueueStack(length int) stack.Stack {
	if length <= 0 {
		length = defaultQueueStackLength
	}

	queue1, _ := queue.NewQueue("array_queue", length)
	queue2, _ := queue.NewQueue("array_queue", length)

	return &QueueStack{
		queue1: queue1,
		queue2: queue2,
	}
}


func (q *QueueStack) IsEmpty() bool {
	return q.queue1.IsEmpty() && q.queue2.IsEmpty()
}

func (q *QueueStack) Pop() (interface{}, error) {
	var storeQ, popQ queue.Queue
	if !q.queue1.IsEmpty() {
		popQ = q.queue1
		storeQ = q.queue2
	} else {
		popQ = q.queue2
		storeQ = q.queue1
	}

	length := popQ.Len()
	for i := 0; i < length-1; i++ {
		data, _ := popQ.Pop()
		storeQ.Push(data)
	}

	return popQ.Pop()
}

func (q *QueueStack) Push(data interface{}) error {
	if q.queue1.IsEmpty() {
		return q.queue2.Push(data)
	}

	return q.queue1.Push(data)
}

func (q *QueueStack) List() []interface{} {
	if q.queue1.IsEmpty() {
		return q.queue2.List()
	}

	return q.queue1.List()
}

func (q *QueueStack) Len() int {
	if q.queue1.IsEmpty() {
		return q.queue2.Len()
	}

	return q.queue1.Len()
}

func init() {
	stack.Register("queue_stack", NewQueueStack)
}
