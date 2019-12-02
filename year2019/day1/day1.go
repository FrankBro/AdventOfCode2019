package day1

import (
	"strconv"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func calculateFuel(mass int) int {
	return mass/3 - 2
}

func recursivelyCalculateFuel(mass int) (sum int) {
	for mass > 0 {
		mass = (mass / 3) - 2
		if mass > 0 {
			sum += mass
		}
	}
	return sum
}

func readMasses() []int {
	lines := util.ReadLines()
	masses := make([]int, len(lines))
	for i, line := range lines {
		mass, err := strconv.Atoi(line)
		util.Check(err)
		masses[i] = mass
	}
	return masses
}
