package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, true, valid(111111))
	require.EqualValues(t, false, valid(223450))
	require.EqualValues(t, false, valid(123789))
	require.EqualValues(t, 2081, part1())
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, true, valid2(112233))
	require.EqualValues(t, false, valid2(123444))
	require.EqualValues(t, true, valid2(111122))
	require.EqualValues(t, false, valid2(344445))
	require.EqualValues(t, 1411, part2())
}
