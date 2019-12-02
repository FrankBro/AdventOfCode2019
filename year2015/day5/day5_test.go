package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.EqualValues(t, true, isNice("ugknbfddgicrmopn"))
	require.EqualValues(t, true, isNice("aaa"))
	require.EqualValues(t, false, isNice("jchzalrnumimnmhp"))
	require.EqualValues(t, false, isNice("haegwjzuvuyypxyu"))
	require.EqualValues(t, false, isNice("dvszwmarrgswjxmb"))
	xs := readStrings()
	var sum int
	for _, x := range xs {
		if isNice(x) {
			sum++
		}
	}
	require.EqualValues(t, 236, sum)
}

func TestPart2(t *testing.T) {
	require.EqualValues(t, true, isNice2("qjhvhtzxzqqjkmpb"))
	require.EqualValues(t, true, isNice2("xxyxx"))
	require.EqualValues(t, false, isNice2("uurcxstgmygtbstg"))
	require.EqualValues(t, false, isNice2("ieodomkazucvgmuy"))
	xs := readStrings()
	var sum int
	for _, x := range xs {
		if isNice2(x) {
			sum++
		}
	}
	require.EqualValues(t, 51, sum)
}
