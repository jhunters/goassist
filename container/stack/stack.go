// stack package provides stack(FILO) feature apis. note not safety in concurrent operation.
package stack

import "sync"

var (
	default_init_size = 16
)

// Stack struct
type Stack[E any] struct {
	data []E
	pos  int

	lock sync.Mutex

	empty E
}

func NewStack[E any]() *Stack[E] {
	return NewStackSize[E](default_init_size)
}

func NewStackSize[E any](initSize int) *Stack[E] {
	if initSize <= 0 {
		initSize = default_init_size
	}
	s := &Stack[E]{data: make([]E, initSize), pos: -1}
	return s
}

func (s *Stack[E]) resize() {
	l := len(s.data)
	newData := make([]E, l*2)
	copy(newData, s.data)
	s.data = newData
}

// Pop push the element into the stack
func (s *Stack[E]) Push(e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.Cap() <= 0 {
		s.resize()
	}
	s.pos++
	s.data[s.pos] = e
}

// Pop pop the element from the stack
func (s *Stack[E]) Pop() (e E) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pos >= 0 {
		e = s.data[s.pos]
		s.data[s.pos] = s.empty
		s.pos--
	}
	return
}

// Cap return capacity count
func (s *Stack[E]) Cap() int {
	return len(s.data) - s.pos - 1
}

// IsEmpty to return if stack has elements
func (s *Stack[E]) IsEmpty() bool {
	return s.pos < 0
}

// Size to get stack elements size
func (s *Stack[E]) Size() int {
	return s.pos + 1
}

// Copy copy to a new stack
func (s *Stack[E]) Copy() *Stack[E] {
	data := make([]E, len(s.data))
	copy(data, s.data)
	r := &Stack[E]{data: data, pos: s.pos}
	return r
}
