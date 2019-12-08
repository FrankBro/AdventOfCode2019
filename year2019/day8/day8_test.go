package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, 1, find("123456789012", 3, 2))
	input := getInput()
	require.EqualValues(t, 2760, find(input, 25, 6))
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, []byte("0110"), render("0222112222120000", 2, 2))
	input := getInput()
	require.EqualValues(t, []byte("011000110010010111101110010010100101001010000100101001010000100101110011100111101011010010100001001010010100101001010000100101001001110011001111011100"), render(input, 25, 6))

}
