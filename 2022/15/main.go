package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	yP1 = 2000000
	l   = 4000000
)

type sensor struct {
	sC, bC coord
}

func (s sensor) dist() int { return dist(s.sC, s.bC) }

func (s sensor) perimeter() []coord {
	var p []coord
	p = append(p, coord{x: s.sC.x, y: s.sC.y - s.dist()})
	for i := 1; i <= s.dist(); i++ {
		p = append(p, coord{x: s.sC.x + i, y: s.sC.y - s.dist() + i})
		p = append(p, coord{x: s.sC.x - i, y: s.sC.y - s.dist() + i})
	}
	p = append(p, coord{x: s.sC.x, y: s.sC.y + s.dist()})
	for i := 1; i < s.dist(); i++ {
		p = append(p, coord{x: s.sC.x + i, y: s.sC.y + s.dist() - i})
		p = append(p, coord{x: s.sC.x - i, y: s.sC.y + s.dist() - i})
	}
	return p
}

type coord struct{ x, y int }

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var all []sensor = parseSensors(bufio.NewScanner(f))
	log.Println(len(all[0].perimeter()))
	// part 1
	var sensors []sensor = part1filter(all, yP1)
	fmt.Println("part 1", count(sensors, yP1, nil, nil))
	// part 2
	// tf := func(i, j int) int { return i*l + j }
	lowB := func(i int, s sensor) bool { return s.sC.x-i >= 0 }
	hiB := func(i int, s sensor) bool { return s.sC.x+i <= l }
	for i := 2844840; i <= 2844850; i++ {
		c := count(all, i, lowB, hiB)
		fmt.Println(i)
		fmt.Println("part 2", i, c)
		if c > l {
			continue
		}
		break
	}

	// 11379394658764 low

	var tuneFreq []struct{ i, j, tf int }
	fmt.Println("part 2", tuneFreq)
}

func dist(c1, c2 coord) int {
	return int(math.Abs(float64(c1.x-c2.x)) + math.Abs(float64(c1.y-c2.y)))
}

func count(sensors []sensor, yRef int, lowB, hiB func(int, sensor) bool) int {
	part1 := lowB == nil || hiB == nil
	if lowB == nil {
		lowB = func(int, sensor) bool { return true }
	}
	if hiB == nil {
		hiB = func(int, sensor) bool { return true }
	}
	m := make(map[int]struct{})
	for _, s := range sensors {
		d := s.dist() - int(math.Abs(float64(s.sC.y-yRef)))
		for i := 0; i <= d; i++ {
			if hiB(i, s) {
				m[s.sC.x+i] = struct{}{}
			}
			if lowB(i, s) {
				m[s.sC.x-i] = struct{}{}
			}
		}
	}
	if !part1 {
		return len(m)
	}
	for _, s := range sensors {
		if _, ok := m[s.bC.x]; ok && s.bC.y == yRef {
			delete(m, s.bC.x)
		}
	}
	return len(m)
}

func parseSensors(scanner *bufio.Scanner) []sensor {
	var sensors []sensor
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()

		coordStr := strings.FieldsFunc(line, func(r rune) bool {
			return r != '-' && !unicode.IsDigit(r)
		})
		sx, _ := strconv.Atoi(coordStr[0])
		sy, _ := strconv.Atoi(coordStr[1])
		dx, _ := strconv.Atoi(coordStr[2])
		dy, _ := strconv.Atoi(coordStr[3])

		s := sensor{sC: coord{x: sx, y: sy}, bC: coord{x: dx, y: dy}}
		sensors = append(sensors, s)
	}
	return sensors
}

func part1filter(sensors []sensor, yRef int) []sensor {
	var filtered []sensor
	for _, s := range sensors {
		if s.sC.y < yRef && s.sC.y+s.dist() >= yRef {
			// log.Println("accepting lower: ", s, s.dist(), s.sC.y+s.dist())
			filtered = append(filtered, s)
			continue
		}
		if s.sC.y > yRef && s.sC.y-s.dist() <= yRef {
			// log.Println("accepting higher: ", s, s.dist(), s.sC.y-s.dist())
			filtered = append(filtered, s)
			continue
		}
		// log.Println("skipping: ", s, s.dist(), int(math.Abs(float64(s.sC.y-s.dist()))))
	}
	return filtered
}
