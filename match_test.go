package match

import (
	"errors"
	"testing"
)

func TestMatch(t *testing.T) {
	var (
		matches   [][]string
		nomatches [][]string

		sm = New[int]()
	)

	for n, test := range [...]struct {
		Add     string
		Err     error
		Match   []string
		NoMatch []string
	}{
		{
			Add:     "abc",
			Match:   []string{"abc"},
			NoMatch: []string{"ab", "abd", "abcd"},
		},
		{
			Add: "abc",
			Err: ErrAmbiguous,
		},
		{
			Add:     "def",
			Match:   []string{"def"},
			NoMatch: []string{"ab", "abd", "abcd", "de", "deg", "defg"},
		},
	} {
		if err := sm.AddString(test.Add, n+1); !errors.Is(err, test.Err) {
			t.Errorf("test %d: expecting error %v, got %v", n+1, test.Err, err)
		} else {
			matches = append(matches, test.Match)
			nomatches = append(nomatches, test.NoMatch)

			for m, toMatch := range matches {
				for l, match := range toMatch {
					if v := sm.Match(match); v != m+1 {
						t.Errorf("test %d.%d.%d: expecting value %d, got %d", n+1, m+1, l+1, m+1, v)
					}
				}
			}

			for m, toNotMatch := range nomatches {
				for l, match := range toNotMatch {
					if v := sm.Match(match); v == m+1 {
						t.Errorf("test %d.%d.%d: expecting to not get value %d, but did", n+1, m+1, l+1, v)
					}
				}
			}
		}
	}
}
