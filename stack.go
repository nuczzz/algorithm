package algorithm

import (
	"errors"
)

type Stack struct {
	maxLength int
	elements  []interface{}
}

const defaultStackLength = 3

var (
	ErrStackUnderflow = errors.New("stack underflow")
	ErrStackOverflow  = errors.New("stack overflow")
)

func NewStack(length int) *Stack {
	return &Stack{
		maxLength: length,
		elements:  make([]interface{}, 0, length),
	}
}

func NewDefaultStack() *Stack {
	return NewStack(defaultStackLength)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}

	length := len(s.elements)
	ret := s.elements[length-1]
	s.elements = s.elements[:length-1]

	return ret, nil
}

func (s *Stack) Push(element interface{}) error {
	if len(s.elements) >= s.maxLength {
		return ErrStackOverflow
	}

	s.elements = append(s.elements, element)

	return nil
}

func (s *Stack) List() []interface{} {
	return s.elements
}
