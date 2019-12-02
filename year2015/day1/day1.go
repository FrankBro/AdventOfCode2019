package day1

import (
	"github.com/FrankBro/AdventOfCodeGo/util"
)

func findFloor(s string) int {
	var floor int
	for _, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			panic("invalid character")
		}
	}
	return floor
}

func findPosition(s string) int {
	var position int
	var floor int
	for _, c := range s {
		position++
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		default:
			panic("invalid character")
		}
		if floor == -1 {
			break
		}
	}
	return position
}

func readInstruction() string {
	return util.ReadLine()
}
