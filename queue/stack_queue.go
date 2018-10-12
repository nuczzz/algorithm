package queue

import "algorithm/stack"

// queue which is made up of stack
type StackQueue struct {
	stack1   stack.Stack
	stack2 stack.Stack
}

const defaultStackQueueLength = 3

func NewStackQueue(length int) Queue {
	if length <= 0 {
		length = defaultStackQueueLength
	}

	stack1, _ := stack.NewStack("array_stack", length)
	stack2, _ := stack.NewStack("array_stack", length)

	return &StackQueue{
		stack1: stack1,
		stack2: stack2,
	}
}

func (q *StackQueue) IsEmpty() bool {
	return q.stack1.IsEmpty() && q.stack2.IsEmpty()
}

func (q *StackQueue) Pop() (interface{}, error) {
	if q.stack1.IsEmpty() && q.stack2.IsEmpty() {
		return nil, stack.ErrStackUnderflow
	}

	if !q.stack2.IsEmpty() {
		return q.stack2.Pop()
	}

	for !q.stack1.IsEmpty() {
		data, _ := q.stack1.Pop()
		q.stack2.Push(data)
	}

	return q.stack2.Pop()
}

func (q *StackQueue) Push(data interface{}) error {
	return q.stack1.Push(data)
}

func (q *StackQueue) List() []interface{} {
	return append(q.stack1.List(), q.stack2.List()...)
}

func init() {
	Register("stack_queue", NewArrayQueue)
}
