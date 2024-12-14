package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed sample
var in string

func main() {
	var cms []clawmachine
	var cp clawpos
	var a, b func(clawpos) clawpos
	for i, line := range strings.Split(in, "\n") {
		switch i % 4 {
		case 0:
			ix := strings.Index(line, "+")
			ixend := strings.Index(line, ",")
			iy := strings.LastIndex(line, "+")
			x, _ := strconv.Atoi(line[ix+1 : ixend])
			y, _ := strconv.Atoi(line[iy+1:])
			a = func(c clawpos) clawpos { return clawpos{x: c.x + x, y: c.y + y} }
		case 1:
			ix := strings.Index(line, "+")
			ixend := strings.Index(line, ",")
			iy := strings.LastIndex(line, "+")
			x, _ := strconv.Atoi(line[ix+1 : ixend])
			y, _ := strconv.Atoi(line[iy+1:])
			b = func(c clawpos) clawpos { return clawpos{x: c.x + x, y: c.y + y} }
		case 2:
			ix := strings.Index(line, "=")
			ixend := strings.Index(line, ",")
			iy := strings.LastIndex(line, "=")
			x, _ := strconv.Atoi(line[ix+1 : ixend])
			y, _ := strconv.Atoi(line[iy+1:])
			cp = clawpos{x: x, y: y}
		default:
			cms = append(cms, clawmachine{A: a, B: b, P: cp})
			a, b, cp = nil, nil, clawpos{}
		}
	}
	log.Print(cms)
}

type clawpos struct {
	x, y int
}

type clawmachine struct {
	A, B func(clawpos) clawpos
	P    clawpos
}
