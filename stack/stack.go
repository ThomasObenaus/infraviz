package stack

import "sync"

type Stack struct {
	lock  sync.Mutex
	stack []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stack = append(s.stack, v)
}

func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.stack)
}

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	numElements := len(s.stack)
	if numElements == 0 {
		return nil
	}

	return s.stack[numElements-1]
}

func (s *Stack) Pop() interface{} {
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

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]interface{}, 0)}
}
