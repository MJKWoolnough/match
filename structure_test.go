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
		Output    *or[byte]
	}{
		{ // 1
			Input:     "",
			Tokeniser: simpleStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 2
			Input:     "a",
			Tokeniser: simpleStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 3
			Input:     "abc",
			Tokeniser: simpleStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'b': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'c': true},
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 4
			Input:     "",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 5
			Input:     "a",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 6
			Input:     "abc",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'b': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'c': true},
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 7
			Input:     "*",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 8
			Input:     "*abc",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'b': true},
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'c': true},
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
		{ // 9
			Input:     "*a*b*c*",
			Tokeniser: partialStringStart,
			Output: &or[byte]{
				sequences: []sequence[byte]{
					{
						parts: []part[byte]{
							{
								partType: partStart,
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'a': true},
								},
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'b': true},
								},
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								char: &char[byte]{
									char: [256]bool{'c': true},
								},
							},
							{
								partType: partMany,
								char: &char[byte]{
									invert: true,
								},
							},
							{
								partType: partEnd,
							},
						},
					},
				},
			},
		},
	} {
		o, err := parse[byte](test.Input, test.Tokeniser)
		if !errors.Is(err, test.Err) {
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
		} else if !reflect.DeepEqual(o, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, o)
		}
	}
}
