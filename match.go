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

func (s *StateMachine[T]) AddString(str string, value T) error {
	o, err := parse[T](str, simpleStart)
	if err != nil {
		return err
	}

	_, err = o.compile(s, 1, make(visitedSet), value)

	return err
}

func (s *StateMachine[T]) Match(str string) T {
	return s.states[s.matchState(str, 1)].value
}

func (s *StateMachine[T]) matchState(str string, state uint32) uint32 {
	for _, c := range strToBytes(str) {
		state = s.states[state].states[c]
	}

	return state
}

func strToBytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

type State[T comparable] struct {
	sm    *StateMachine[T]
	state uint32
}

func (s *StateMachine[T]) MatchState(str string) State[T] {
	return State[T]{
		sm:    s,
		state: s.matchState(str, 1),
	}
}

func (s *State[T]) Match(str string) T {
	return s.sm.states[s.sm.matchState(str, s.state)].value
}

func (s *State[T]) MatchState(str string) State[T] {
	return State[T]{
		sm:    s.sm,
		state: s.sm.matchState(str, s.state),
	}
}

var ErrAmbiguous = errors.New("ambiguous states")
