package match

import (
	"testing"

	"vimagination.zapto.org/parser"
)

func TestStringTokeniser(t *testing.T) {
	for n, test := range [...]struct {
		Input  string
		Output []parser.Token
	}{
		{
			"",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{
			"a",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
		{
			"abc",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenChar, Data: "b"},
				{Type: tokenChar, Data: "c"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
		},
	} {
		p := parser.NewStringTokeniser(test.Input)

		p.TokeniserState(simpleStart)

		for m, tkn := range test.Output {
			if tk, _ := p.GetToken(); tk.Type != tkn.Type {
				if tk.Type == parser.TokenError {
					t.Errorf("test %d.%d: unexpected error: %s", n+1, m+1, tk.Data)
				} else {
					t.Errorf("test %d.%d: Incorrect type, expecting %d, got %d", n+1, m+1, tkn.Type, tk.Type)
				}

				break
			} else if tk.Data != tkn.Data {
				t.Errorf("test %d.%d: Incorrect data, expecting %q, got %q", n+1, m+1, tkn.Data, tk.Data)

				break
			}
		}
	}
}
