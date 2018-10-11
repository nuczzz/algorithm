package queue

import (
	"errors"
)

type Queue interface {
	IsEmpty() bool
	Pop() (interface{}, error)
	Push(interface{}) error
	List() []interface{}
}

type Instance func(int) Queue

var (
	adapters = make(map[string]Instance)

	ErrQueueRegisterTwice = errors.New("register queue twice")
	ErrQueueNotRegister = errors.New("queue not register")
)

func Register(name string, instance Instance) error {
	if _, ok := adapters[name]; ok {
		return ErrQueueRegisterTwice
	}

	adapters[name] = instance
	return nil
}

func NewQueue(name string, length int) (queue Queue, err error) {
	adapter, ok := adapters[name]
	if !ok {
		err = ErrQueueNotRegister
		return
	}

	queue = adapter(length)
	return
}
