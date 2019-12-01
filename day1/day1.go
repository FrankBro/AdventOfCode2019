package day1

import (
	"bufio"
	"os"
	"strconv"

	"github.com/FrankBro/aoc19/util"
)

func calculateFuel(mass int) int {
	return mass/3 - 2
}

func recursivelyCalculateFuel(mass int) (sum int) {
	for {
		mass = (mass / 3) - 2
		if mass < 0 {
			break
		}
		sum += mass
		if mass < 3 {
			break
		}
	}
	return sum
}

func readMasses() []int {
	masses := make([]int, 0)
	f, err := os.Open("part1.txt")
	util.Check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		util.Check(err)
		masses = append(masses, mass)
	}
	return masses
}