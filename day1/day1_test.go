package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, 2, calculateFuel(12))
	require.EqualValues(t, 2, calculateFuel(14))
	require.EqualValues(t, 654, calculateFuel(1969))
	require.EqualValues(t, 33583, calculateFuel(100756))
	masses := readMasses()
	var sum int
	for _, mass := range masses {
		sum += calculateFuel(mass)
	}
	require.EqualValues(t, 3256794, sum)
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, 2, recursivelyCalculateFuel(14))
	require.EqualValues(t, 966, recursivelyCalculateFuel(1969))
	require.EqualValues(t, 50346, recursivelyCalculateFuel(100756))
	masses := readMasses()
	var sum int
	for _, mass := range masses {
		sum += recursivelyCalculateFuel(mass)
	}
	require.EqualValues(t, 4882337, sum)
}
