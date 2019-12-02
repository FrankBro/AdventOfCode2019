package day2

import (
	"strconv"
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func surface(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

type present struct {
	l int
	w int
	h int
}

func (present *present) surface() int {
	return surface(present.l, present.w, present.h)
}

func (present *present) smallestSide() int {
	lw := present.l * present.w
	wh := present.w * present.h
	hl := present.h * present.l
	min := lw
	if wh < min {
		min = wh
	}
	if hl < min {
		min = hl
	}
	return min
}

func (present *present) wrappingPaper() int {
	return present.surface() + present.smallestSide()
}

func (present *present) ribbon() int {
	min1 := present.l
	min2 := present.w
	if min2 < min1 {
		temp := min1
		min1 = min2
		min2 = temp
	}
	if present.h < min2 {
		min2 = present.h
	}
	return min1*2 + min2*2
}

func (present *present) bow() int {
	return present.l * present.w * present.h
}

func (present *present) ribbonAndBow() int {
	return present.ribbon() + present.bow()
}

func readPresent(line string) present {
	dimensions := strings.Split(line, "x")
	l, err := strconv.Atoi(dimensions[0])
	util.Check(err)
	w, err := strconv.Atoi(dimensions[1])
	util.Check(err)
	h, err := strconv.Atoi(dimensions[2])
	util.Check(err)
	return present{l: l, w: w, h: h}
}

func readPresents() []present {
	lines := util.ReadLines()
	presents := make([]present, len(lines))
	for i, line := range lines {
		presents[i] = readPresent(line)
	}
	return presents
}
