package stack

import (
	"errors"
)

type ArrayStack struct {
	maxLength int
	elements  []interface{}
}

const defaultArrayStackLength = 3

var (
	ErrStackUnderflow = errors.New("stack underflow")
	ErrStackOverflow  = errors.New("stack overflow")
)

func NewArrayStack(length int) Stack {
	if length <= 0 {
		length = defaultArrayStackLength
	}

	return &ArrayStack{
		maxLength: length,
		elements:  make([]interface{}, 0, length),
	}
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *ArrayStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}

	length := len(s.elements)
	ret := s.elements[length-1]
	s.elements = s.elements[:length-1]

	return ret, nil
}

func (s *ArrayStack) Push(element interface{}) error {
	if len(s.elements) >= s.maxLength {
		return ErrStackOverflow
	}

	s.elements = append(s.elements, element)

	return nil
}

func (s *ArrayStack) List() []interface{} {
	return s.elements
}

func (s *ArrayStack) Len() int {
	return len(s.elements)
}

func init() {
	Register("array_stack", NewArrayStack)
}
