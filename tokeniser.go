package match

import "vimagination.zapto.org/parser"

const (
	tokenStart parser.TokenType = iota
	tokenChar
	tokenEnd
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
