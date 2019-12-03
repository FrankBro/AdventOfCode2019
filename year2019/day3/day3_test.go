package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, 6, closest("R8,U5,L5,D3", "U7,R6,D4,L4"))
	require.EqualValues(t, 159, closest("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"))
	require.EqualValues(t, 135, closest(
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))
	a, b := readPath()
	require.EqualValues(t, 207, closest(a, b))
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, 30, closestPath("R8,U5,L5,D3", "U7,R6,D4,L4"))
	require.EqualValues(t, 610, closestPath("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"))
	require.EqualValues(t, 410, closestPath(
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))
	a, b := readPath()
	require.EqualValues(t, 21196, closestPath(a, b))
}
