package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, 2, travelled(">"))
	require.EqualValues(t, 4, travelled("^>v<"))
	require.EqualValues(t, 2, travelled("^v^v^v^v^v"))
	route := readRoute()
	require.EqualValues(t, 2592, travelled(route))
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, 3, roboTravelled("^v"))
	require.EqualValues(t, 3, roboTravelled("^>v<"))
	require.EqualValues(t, 11, roboTravelled("^v^v^v^v^v"))
	route := readRoute()
	require.EqualValues(t, 2360, roboTravelled(route))
}
