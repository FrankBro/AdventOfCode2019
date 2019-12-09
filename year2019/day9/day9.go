package day9

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/FrankBro/AdventOfCodeGo/intcode"
	"github.com/FrankBro/AdventOfCodeGo/util"
)

func run(input string, inputs ...int64) string {
	codes := readString(input)
	var wg sync.WaitGroup
	inputChan := make(chan int64, len(inputs))
	for _, input := range inputs {
		inputChan <- input
	}
	outputChan := make(chan int64, 100)
	wg.Add(1)
	computer := intcode.NewComputer(0, codes, inputChan, outputChan, &wg)
	go computer.Eval()
	wg.Wait()
	var exit bool
	outputs := make([]string, 0)
	for {
		if exit {
			break
		}
		select {
		case output := <-outputChan:
			outputs = append(outputs, fmt.Sprintf("%d", output))
		default:
			exit = true
		}
	}
	return strings.Join(outputs, ",")
}

func readString(line string) []int64 {
	splits := strings.Split(line, ",")
	codes := make([]int64, len(splits))
	for i, split := range splits {
		code, err := strconv.ParseInt(split, 10, 64)
		util.Check(err)
		codes[i] = code
	}
	return codes
}

func getInput() string {
	return util.ReadLine()
}
