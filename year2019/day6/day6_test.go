package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testData := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	require.EqualValues(t, 42, orbits(testData))
	require.EqualValues(t, 106065, orbits(getData()))
}

func TestPart2(t *testing.T) {
	testData := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	require.EqualValues(t, 4, closest(testData))
	require.EqualValues(t, 253, closest(getData()))
}
