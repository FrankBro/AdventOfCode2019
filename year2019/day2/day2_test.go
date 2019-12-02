package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, "2,0,0,0,99", evalString("1,0,0,0,99"))
	require.EqualValues(t, "2,3,0,6,99", evalString("2,3,0,3,99"))
	require.EqualValues(t, "2,4,4,5,99,9801", evalString("2,4,4,5,99,0"))
	require.EqualValues(t, "30,1,1,4,2,5,6,0,99", evalString("1,1,1,4,99,5,6,0,99"))
	codes := readInput()
	codes[1] = 12
	codes[2] = 2
	evalled := eval(codes)
	require.EqualValues(t, 4945026, evalled[0])
}

func TestPart2(t *testing.T) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			codes := readInput()
			codes[1] = noun
			codes[2] = verb
			evalled := eval(codes)
			if evalled[0] == 19690720 {
				require.EqualValues(t, 5296, 100*noun+verb)
				return
			}
		}
	}
}
