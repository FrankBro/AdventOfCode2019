package day2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/FrankBro/AdventOfCodeGo/intcode"
	"github.com/FrankBro/AdventOfCodeGo/util"
)

func eval(codes []int) []int {
	var wg sync.WaitGroup
	inputChan := make(chan int, 2)
	outputChan := make(chan int, 1)
	wg.Add(1)
	computer := intcode.NewComputer(0, codes, inputChan, outputChan, &wg)
	go computer.Eval()
	wg.Wait()
	return codes
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
