package day3

import (
	"github.com/FrankBro/AdventOfCodeGo/util"
)

type pos struct {
	x int
	y int
}

func addVisit(pos pos, houses map[pos]int) {
	if _, ok := houses[pos]; ok {
		houses[pos]++
	} else {
		houses[pos] = 1
	}
}

func travel(route string) map[pos]int {
	houses := make(map[pos]int)
	lastPos := pos{x: 0, y: 0}
	addVisit(lastPos, houses)
	for _, dir := range route {
		switch dir {
		case '>':
			lastPos.x++
		case '<':
			lastPos.x--
		case '^':
			lastPos.y++
		case 'v':
			lastPos.y--
		}
		addVisit(lastPos, houses)
	}
	return houses
}

func travelled(route string) int {
	houses := travel(route)
	return len(houses)
}

func roboTravel(route string) map[pos]int {
	houses := make(map[pos]int)
	santaPos := pos{x: 0, y: 0}
	roboPos := pos{x: 0, y: 0}
	addVisit(santaPos, houses)
	for i, dir := range route {
		switch dir {
		case '>':
			if i%2 == 0 {
				santaPos.x++
			} else {
				roboPos.x++
			}
		case '<':
			if i%2 == 0 {
				santaPos.x--
			} else {
				roboPos.x--
			}
		case '^':
			if i%2 == 0 {
				santaPos.y++
			} else {
				roboPos.y++
			}
		case 'v':
			if i%2 == 0 {
				santaPos.y--
			} else {
				roboPos.y--
			}
		}
		if i%2 == 0 {
			addVisit(santaPos, houses)
		} else {
			addVisit(roboPos, houses)
		}
	}
	return houses
}

func roboTravelled(route string) int {
	return len(roboTravel(route))
}

func readRoute() string {
	return util.ReadLine()
}
