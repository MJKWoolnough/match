package match

import (
	"testing"

	"vimagination.zapto.org/parser"
)

func TestTokeniser(t *testing.T) {
	for n, test := range [...]struct {
		Input        string
		Output       []parser.Token
		InitialState parser.TokenFunc
	}{
		{
			"",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			simpleStart,
		},
		{
			"a",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			simpleStart,
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
			simpleStart,
		},
		{
			"",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
		{
			"a",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
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
			partialStringStart,
		},
		{
			"*",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenAnyChar, Data: ""},
				{Type: tokenRepeat, Data: "*"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
		{
			"*abc",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenAnyChar, Data: ""},
				{Type: tokenRepeat, Data: "*"},
				{Type: tokenChar, Data: "a"},
				{Type: tokenChar, Data: "b"},
				{Type: tokenChar, Data: "c"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
		{
			"*abc*",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenAnyChar, Data: ""},
				{Type: tokenRepeat, Data: "*"},
				{Type: tokenChar, Data: "a"},
				{Type: tokenChar, Data: "b"},
				{Type: tokenChar, Data: "c"},
				{Type: tokenAnyChar, Data: ""},
				{Type: tokenRepeat, Data: "*"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
	} {
		p := parser.NewStringTokeniser(test.Input)

		p.TokeniserState(test.InitialState)

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
