package algorithm

import "errors"

type Queue struct {
	maxLength int
	elements  []interface{}
}

const defaultQueueLength = 3

var (
	ErrQueueUnderflow = errors.New("queue underflow")
	ErrQueueOverflow  = errors.New("queue overflow")
)

func NewQueue(length int) *Queue {
	return &Queue{
		maxLength: length,
		elements:  make([]interface{}, 0, length),
	}
}

func NewDefaultQueue() *Queue {
	return NewQueue(defaultQueueLength)
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrQueueUnderflow
	}

	ret := q.elements[0]
	q.elements = q.elements[1:]

	return ret, nil
}

func (q *Queue) Push(element interface{}) error {
	if len(q.elements) >= q.maxLength {
		return ErrQueueOverflow
	}

	q.elements = append(q.elements, element)

	return nil
}

func (q *Queue) List() []interface{} {
	return q.elements
}
