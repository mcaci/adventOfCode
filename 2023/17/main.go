package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var f field
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, []byte(line))
	}
	fmt.Print(f)
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
