package day8

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"github.com/FrankBro/AdventOfCodeGo/util"
)

func find(data string, w, h int) int {
	layerLength := w * h
	count := len(data) / layerLength
	layers := make([]string, count)
	for i := 0; i < count; i++ {
		start := i * layerLength
		end := (i + 1) * layerLength
		layer := data[start:end]
		layers[i] = layer
	}
	var fewest0 int
	min := math.MaxInt64
	for i := 0; i < count; i++ {
		count0 := strings.Count(layers[i], "0")
		if count0 < min {
			fewest0 = i
			min = count0
		}
	}
	return strings.Count(layers[fewest0], "1") * strings.Count(layers[fewest0], "2")
}

func render(data string, w, h int) []byte {
	layerLength := w * h
	count := len(data) / layerLength
	layers := make([]string, count)
	for i := 0; i < count; i++ {
		start := i * layerLength
		end := (i + 1) * layerLength
		layer := data[start:end]
		layers[i] = layer
	}
	render := bytes.Repeat([]byte("2"), layerLength)
	for i := 0; i < count; i++ {
		layer := layers[i]
		for j := 0; j < layerLength; j++ {
			switch {
			case render[j] == '2' && layer[j] == '0':
				render[j] = '0'
			case render[j] == '2' && layer[j] == '1':
				render[j] = '1'
			}
		}
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			bit := render[y*w+x]
			switch bit {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
	return render
}

func getInput() string {
	return util.ReadLine()
}
