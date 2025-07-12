package match

import (
	"errors"

	"vimagination.zapto.org/parser"
)

const (
	tokenStart parser.TokenType = iota
	tokenEnd
	tokenChar
	tokenAnyChar
	tokenRepeat
)

func simpleStart(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return t.Return(tokenStart, simpleString)
}

func simpleString(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	if t.Next() == -1 {
		return t.Return(tokenEnd, (*parser.Tokeniser).Done)
	}

	return t.Return(tokenChar, simpleString)
}

func partialStringStart(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	return t.Return(tokenStart, partialString)
}

func partialString(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	switch t.Peek() {
	case -1:
		return t.Return(tokenEnd, (*parser.Tokeniser).Done)
	case '*':
		return t.Return(tokenAnyChar, partialStringWildcard)
	case '\\':
		t.Next()
		t.Get()

		if t.Peek() == -1 {
			return t.ReturnError(ErrInvalidEscape)
		}
	}

	t.Next()

	return t.Return(tokenChar, partialString)
}

func partialStringWildcard(t *parser.Tokeniser) (parser.Token, parser.TokenFunc) {
	t.Next()

	return t.Return(tokenRepeat, partialString)
}

var ErrInvalidEscape = errors.New("invalid escape sequence")
