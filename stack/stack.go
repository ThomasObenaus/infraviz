package stack

import "sync"

type stack struct {
	lock  sync.Mutex
	stack []interface{}
}

func (s *stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stack = append(s.stack, v)
}

func (s *stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.stack)
}

func (s *stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	numElements := len(s.stack)
	if numElements == 0 {
		return nil
	}

	element := s.stack[numElements-1]
	s.stack = s.stack[:numElements-1]
	return element
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]interface{}, 0)}
}
