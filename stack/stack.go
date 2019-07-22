package stack

import "sync"

type Stacker interface {
	Push(interface{})
	Pop() interface{}
}

type item struct {
	value interface{}
	next  *item
}

type Stack struct {
	mux  sync.Mutex
	top  *item
	size int
}

func NewStack() Stacker { return &Stack{} }

func (s *Stack) Push(value interface{}) {
	s.mux.Lock()
	s.top = &item{
		value: value,
		next:  s.top,
	}
	s.size++
	s.mux.Unlock()
}

func (s *Stack) Pop() interface{} {
	s.mux.Lock()
	if s.size <= 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	s.mux.Unlock()
	return value
}
