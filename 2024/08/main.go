package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed sample
var in string

func main() {
	var g grid
	var l int
	for i, line := range strings.Split(in, "\n") {
		l++
		for j, l := range line {
			switch l {
			case '.':
				continue
			default:
				g.antennas = append(g.antennas, antenna{letter: l, x: j, y: i})
			}
		}
	}
	g.w = l
	g.h = l
	log.Print(g)
}

type grid struct {
	h, w     int
	antennas []antenna
}

type antenna struct {
	letter rune
	x, y   int
}
