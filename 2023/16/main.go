package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
	"sync"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var f, e field
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, []byte(line))
		e = append(e, make([]byte, len(line)))
	}
	var wg sync.WaitGroup
	var starts []vect
	for i := 0; i < len(f); i++ {
		starts = append(starts, vect{x: 0, y: i, dir: S})
		starts = append(starts, vect{x: i, y: 0, dir: E})
		starts = append(starts, vect{x: i, y: len(f) - 1, dir: W})
		starts = append(starts, vect{x: len(f) - 1, y: i, dir: N})
	}
	maxC := make(chan int, len(starts))
	wg.Add(len(starts))
	for _, s := range starts {
		go func(s vect) {
			defer wg.Done()
			e1 := make(field, len(e))
			for i := range e {
				e1[i] = make([]byte, len(e[i]))
				copy(e1[i], e[i])
			}
			cover([]vect{s}, f, e1)
			scor := score(e1)
			log.Println(s, scor)
			maxC <- scor
		}(s)
	}
	wg.Wait()
	close(maxC)
	var max int
	for m := range maxC {
		if max > m {
			continue
		}
		max = m
	}
	log.Println(max)
}

type vect struct {
	x, y int
	dir  direction
}

func (v vect) String() string { return fmt.Sprintf("{%d %d %v}", v.x, v.y, v.dir) }

type direction byte

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

const (
	N = direction(iota)
	E
	S
	W
)

type field [][]byte

func cover(vs []vect, f, e field) {
	lastChanged := []bool{true, true, true, true, true, true, true, true, true, true, true}
	for slices.Contains(lastChanged, true) {
		var ovs []vect
		var changed bool
		for _, v := range vs {
			if e[v.x][v.y] == 0 {
				e[v.x][v.y] = 1
				changed = true
			}
			a, b, bOk := meet(v, f[v.x][v.y])
			a = move(a)
			if (a.x >= 0 && a.y >= 0) && (a.x < len(f[0]) && a.y < len(f)) {
				ovs = append(ovs, a)
			}
			if bOk {
				b = move(b)
				if (b.x >= 0 && b.y >= 0) && (b.x < len(f[0]) && b.y < len(f)) {
					ovs = append(ovs, b)
				}
			}
		}
		vs = ovs
		if !changed {
			lastChanged = append(lastChanged[1:], false)
			continue
		}
		lastChanged = append(lastChanged[1:], true)
	}
}

func meet(v vect, cell byte) (vect, vect, bool) {
	switch cell {
	case '\\':
		o := v
		switch v.dir {
		case N:
			o.dir = W
		case S:
			o.dir = E
		case E:
			o.dir = S
		case W:
			o.dir = N
		}
		return o, vect{}, false
	case '/':
		o := v
		switch v.dir {
		case N:
			o.dir = E
		case S:
			o.dir = W
		case E:
			o.dir = N
		case W:
			o.dir = S
		}
		return o, vect{}, false
	case '-':
		switch v.dir {
		case N:
			o1, o2 := v, v
			o1.dir = E
			o2.dir = W
			return o1, o2, true
		case S:
			o1, o2 := v, v
			o1.dir = E
			o2.dir = W
			return o1, o2, true
		case W:
			return v, vect{}, false
		case E:
			return v, vect{}, false
		}
	case '|':
		switch v.dir {
		case N:
			return v, vect{}, false
		case S:
			return v, vect{}, false
		case W:
			o1, o2 := v, v
			o1.dir = N
			o2.dir = S
			return o1, o2, true
		case E:
			o1, o2 := v, v
			o1.dir = N
			o2.dir = S
			return o1, o2, true
		}
	default:
		// case '.':
		return v, vect{}, false
	}
	log.Fatal("nope")
	return vect{}, vect{}, false
}

func move(v vect) vect {
	switch v.dir {
	case W:
		return vect{x: v.x, y: v.y - 1, dir: v.dir}
	case E:
		return vect{x: v.x, y: v.y + 1, dir: v.dir}
	case N:
		return vect{x: v.x - 1, y: v.y, dir: v.dir}
	case S:
		return vect{x: v.x + 1, y: v.y, dir: v.dir}
	default:
		log.Fatal("nope")
		return vect{}
	}
}

func score(f field) int {
	return bytes.Count(bytes.Join(f, []byte{}), []byte{1})
}
