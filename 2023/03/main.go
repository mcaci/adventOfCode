package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	b, _ := os.ReadFile("input")
	var x, y int
	var schematics []schematicByte
	for i := range b {
		switch b[i] {
		case '.':
			x++
		case '\r':
			// for windows only
		case '\n':
			x = 0
			y++
		default:
			schematics = append(schematics, schematicByte{x: x, y: y, b: b[i]})
			x++
		}
	}
	fmt.Println(part1sum(sBytesToSymbols(schematics), sToNoDigitBytes(schematics)))
	fmt.Println(part2sum(sBytesToSymbols(schematics), sToGearBytes(schematics)))
}

type gear struct{ a, b schematicSymbol }

func (g gear) String() string { return fmt.Sprintf("%v*%v", g.a, g.b) }

func sBytesToGears(sss []schematicSymbol, gbytes []schematicByte) []gear {
	// fmt.Println(sss, gbytes)
	var gears []gear
	for i1, s1 := range sss {
		g1, ok := adjacent(s1, gbytes)
		if !ok {
			continue
		}
		for _, s2 := range sss[i1+1:] {
			g2, ok := adjacent(s2, gbytes)
			if !ok {
				continue
			}
			if g1 != g2 {
				continue
			}
			gears = append(gears, gear{a: s1, b: s2})
			break
		}
	}
	// fmt.Println(gears)
	return gears
}

func part2sum(sss []schematicSymbol, gbytes []schematicByte) int {
	var sum int
	g := sBytesToGears(sss, gbytes)
	for i := range g {
		sum += g[i].a.n() * g[i].b.n()
	}
	return sum
}

func part1sum(sss []schematicSymbol, ndbytes []schematicByte) int {
	var sum int
	for i := range sss {
		if _, ok := adjacent(sss[i], ndbytes); !ok {
			continue
		}
		sum += sss[i].n()
	}
	return sum
}

func adjacent(s schematicSymbol, ndbytes []schematicByte) (schematicByte, bool) {
	minX, maxX := s[0].x-1, s[len(s)-1].x+1
	minY, maxY := s[0].y-1, s[0].y+1
	for _, nd := range ndbytes {
		if nd.y < minY {
			continue
		}
		if nd.y > maxY {
			continue
		}
		if nd.x < minX {
			continue
		}
		if nd.x > maxX {
			continue
		}
		return nd, true
	}
	return schematicByte{}, false
}

type schematicByte struct {
	x, y int
	b    byte
}

func (s schematicByte) String() string {
	return fmt.Sprintf("%d:%d n:%s.", s.x, s.y, string(s.b))
}

type schematicSymbol []schematicByte

func (s schematicSymbol) n() int {
	var b []byte
	for i := range s {
		b = append(b, s[i].b)
	}
	n, _ := strconv.Atoi(string(b))
	return n
}

func (s schematicSymbol) String() string {
	var b []byte
	for i := range s {
		b = append(b, s[i].b)
	}
	return fmt.Sprintf("<%d:%d>:%d n:%s.", s[0].x, s[len(s)-1].x, s[0].y, string(b))
}

func sBytesToSymbols(sbs []schematicByte) []schematicSymbol {
	var sss []schematicSymbol
nextSymbol:
	for i := 0; i < len(sbs); {
		if !unicode.IsDigit(rune(sbs[i].b)) {
			i++
			continue
		}
		symbol := schematicSymbol{sbs[i]}
		for i < len(sbs) {
			if i+1 >= len(sbs) {
				sss = append(sss, symbol)
				break nextSymbol
			}
			if !unicode.IsDigit(rune(sbs[i+1].b)) {
				sss = append(sss, symbol)
				i++
				continue nextSymbol
			}

			switch sbs[i+1].x - sbs[i].x {
			case 1:
				symbol = append(symbol, sbs[i+1])
				i++
				continue
			default:
				sss = append(sss, symbol)
				i++
				continue nextSymbol
			}
		}
	}
	return sss
}

func sToNoDigitBytes(sbs []schematicByte) []schematicByte {
	var out []schematicByte
	for i := range sbs {
		if unicode.IsDigit(rune(sbs[i].b)) {
			continue
		}
		out = append(out, sbs[i])
	}
	return out
}

func sToGearBytes(sbs []schematicByte) []schematicByte {
	var out []schematicByte
	for i := range sbs {
		if sbs[i].b != '*' {
			continue
		}
		out = append(out, sbs[i])
	}
	return out
}
