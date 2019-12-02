package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	{
		present := readPresent("2x3x4")
		require.EqualValues(t, 58, present.wrappingPaper())
	}
	{
		present := readPresent("1x1x10")
		require.EqualValues(t, 43, present.wrappingPaper())
	}
	presents := readPresents()
	var sum int
	for _, present := range presents {
		sum += present.wrappingPaper()
	}
	require.EqualValues(t, 1598415, sum)
}

func TestPart2(t *testing.T) {
	{
		present := readPresent("2x3x4")
		require.EqualValues(t, 34, present.ribbonAndBow())
	}
	{
		present := readPresent("1x1x10")
		require.EqualValues(t, 14, present.ribbonAndBow())
	}
	presents := readPresents()
	var sum int
	for _, present := range presents {
		sum += present.ribbonAndBow()
	}
	require.EqualValues(t, 3812909, sum)
}
