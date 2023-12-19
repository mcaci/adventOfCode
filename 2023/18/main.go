package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

//go:embed sample
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var instrs []instruction
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Fields(line)
		n, _ := strconv.Atoi(f[1])
		instr := instruction{
			d:      newLRUD(f[0]),
			nSteps: n,
			c:      newColor(f[2]),
		}
		instrs = append(instrs, instr)
	}
	fmt.Print(instrs)
}

type instruction struct {
	d      direction
	nSteps int
	c      color.Color
}

func (i instruction) String() string {
	return fmt.Sprintf("{%v %d %v}", i.d, i.nSteps, i.c)
}

type field [][]byte

func (f field) String() string {
	var s strings.Builder
	for _, l := range f {
		s.Write(l)
		s.WriteByte('\n')
	}
	return s.String()
}

type vect struct {
	x, y int
	dir  direction
}

func (v vect) String() string { return fmt.Sprintf("{%d %d %v}", v.x, v.y, v.dir) }

type direction byte

const (
	N = direction(iota)
	E
	S
	W
)

func newLRUD(s string) direction {
	switch s {
	case "L":
		return W
	case "R":
		return E
	case "U":
		return N
	case "D":
		return S
	default:
		return direction(s[0])
	}
}

func (d direction) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	default:
		return "#"
	}
}

func newColor(s string) color.Color {
	c := color.RGBA{A: 255}
	r, _ := strconv.ParseInt(s[2:4], 16, 8)
	c.R = uint8(r)
	g, _ := strconv.ParseInt(s[4:6], 16, 8)
	c.G = uint8(g)
	b, _ := strconv.ParseInt(s[6:8], 16, 8)
	c.B = uint8(b)
	return c
}
