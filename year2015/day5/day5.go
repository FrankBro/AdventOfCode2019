package day5

import (
	"github.com/FrankBro/AdventOfCodeGo/util"
)

type state struct {
	vowels int
	double bool
	bad    bool
	last   rune
}

func (state *state) checkVowel(r rune) {
	switch r {
	case 'a':
		state.vowels++
	case 'e':
		state.vowels++
	case 'i':
		state.vowels++
	case 'o':
		state.vowels++
	case 'u':
		state.vowels++
	}
}

func (state *state) checkDouble(r rune) {
	if r == state.last {
		state.double = true
	}
}

func (state *state) checkBad(r rune) {
	switch {
	case state.last == 'a' && r == 'b':
		state.bad = true
	case state.last == 'c' && r == 'd':
		state.bad = true
	case state.last == 'p' && r == 'q':
		state.bad = true
	case state.last == 'x' && r == 'y':
		state.bad = true
	}
}

func isNice(s string) bool {
	var state state
	for _, r := range s {
		state.checkVowel(r)
		state.checkDouble(r)
		state.checkBad(r)
		if state.bad {
			return false
		}
		state.last = r
	}
	return state.vowels >= 3 && state.double
}

type state2 struct {
	firstDouble map[string]int
	double      bool
	mirror      bool
	last        rune
	last2       rune
}

func (state *state2) checkDouble(r rune, pos int) {
	if state.last == rune(0) {
		return
	}
	s := string(r) + string(state.last)
	if firstPos, ok := state.firstDouble[s]; ok {
		if r != state.last || firstPos != pos-1 {
			state.double = true
		}
	} else {
		state.firstDouble[s] = pos
	}

}

func (state *state2) checkMirror(r rune) {
	if state.last2 == r {
		state.mirror = true
	}
}

func isNice2(s string) bool {
	state := state2{
		firstDouble: make(map[string]int),
	}
	for i, r := range s {
		state.checkDouble(r, i)
		state.checkMirror(r)
		state.last2 = state.last
		state.last = r
	}
	return state.double && state.mirror
}

func readStrings() []string {
	return util.ReadLines()
}
