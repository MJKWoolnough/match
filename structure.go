package match

import (
	"unsafe"

	"vimagination.zapto.org/parser"
)

func parse(str string, fn parser.TokenFunc) (*or, error) {
	tk := parser.NewStringTokeniser(str)

	tk.TokeniserState(fn)

	p := parser.New(tk)

	var or or

	if err := or.parse(&p); err != nil {
		return nil, err
	}

	return &or, nil
}

type or struct {
	sequences []sequence
}

func (o *or) parse(p *parser.Parser) error {
	var s sequence

	if err := s.parse(p); err != nil {
		return err
	}

	o.sequences = append(o.sequences, s)

	return nil
}

type sequence struct {
	parts []part
}

func (s *sequence) parse(p *parser.Parser) error {
	for {
		if p.Peek().Type == parser.TokenDone {
			break
		}

		var pt part

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

type partType uint8

const (
	partOne partType = iota
	partMany
	partStart
	partEnd
)

type part struct {
	partType
	char *char
}

func (pt *part) parse(p *parser.Parser) error {
	if p.Accept(tokenStart) {
		pt.partType = partStart
	} else if p.Accept(tokenEnd) {
		pt.partType = partEnd
	} else {
		pt.char = new(char)

		if err := pt.char.parse(p); err != nil {
			return err
		}
	}

	return nil
}

type char struct {
	invert bool
	char   [256]bool
}

func (c *char) parse(p *parser.Parser) error {
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
