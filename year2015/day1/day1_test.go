package day1

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, 0, findFloor("(())"))
	require.EqualValues(t, 0, findFloor("()()"))
	require.EqualValues(t, 3, findFloor("((("))
	require.EqualValues(t, 3, findFloor("(()(()("))
	require.EqualValues(t, 3, findFloor("))((((("))
	require.EqualValues(t, -1, findFloor("())"))
	require.EqualValues(t, -1, findFloor("))("))
	require.EqualValues(t, -3, findFloor(")))"))
	require.EqualValues(t, -3, findFloor(")())())"))
	file, err := os.Open("input")
	require.NoError(t, err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	require.EqualValues(t, 232, findFloor(line))
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, 1, findPosition(")"))
	require.EqualValues(t, 5, findPosition("()())"))
	file, err := os.Open("input")
	require.NoError(t, err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	require.EqualValues(t, 1783, findPosition(line))
}
