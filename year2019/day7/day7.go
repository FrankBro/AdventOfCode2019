package day7

import (
	"strconv"
	"strings"
	"sync"

	"github.com/FrankBro/AdventOfCodeGo/intcode"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func maximize(codes []int) int {
	var max int
	for _, permutation := range permutations([]int{0, 1, 2, 3, 4}) {
		var output int
		var wg sync.WaitGroup
		for i, input := range permutation {
			codesCopy := make([]int, len(codes))
			copy(codesCopy, codes)
			inputChan := make(chan int, 2)
			inputChan <- input
			inputChan <- output
			outputChan := make(chan int, 1)
			wg.Add(1)
			computer := intcode.NewComputer(i, codesCopy, inputChan, outputChan, &wg)
			go computer.Eval()
			wg.Wait()
			output = <-outputChan
		}
		if output > max {
			max = output
		}
	}
	return max
}

func maximizeString(line string) int {
	codes := readString(line)
	output := maximize(codes)
	return output
}

func maximizeFeedback(codes []int) int {
	var max int
	for _, permutation := range permutations([]int{5, 6, 7, 8, 9}) {
		inputs := make([]chan int, len(permutation))
		outputs := make([]chan int, len(permutation))
		for i := range permutation {
			outputs[i] = make(chan int, 2)
		}
		for i, input := range permutation {
			if i == 0 {
				inputs[i] = outputs[4]
				inputs[i] <- input
				inputs[i] <- 0
			} else {
				inputs[i] = outputs[i-1]
				inputs[i] <- input
			}
		}
		var wg sync.WaitGroup
		wg.Add(5)
		for i := range permutation {
			codesCopy := make([]int, len(codes))
			copy(codesCopy, codes)
			computer := intcode.NewComputer(i, codesCopy, inputs[i], outputs[i], &wg)
			go computer.Eval()
		}
		wg.Wait()
		output := <-outputs[4]
		if output > max {
			max = output
		}
	}
	return max
}

func maximizeStringFeedback(line string) int {
	codes := readString(line)
	output := maximizeFeedback(codes)
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
