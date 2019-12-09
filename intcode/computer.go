package intcode

import (
	"math"
	"sync"
)

type Computer struct {
	id           int
	input        chan int64
	output       chan int64
	codes        []int64
	wg           *sync.WaitGroup
	relativeBase int64
}

func NewComputer(id int, codes []int64, input, output chan int64, wg *sync.WaitGroup) *Computer {
	computer := Computer{
		id:     id,
		input:  input,
		output: output,
		codes:  codes,
		wg:     wg,
	}
	return &computer
}

func (computer *Computer) alloc(index int64) {
	codes := make([]int64, index+1)
	copy(codes, computer.codes)
	computer.codes = codes
}

func (computer *Computer) getOrAlloc(index int64) int64 {
	if int(index) >= len(computer.codes) {
		computer.alloc(index)
	}
	return computer.codes[index]
}

func (computer *Computer) setOrAlloc(index int64, value int64) {
	if int(index) >= len(computer.codes) {
		computer.alloc(index)
	}
	computer.codes[index] = value
}

func (computer *Computer) get(addr int64, arg int64) int64 {
	mode := (computer.getOrAlloc(addr) / int64(math.Pow10(int(arg)+1))) % 10
	switch mode {
	case 0:
		return computer.getOrAlloc(computer.getOrAlloc(addr + arg))
	case 1:
		return computer.getOrAlloc(addr + arg)
	case 2:
		return computer.getOrAlloc(computer.relativeBase + computer.getOrAlloc(addr+arg))
	default:
		panic("unknown mode")
	}
}

func (computer *Computer) set(addr int64, arg int64, value int64) {
	mode := (computer.getOrAlloc(addr) / int64(math.Pow10(int(arg)+1))) % 10
	switch mode {
	case 0:
		computer.setOrAlloc(computer.getOrAlloc(addr+arg), value)
	case 1:
		panic("immediate mode writing")
	case 2:
		computer.setOrAlloc(computer.relativeBase+computer.getOrAlloc(addr+arg), value)
	default:
		panic("unknown mode")
	}
}

func (computer *Computer) Eval() {
	var addr int64
	for {
		op := computer.codes[addr] % 100
		switch op {
		case 1:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			computer.set(addr, 3, a1+a2)
			addr += 4
		case 2:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			computer.set(addr, 3, a1*a2)
			addr += 4
		case 3:
			// fmt.Printf("Computer %d wants to read input\n", computer.id)
			input := <-computer.input
			// fmt.Printf("Computer %d read input %d\n", computer.id, input)
			computer.set(addr, 1, input)
			addr += 2
		case 4:
			a1 := computer.get(addr, 1)
			// fmt.Printf("Computer %d wants to write output %d\n", computer.id, a1)
			computer.output <- a1
			// fmt.Printf("Computer %d wrote output %d\n", computer.id, a1)
			addr += 2
		case 5:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			if a1 != 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 6:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			if a1 == 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 7:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			var value int64
			if a1 < a2 {
				value = 1
			}
			computer.set(addr, 3, value)
			addr += 4
		case 8:
			a1 := computer.get(addr, 1)
			a2 := computer.get(addr, 2)
			var value int64
			if a1 == a2 {
				value = 1
			}
			computer.set(addr, 3, value)
			addr += 4
		case 9:
			a1 := computer.get(addr, 1)
			computer.relativeBase += a1
			addr += 2
		case 99:
			computer.wg.Done()
			return
		}
	}
}
