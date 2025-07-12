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
		{ // 1
			Add:     "abc",
			Match:   []string{"abc"},
			NoMatch: []string{"ab", "abd", "abcd"},
		},
		{ // 1
			Add: "abc",
			Err: ErrAmbiguous,
		},
		{ // 1
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

func TestState(t *testing.T) {
	sm := New[int]()

	sm.AddString("abcde", 1)

	if state := sm.MatchState("b"); state != (State[int]{sm, 0}) {
		t.Errorf("test 1: expecting state 0, got %d", state.state)
	}

	if state := sm.MatchState("a"); state != (State[int]{sm, 2}) {
		t.Errorf("test 2: expecting state 2, got %d", state.state)
	} else if state = state.MatchState("b"); state != (State[int]{sm, 3}) {
		t.Errorf("test 3: expecting state 3, got %d", state.state)
	} else if v := state.Match("c"); v != 0 {
		t.Errorf("test 4: expecting value 0, got %d", v)
	} else if v := state.Match("cde"); v != 1 {
		t.Errorf("test 5: expecting value 1, got %d", v)
	}
}
