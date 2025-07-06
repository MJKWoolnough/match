package match

import (
	"errors"
	"reflect"
	"testing"
)

func TestStructure(t *testing.T) {
	for n, test := range [...]struct {
		Input  string
		Err    error
		Output *or
	}{
		{
			Input: "",
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
			Input: "a",
			Output: &or{
				sequences: []sequence{
					{
						parts: []part{
							{
								start: true,
							},
							{
								char: &char{
									char: "a",
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
			Input: "abc",
			Output: &or{
				sequences: []sequence{
					{
						parts: []part{
							{
								start: true,
							},
							{
								char: &char{
									char: "a",
								},
							},
							{
								char: &char{
									char: "b",
								},
							},
							{
								char: &char{
									char: "c",
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
		o, err := parse(test.Input)
		if !errors.Is(err, test.Err) {
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
		} else if !reflect.DeepEqual(o, test.Output) {
			t.Errorf("test %d: expecting output %v, got %v", n+1, test.Output, o)
		}
	}
}
