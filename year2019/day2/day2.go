package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func eval(codes []int) []int {
	var addr int
	for {
		switch codes[addr] {
		case 1:
			codes[codes[addr+3]] = codes[codes[addr+1]] + codes[codes[addr+2]]
		case 2:
			codes[codes[addr+3]] = codes[codes[addr+1]] * codes[codes[addr+2]]
		case 99:
			return codes
		}
		addr += 4
	}
}

func evalString(line string) string {
	codes := readString(line)
	codes = eval(codes)
	var s string
	for i, code := range codes {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("%d", code)
	}
	return s
}

func readString(line string) []int {
	splits := strings.Split(line, ",")
	codes := make([]int, len(splits))
	for i, split := range splits {
		code, err := strconv.Atoi(split)
		util.Check(err)
		codes[i] = code
	}
	return codes
}

func readInput() []int {
	line := util.ReadLine()
	return readString(line)
}
