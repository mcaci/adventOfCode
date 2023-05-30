package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dirs := []dir{north, south, west, east}

	elves := parseInput(f)
	// plot(elves)

	n := 0
	for {
		n++
		for i := range elves {
			scan(elves[i], elves, dirs)
		}
		dest := make(map[struct{ i, j int }]int)
		for _, e := range elves {
			if !e.moving {
				continue
			}
			dest[destination(e)]++
		}
		if len(dest) == 0 {
			break
		}
		for i := range elves {
			e := elves[i]
			if !e.moving {
				continue
			}
			e.moving = false
			if dest[destination(e)] > 1 {
				e.iNext = e.i
				e.jNext = e.j
				continue
			}
			e.i = e.iNext
			e.j = e.jNext
		}
		dirs = append(dirs[1:], dirs[0])
		// fmt.Printf("TURN %d\n", n)
		// plot(elves)
	}
	var minJ, maxJ int = math.MaxInt, math.MinInt
	var minI, maxI int = math.MaxInt, math.MinInt
	for i := range elves {
		e := elves[i]
		if e.i > maxI {
			maxI = e.i
		}
		if e.i < minI {
			minI = e.i
		}
		if e.j < minJ {
			minJ = e.j
		}
		if e.j > maxJ {
			maxJ = e.j
		}
	}
	// part 1
	fmt.Println("part 1:", ((maxI-minI+1)*(maxJ-minJ+1))-len(elves))

	// part 2
	fmt.Println("part 2:", n)
}

type elf struct {
	i, j         int
	iNext, jNext int
	d            dir
	moving       bool
}

func origin(e *elf) struct{ i, j int }      { return struct{ i, j int }{i: e.i, j: e.j} }
func destination(e *elf) struct{ i, j int } { return struct{ i, j int }{i: e.iNext, j: e.jNext} }

func scan(e *elf, elves []*elf, dirs []dir) {
	around := neighbours(e, elves)
	if len(around) == 0 {
		return
	}
freePlaceSearch:
	for _, d := range dirs {
		switch d {
		case north:
			if freeSpace(e, around, []elf{{i: e.i - 1, j: e.j - 1}, {i: e.i - 1, j: e.j}, {i: e.i - 1, j: e.j + 1}}) {
				e.d = north
				e.iNext = e.i - 1
				e.jNext = e.j
				e.moving = true
				break freePlaceSearch
			}
		case south:
			if freeSpace(e, around, []elf{{i: e.i + 1, j: e.j - 1}, {i: e.i + 1, j: e.j}, {i: e.i + 1, j: e.j + 1}}) {
				e.d = south
				e.iNext = e.i + 1
				e.jNext = e.j
				e.moving = true
				break freePlaceSearch
			}
		case west:
			if freeSpace(e, around, []elf{{i: e.i - 1, j: e.j - 1}, {i: e.i, j: e.j - 1}, {i: e.i + 1, j: e.j - 1}}) {
				e.d = west
				e.iNext = e.i
				e.jNext = e.j - 1
				e.moving = true
				break freePlaceSearch
			}
		case east:
			if freeSpace(e, around, []elf{{i: e.i - 1, j: e.j + 1}, {i: e.i, j: e.j + 1}, {i: e.i + 1, j: e.j + 1}}) {
				e.d = east
				e.iNext = e.i
				e.jNext = e.j + 1
				e.moving = true
				break freePlaceSearch
			}
		}
	}
}

func freeSpace(e *elf, elves []*elf, neighbours []elf) bool {
	for i := range elves {
		for j := range neighbours {
			if origin(elves[i]) != origin(&neighbours[j]) {
				continue
			}
			return false
		}
	}
	return true
}

func neighbours(e *elf, elves []*elf) []*elf {
	neighbouringPlaces := []struct{ i, j int }{{i: e.i - 1, j: e.j}, {i: e.i - 1, j: e.j - 1}, {i: e.i, j: e.j - 1}, {i: e.i + 1, j: e.j - 1}, {i: e.i + 1, j: e.j}, {i: e.i + 1, j: e.j + 1}, {i: e.i, j: e.j + 1}, {i: e.i - 1, j: e.j + 1}}
	var neighbouringElves []*elf
	for i := range elves {
		for j := range neighbouringPlaces {
			if origin(elves[i]) != neighbouringPlaces[j] {
				continue
			}
			neighbouringElves = append(neighbouringElves, elves[i])
		}
	}
	return neighbouringElves
}

type dir uint8

const (
	north dir = iota
	south
	west
	east
)

func parseInput(r io.Reader) []*elf {
	scanner := bufio.NewScanner(r)
	var elves []*elf
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			if c != '#' {
				continue
			}
			elves = append(elves, &elf{i: i, j: j})
		}
	}
	return elves
}

func plot(elves []*elf) {
	var minJ, maxJ int = math.MaxInt, math.MinInt
	var minI, maxI int = math.MaxInt, math.MinInt
	for i := range elves {
		e := elves[i]
		if e.i > maxI {
			maxI = e.i
		}
		if e.i < minI {
			minI = e.i
		}
		if e.j < minJ {
			minJ = e.j
		}
		if e.j > maxJ {
			maxJ = e.j
		}
	}
	maxI++
	maxJ++
	m := make([][]byte, maxI-minI)
	for i := range m {
		m[i] = make([]byte, maxJ-minJ)
		for j := range m[i] {
			m[i][j] = '.'
		}
	}
	for _, e := range elves {
		m[e.i-minI][e.j-minJ] = '#'
	}
	for i := range m {
		fmt.Println(string(m[i]))
	}
}
