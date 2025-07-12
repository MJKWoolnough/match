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
		{ // 1
			"",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			simpleStart,
		},
		{ // 2
			"a",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			simpleStart,
		},
		{ // 3
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
		{ // 4
			"",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
		{ // 5
			"a",
			[]parser.Token{
				{Type: tokenStart, Data: ""},
				{Type: tokenChar, Data: "a"},
				{Type: tokenEnd, Data: ""},
				{Type: parser.TokenDone, Data: ""},
			},
			partialStringStart,
		},
		{ // 6
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
		{ // 7
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
		{ // 8
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
		{ // 9
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
