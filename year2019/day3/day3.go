package day3

import (
	"math"
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

type pos struct {
	x int
	y int
}

type tile struct {
	a int
	b int
}

type wire int

const (
	wireA wire = iota
	wireB
)

func passed(wire wire, pos pos, cost int, grid map[pos]tile) {
	// fmt.Printf("x: %d, y: %d, wire: %d\n", pos.x, pos.y, wire)
	switch wire {
	case wireA:
		tile := grid[pos]
		if tile.a == 0 {
			tile.a = cost
		}
		grid[pos] = tile
	case wireB:
		tile := grid[pos]
		if tile.b == 0 {
			tile.b = cost
		}
		grid[pos] = tile
	}
}

func apply(wire wire, instruction string, grid map[pos]tile) {
	steps := strings.Split(instruction, ",")
	pos := pos{x: 0, y: 0}
	var cost int
	for _, step := range steps {
		amount, err := strconv.Atoi(step[1:])
		util.Check(err)
		switch step[0] {
		case 'U':
			for y := 1; y <= amount; y++ {
				pos.y++
				cost++
				passed(wire, pos, cost, grid)
			}
		case 'R':
			for x := 1; x <= amount; x++ {
				pos.x++
				cost++
				passed(wire, pos, cost, grid)
			}
		case 'L':
			for x := 1; x <= amount; x++ {
				pos.x--
				cost++
				passed(wire, pos, cost, grid)
			}
		case 'D':
			for y := 1; y <= amount; y++ {
				pos.y--
				cost++
				passed(wire, pos, cost, grid)
			}
		}
	}
}

func crosses(grid map[pos]tile) map[pos]tile {
	crosses := make(map[pos]tile)
	for pos, tile := range grid {
		if tile.a != 0 && tile.b != 0 {
			crosses[pos] = tile
		}
	}
	return crosses
}

func closest(a, b string) int {
	grid := make(map[pos]tile)
	apply(wireA, a, grid)
	apply(wireB, b, grid)
	crosses := crosses(grid)
	min := math.MaxInt32
	for pos := range crosses {
		dx := pos.x
		if dx < 0 {
			dx = -dx
		}
		dy := pos.y
		if dy < 0 {
			dy = -dy
		}
		distance := dx + dy
		if distance < min && distance != 0 {
			min = distance
		}
	}
	return min
}

func closestPath(a, b string) int {
	grid := make(map[pos]tile)
	apply(wireA, a, grid)
	apply(wireB, b, grid)
	crosses := crosses(grid)
	min := math.MaxInt32
	for _, cross := range crosses {
		cost := cross.a + cross.b
		if cost < min && cost != 0 {
			min = cost
		}
	}
	return min
}

func readPath() (string, string) {
	lines := util.ReadLines()
	return lines[0], lines[1]
}
