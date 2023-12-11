package main

import (
	"bufio"
	_ "embed"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var constellations []xy
	var rows int
	for scanner.Scan() {
		line := scanner.Text()
		for x := range line {
			if line[x] != '#' {
				continue
			}
			constellations = append(constellations, xy{x: x, y: rows})
		}
		rows++
	}
	empty := make(map[int]byte)
	for i := 0; i < rows; i++ {
		if slices.ContainsFunc(constellations, func(o xy) bool { return o.y == i }) {
			continue
		}
		empty[i] = 'R'
	}
	columns := strings.Index(input, "\n") - 1
	for i := 0; i < columns; i++ {
		if slices.ContainsFunc(constellations, func(o xy) bool { return o.x == i }) {
			continue
		}
		_, ok := empty[i]
		if !ok {
			empty[i] = 'C'
			continue
		}
		empty[i] = 'A'
	}

	var sortedKeys []int
	for k := range empty {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	for i := range constellations {
		var dx, dy int
		for _, k := range sortedKeys {
			v := empty[k]
			switch v {
			case 'A':
				if k < constellations[i].x {
					// dx += 1 // part1
					dx += 1000000 - 1 // part2
				}
				if k < constellations[i].y {
					// dy += 1 // part1
					dy += 1000000 - 1 // part2
				}
			case 'C':
				if k < constellations[i].x {
					// dx += 1 // part1
					dx += 1000000 - 1 // part2
				}
			case 'R':
				if k < constellations[i].y {
					// dy += 1 // part1
					dy += 1000000 - 1 // part2
				}
			default:
				log.Fatal(k, v, "nope")
			}
		}
		constellations[i].x += dx
		constellations[i].y += dy
	}

	var sum int
	for i, c1 := range constellations {
		for _, c2 := range constellations[i+1:] {
			sum += dist(c1, c2)
		}
	}
	log.Println(sum)
}

type xy struct{ x, y int }

func dist(a, b xy) int {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return int(math.Abs(dx) + math.Abs(dy))
}
