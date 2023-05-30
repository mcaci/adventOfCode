package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct{ x, y int }

func closeEnough(h, t coord) bool {
	for x := t.x - 1; x <= t.x+1; x++ {
		for y := t.y - 1; y <= t.y+1; y++ {
			c := coord{x: x, y: y}
			if c != h {
				continue
			}
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var posH, pos1, pos2, pos3, pos4, pos5, pos6, pos7, pos8, posT coord
	var pathP1 []coord = []coord{pos1}
	var pathP2 []coord = []coord{posT}
	scanner := bufio.NewScanner(f)
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()
		fields := strings.Fields(line)
		n, _ := strconv.Atoi(fields[1])
		for i := 0; i < n; i++ {
			posH = moveH(posH, fields[0])
			if !closeEnough(posH, pos1) {
				pos1 = moveT(posH, pos1)
				pathP1 = append(pathP1, pos1)
			}
			if !closeEnough(pos1, pos2) {
				pos2 = moveT(pos1, pos2)
			}
			if !closeEnough(pos2, pos3) {
				pos3 = moveT(pos2, pos3)
			}
			if !closeEnough(pos3, pos4) {
				pos4 = moveT(pos3, pos4)
			}
			if !closeEnough(pos4, pos5) {
				pos5 = moveT(pos4, pos5)
			}
			if !closeEnough(pos5, pos6) {
				pos6 = moveT(pos5, pos6)
			}
			if !closeEnough(pos6, pos7) {
				pos7 = moveT(pos6, pos7)
			}
			if !closeEnough(pos7, pos8) {
				pos8 = moveT(pos7, pos8)
			}
			if !closeEnough(pos8, posT) {
				posT = moveT(pos8, posT)
				pathP2 = append(pathP2, posT)
			}
		}
	}
	m := make(map[coord]struct{})
	for i := range pathP1 {
		m[pathP1[i]] = struct{}{}
	}
	fmt.Println("part 1", len(m), len(pathP1))
	// part 2
	m2 := make(map[coord]struct{})
	for i := range pathP2 {
		m2[pathP2[i]] = struct{}{}
	}
	fmt.Println("part 2", len(m2), len(pathP2))

}

func moveH(c coord, d string) coord {
	switch d {
	case "U":
		return coord{x: c.x, y: c.y + 1}
	case "L":
		return coord{x: c.x - 1, y: c.y}
	case "D":
		return coord{x: c.x, y: c.y - 1}
	case "R":
		return coord{x: c.x + 1, y: c.y}
	}
	return c
}

func moveT(h, t coord) coord {
	if h.x == t.x {
		if h.y > t.y {
			return coord{x: t.x, y: t.y + 1}
		}
		return coord{x: t.x, y: t.y - 1}
	}
	if h.y == t.y {
		if h.x > t.x {
			return coord{x: t.x + 1, y: t.y}
		}
		return coord{x: t.x - 1, y: t.y}
	}
	if h.x > t.x {
		if h.y > t.y {
			return coord{x: t.x + 1, y: t.y + 1}
		}
		return coord{x: t.x + 1, y: t.y - 1}
	}
	if h.x < t.x {
		if h.y > t.y {
			return coord{x: t.x - 1, y: t.y + 1}
		}
		return coord{x: t.x - 1, y: t.y - 1}
	}
	return t
}
