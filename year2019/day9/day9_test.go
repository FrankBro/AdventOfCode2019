package day9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", run("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"))
	require.EqualValues(t, 16, len(run("1102,34915192,34915192,7,4,7,99,0")))
	require.EqualValues(t, "1125899906842624", run("104,1125899906842624,99"))
	input := getInput()
	require.EqualValues(t, "3638931938", run(input, 1))
}

func TestPart2(t *testing.T) {
	input := getInput()
	require.EqualValues(t, "86025", run(input, 2))
}
