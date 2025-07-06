package match

import (
	"errors"
	"unsafe"
)

type stateValue[T comparable] struct {
	states [256]uint32
	value  T
}

type StateMachine[T comparable] struct {
	states []stateValue[T]
}

func New[T comparable]() *StateMachine[T] {
	return &StateMachine[T]{
		states: make([]stateValue[T], 2),
	}
}

func (s *StateMachine[T]) compile(state uint32, data []byte, value T) error {
	if len(data) == 0 {
		var d T

		if s.states[state].value != d {
			return ErrAmbiguous
		}

		s.states[state].value = value

		return nil
	}

	c := data[0]
	data = data[1:]

	next := s.states[state].states[c]
	if next == 0 {
		next = uint32(len(s.states))
		s.states = append(s.states, stateValue[T]{})
		s.states[state].states[c] = next
	}

	return s.compile(next, data, value)
}

func (s *StateMachine[T]) AddString(str string, value T) error {
	return s.compile(1, strToBytes(str), value)
}

func (s *StateMachine[T]) Match(str string) T {
	state := uint32(1)

	for _, c := range strToBytes(str) {
		state = s.states[state].states[c]
	}

	return s.states[state].value
}

func strToBytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

var ErrAmbiguous = errors.New("ambiguous states")
