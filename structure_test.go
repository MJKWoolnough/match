package match

import (
	"errors"
	"reflect"
	"testing"

	"vimagination.zapto.org/parser"
)

func TestStructure(t *testing.T) {
	for n, test := range [...]struct {
		Input     string
		Tokeniser parser.TokenFunc
		Err       error
		Output    *or
	}{
		{
			Input:     "",
			Tokeniser: simpleStart,
			Output: &or{
				sequences: []sequence{
					{
						parts: []part{
							{
								start: true,
							},
							{
								end: true,
							},
						},
					},
				},
			},
		},
		{
			Input:     "a",
			Tokeniser: simpleStart,
			Output: &or{
				sequences: []sequence{
					{
						parts: []part{
							{
								start: true,
							},
							{
								char: &char{
									char: [256]bool{'a': true},
								},
							},
							{
								end: true,
							},
						},
					},
				},
			},
		},
		{
			Input:     "abc",
			Tokeniser: simpleStart,
			Output: &or{
				sequences: []sequence{
					{
						parts: []part{
							{
								start: true,
							},
							{
								char: &char{
									char: [256]bool{'a': true},
								},
							},
							{
								char: &char{
									char: [256]bool{'b': true},
								},
							},
							{
								char: &char{
									char: [256]bool{'c': true},
								},
							},
							{
								end: true,
							},
						},
					},
				},
			},
		},
	} {
		o, err := parse(test.Input, test.Tokeniser)
		if !errors.Is(err, test.Err) {
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
		} else if !reflect.DeepEqual(o, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, o)
		}
	}
}
