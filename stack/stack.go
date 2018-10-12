package stack

import "errors"

type Stack interface {
	// is stack empty
	IsEmpty() bool

	// pop element from stack
	Pop() (interface{}, error)

	// push element to stack
	Push(interface{}) error

	// stack elements list
	List() []interface{}

	// length of stack elements
	Len() int
}

type Instance func(int) Stack

var (
	adapters = make(map[string]Instance)

	ErrStackRegisterTwice = errors.New("register stack twice")
	ErrStackNotRegister   = errors.New("stack not register")
)

func Register(name string, instance Instance) error {
	if _, ok := adapters[name]; ok {
		return ErrStackRegisterTwice
	}

	adapters[name] = instance
	return nil
}

func NewStack(name string, length int) (stack Stack, err error) {
	adapter, ok := adapters[name]
	if !ok {
		err = ErrStackNotRegister
		return
	}

	stack = adapter(length)
	return
}
