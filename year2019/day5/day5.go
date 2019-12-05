package day5

import (
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func get(immediateMode bool, codes []int, addr int) int {
	if immediateMode {
		return codes[addr]
	}
	return codes[codes[addr]]
}

func eval(codes []int, input int) int {
	var output int
	var addr int
	for {
		op := codes[addr] % 100
		switch op {
		case 1:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			codes[codes[addr+3]] = a1 + a2
			addr += 4
		case 2:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			codes[codes[addr+3]] = a1 * a2
			addr += 4
		case 3:
			codes[codes[addr+1]] = input
			addr += 2
		case 4:
			m1 := codes[addr]/100%10 == 1
			a1 := get(m1, codes, addr+1)
			output = a1
			addr += 2
		case 5:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			if a1 != 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 6:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			if a1 == 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 7:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			var value int
			if a1 < a2 {
				value = 1
			}
			codes[codes[addr+3]] = value
			addr += 4
		case 8:
			m1 := codes[addr]/100%10 == 1
			m2 := codes[addr]/1000%10 == 1
			a1 := get(m1, codes, addr+1)
			a2 := get(m2, codes, addr+2)
			var value int
			if a1 == a2 {
				value = 1
			}
			codes[codes[addr+3]] = value
			addr += 4
		case 99:
			return output
		}
	}
}

func evalString(line string, input int) int {
	codes := readString(line)
	output := eval(codes, input)
	return output
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
