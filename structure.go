package match

import (
	"unsafe"

	"vimagination.zapto.org/parser"
)

type visitedSet map[uint32]struct{}

func parse[T comparable](str string, fn parser.TokenFunc) (*or[T], error) {
	tk := parser.NewStringTokeniser(str)

	tk.TokeniserState(fn)

	p := parser.New(tk)

	var or or[T]

	if err := or.parse(&p); err != nil {
		return nil, err
	}

	return &or, nil
}

type or[T comparable] struct {
	sequences []sequence[T]
}

func (o *or[T]) parse(p *parser.Parser) error {
	var s sequence[T]

	if err := s.parse(p); err != nil {
		return err
	}

	o.sequences = append(o.sequences, s)

	return nil
}

func (o *or[T]) compile(sm *StateMachine[T], state uint32, visited visitedSet, value T) ([]*uint32, error) {
	return nil, nil
}

type sequence[T comparable] struct {
	parts []part[T]
}

func (s *sequence[T]) parse(p *parser.Parser) error {
	for {
		if p.Peek().Type == parser.TokenDone {
			break
		}

		var pt part[T]

		if err := pt.parse(p); err != nil {
			return err
		}

		if p.Accept(tokenRepeat) {
			pt.partType = partMany
		}

		s.parts = append(s.parts, pt)
	}

	return nil
}

func (s *sequence[T]) compile(sm *StateMachine[T], state uint32, visited visitedSet, value T) ([]*uint32, error) {
	return nil, nil
}

type partType uint8

const (
	partOne partType = iota
	partMany
	partStart
	partEnd
)

type part[T comparable] struct {
	partType
	char *char[T]
}

func (pt *part[T]) parse(p *parser.Parser) error {
	if p.Accept(tokenStart) {
		pt.partType = partStart
	} else if p.Accept(tokenEnd) {
		pt.partType = partEnd
	} else {
		pt.char = new(char[T])

		if err := pt.char.parse(p); err != nil {
			return err
		}
	}

	return nil
}

func (pt *part[T]) compile(sm *StateMachine[T], state uint32, visited visitedSet, value T) ([]*uint32, error) {
	return nil, nil
}

type char[T comparable] struct {
	invert bool
	char   [256]bool
}

func (c *char[T]) parse(p *parser.Parser) error {
	tk := p.Next()

	if tk.Type == parser.TokenError {
		return p.GetError()
	} else if tk.Type == tokenAnyChar {
		c.invert = true
	}

	for _, b := range unsafe.Slice(unsafe.StringData(tk.Data), len(tk.Data)) {
		c.char[b] = true
	}

	return nil
}

func (c *char[T]) compile(sm *StateMachine[T], state uint32, visited visitedSet, value T) ([]*uint32, error) {
	return nil, nil
}
