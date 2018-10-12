package queue

import (
	"errors"
)

type Queue interface {
	// is queue empty
	IsEmpty() bool

	// pop element from queue
	Pop() (interface{}, error)

	// push data to queue
	Push(interface{}) error

	// queue elements list
	List() []interface{}

	// length of queue elements
	Len() int
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
