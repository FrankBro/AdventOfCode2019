package intcode

import (
	"sync"
)

type Computer struct {
	id     int
	input  chan int
	output chan int
	codes  []int
	wg     *sync.WaitGroup
}

func NewComputer(id int, codes []int, input, output chan int, wg *sync.WaitGroup) *Computer {
	computer := Computer{
		id:     id,
		input:  input,
		output: output,
		codes:  codes,
		wg:     wg,
	}
	return &computer
}

func (computer *Computer) get(immediateMode bool, addr int) int {
	if immediateMode {
		return computer.codes[addr]
	}
	return computer.codes[computer.codes[addr]]
}

func (computer *Computer) Eval() {
	var addr int
	for {
		op := computer.codes[addr] % 100
		switch op {
		case 1:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			computer.codes[computer.codes[addr+3]] = a1 + a2
			addr += 4
		case 2:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			computer.codes[computer.codes[addr+3]] = a1 * a2
			addr += 4
		case 3:
			// fmt.Printf("Computer %d wants to read input\n", computer.id)
			input := <-computer.input
			// fmt.Printf("Computer %d read input %d\n", computer.id, input)
			computer.codes[computer.codes[addr+1]] = input
			addr += 2
		case 4:
			m1 := computer.codes[addr]/100%10 == 1
			a1 := computer.get(m1, addr+1)
			// fmt.Printf("Computer %d wants to write output %d\n", computer.id, a1)
			computer.output <- a1
			// fmt.Printf("Computer %d wrote output %d\n", computer.id, a1)
			addr += 2
		case 5:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			if a1 != 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 6:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			if a1 == 0 {
				addr = a2
			} else {
				addr += 3
			}
		case 7:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			var value int
			if a1 < a2 {
				value = 1
			}
			computer.codes[computer.codes[addr+3]] = value
			addr += 4
		case 8:
			m1 := computer.codes[addr]/100%10 == 1
			m2 := computer.codes[addr]/1000%10 == 1
			a1 := computer.get(m1, addr+1)
			a2 := computer.get(m2, addr+2)
			var value int
			if a1 == a2 {
				value = 1
			}
			computer.codes[computer.codes[addr+3]] = value
			addr += 4
		case 99:
			computer.wg.Done()
			return
		}
	}
}
