package stack

import "errors"

type Stack interface {
	IsEmpty() bool
	Pop() (interface{}, error)
	Push(element interface{}) error
	List() []interface{}
}

type Instance func(int) Stack

var (
	adapters = make(map[string]Instance)

	ErrStackRegisterTwice = errors.New("register stack twice")
	ErrStackNotRegister   = errors.New("stack not register")
)

func Register(name string, instance Instance) error {
	instance, ok := adapters[name]
	if ok {
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
