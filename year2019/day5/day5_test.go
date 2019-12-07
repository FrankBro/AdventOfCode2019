package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	codes := readInput()
	evalled := eval1(codes, 1)
	require.EqualValues(t, 7692125, evalled)
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, 0, eval2String("3,9,8,9,10,9,4,9,99,-1,8", 7))
	require.EqualValues(t, 1, eval2String("3,9,8,9,10,9,4,9,99,-1,8", 8))
	require.EqualValues(t, 0, eval2String("3,9,8,9,10,9,4,9,99,-1,8", 9))

	require.EqualValues(t, 1, eval2String("3,9,7,9,10,9,4,9,99,-1,8", 7))
	require.EqualValues(t, 0, eval2String("3,9,7,9,10,9,4,9,99,-1,8", 8))
	require.EqualValues(t, 0, eval2String("3,9,7,9,10,9,4,9,99,-1,8", 9))

	require.EqualValues(t, 0, eval2String("3,3,1108,-1,8,3,4,3,99", 7))
	require.EqualValues(t, 1, eval2String("3,3,1108,-1,8,3,4,3,99", 8))
	require.EqualValues(t, 0, eval2String("3,3,1108,-1,8,3,4,3,99", 9))

	require.EqualValues(t, 1, eval2String("3,3,1107,-1,8,3,4,3,99", 7))
	require.EqualValues(t, 0, eval2String("3,3,1107,-1,8,3,4,3,99", 8))
	require.EqualValues(t, 0, eval2String("3,3,1107,-1,8,3,4,3,99", 9))

	require.EqualValues(t, 0, eval2String("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0))
	require.EqualValues(t, 1, eval2String("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1))
	require.EqualValues(t, 1, eval2String("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 2))
	codes := readInput()
	evalled := eval2(codes, 5)
	require.EqualValues(t, 14340395, evalled)
}
