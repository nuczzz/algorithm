package queue

import (
	"errors"
)

type ArrayQueue struct {
	maxLength int
	elements  []interface{}
}

const defaultArrayQueueLength = 3

var (
	ErrQueueUnderflow = errors.New("queue underflow")
	ErrQueueOverflow  = errors.New("queue overflow")
)

func NewArrayQueue(length int) Queue {
	if length <= 0 {
		length = defaultArrayQueueLength
	}

	return &ArrayQueue{
		maxLength: length,
		elements:  make([]interface{}, 0, length),
	}
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *ArrayQueue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueUnderflow
	}

	ret := q.elements[0]
	q.elements = q.elements[1:]

	return ret, nil
}

func (q *ArrayQueue) Push(element interface{}) error {
	if len(q.elements) >= q.maxLength {
		return ErrQueueOverflow
	}

	q.elements = append(q.elements, element)

	return nil
}

func (q *ArrayQueue) List() []interface{} {
	return q.elements
}

func init() {
	Register("array_queue", NewArrayQueue)
}
