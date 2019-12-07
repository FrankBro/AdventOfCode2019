package day6

import (
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func count(origin string, orbits []string, connections map[string][]string) (sum int) {
	for _, orbit := range orbits {
		var amount int
		amount += len(orbits)
		sub := connections[orbit]
		amount += count(origin, sub, connections)
		sum += amount
	}
	return sum
}

func orbits(links []string) int {
	connections := make(map[string][]string)
	for _, link := range links {
		split := strings.Split(link, ")")
		origin, destination := split[0], split[1]
		if orbits, ok := connections[destination]; ok {
			connections[destination] = append(orbits, origin)
		} else {
			connections[destination] = []string{origin}
		}
	}
	var sum int
	for destination, orbits := range connections {
		sum += count(destination, orbits, connections)
	}
	return sum
}

func closest(links []string) int {
	path := make(map[string]string)
	for _, link := range links {
		split := strings.Split(link, ")")
		origin, destination := split[0], split[1]
		path[destination] = origin
	}
	next := path["YOU"]
	var distance int
	distances := make(map[string]int)
	for {
		distances[next] = distance
		if parent, ok := path[next]; ok {
			next = parent
			distance++
		} else {
			break
		}
	}
	next = path["SAN"]
	distance = 0
	for {
		if pathDistance, ok := distances[next]; ok {
			distance += pathDistance
			break
		} else {
			parent := path[next]
			next = parent
			distance++
		}
	}

	return distance
}

func getData() []string {
	return util.ReadLines()
}
