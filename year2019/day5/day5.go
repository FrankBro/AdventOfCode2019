package day5

import (
	"strconv"
	"strings"
	"sync"

	"github.com/FrankBro/AdventOfCodeGo/intcode"
	"github.com/FrankBro/AdventOfCodeGo/util"
)

func eval1(codes []int, input int) int {
	var wg sync.WaitGroup
	inputChan := make(chan int, 2)
	outputChan := make(chan int, 20)
	wg.Add(1)
	inputChan <- input
	computer := intcode.NewComputer(0, codes, inputChan, outputChan, &wg)
	go computer.Eval()
	wg.Wait()
	for {
		output := <-outputChan
		if output != 0 {
			return output
		}
	}
}

func eval2(codes []int, input int) int {
	var wg sync.WaitGroup
	inputChan := make(chan int, 2)
	outputChan := make(chan int, 20)
	wg.Add(1)
	inputChan <- input
	computer := intcode.NewComputer(0, codes, inputChan, outputChan, &wg)
	go computer.Eval()
	wg.Wait()
	return <-outputChan
}

func eval2String(line string, input int) int {
	codes := readString(line)
	output := eval2(codes, input)
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
