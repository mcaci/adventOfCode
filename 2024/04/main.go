package main

import (
	"bytes"
	_ "embed"
	"log"
)

//go:embed input
var in []byte

func main() {
	mat := bytes.Split(in, []byte("\n"))
	xmas1 := xmasSingleDir(mat)
	xmas2 := xmasAnydir(mat)
	xmas3 := x_masSingleDir(mat)
	// 3358 too high
	log.Println(xmas1)
	log.Println(xmas2)
	// 1749 too high
	log.Println(xmas3)
}

type XY struct{ x, y int }
type XMAS [4]XY
type X_MAS [3]XY

func xmasLine(mat [][]byte, a, b int) []XMAS {
	xmasRight := XMAS{XY{a, b}, XY{a, b + 1}, XY{a, b + 2}, XY{a, b + 3}}
	xmasLeft := XMAS{XY{a, b}, XY{a, b - 1}, XY{a, b - 2}, XY{a, b - 3}}
	xmasDown := XMAS{XY{a, b}, XY{a + 1, b}, XY{a + 2, b}, XY{a + 3, b}}
	xmasUp := XMAS{XY{a, b}, XY{a - 1, b}, XY{a - 2, b}, XY{a - 3, b}}
	xmasDownRight := XMAS{XY{a, b}, XY{a + 1, b + 1}, XY{a + 2, b + 2}, XY{a + 3, b + 3}}
	xmasDownLeft := XMAS{XY{a, b}, XY{a + 1, b - 1}, XY{a + 2, b - 2}, XY{a + 3, b - 3}}
	xmasUpRight := XMAS{XY{a, b}, XY{a - 1, b + 1}, XY{a - 2, b + 2}, XY{a - 3, b + 3}}
	xmasUpLeft := XMAS{XY{a, b}, XY{a - 1, b - 1}, XY{a - 2, b - 2}, XY{a - 3, b - 3}}
	xmases := []XMAS{xmasRight, xmasLeft, xmasDown, xmasUp, xmasDownRight, xmasDownLeft, xmasUpRight, xmasUpLeft}
	// check if xmases points are in bounds and contain the right characters in the right order
	var xmas []XMAS
	for _, x := range xmases {
		if x[0].x < 0 || x[0].x >= len(mat) {
			continue
		}
		if x[0].y < 0 || x[0].y >= len(mat[x[0].x]) {
			continue
		}
		if x[1].x < 0 || x[1].x >= len(mat) {
			continue
		}
		if x[1].y < 0 || x[1].y >= len(mat[x[1].x]) {
			continue
		}
		if x[2].x < 0 || x[2].x >= len(mat) {
			continue
		}
		if x[2].y < 0 || x[2].y >= len(mat[x[2].x]) {
			continue
		}
		if x[3].x < 0 || x[3].x >= len(mat) {
			continue
		}
		if x[3].y < 0 || x[3].y >= len(mat[x[3].x]) {
			continue
		}
		if mat[x[0].x][x[0].y] != 'X' {
			continue
		}
		if mat[x[1].x][x[1].y] != 'M' {
			continue
		}
		if mat[x[2].x][x[2].y] != 'A' {
			continue
		}
		if mat[x[3].x][x[3].y] != 'S' {
			continue
		}
		xmas = append(xmas, x)
	}
	return xmas
}

func x_masLine(mat [][]byte, a, b int) int {
	xmasDownRight := X_MAS{XY{a - 1, b - 1}, XY{a, b}, XY{a + 1, b + 1}}
	xmasDownLeft := X_MAS{XY{a - 1, b + 1}, XY{a, b}, XY{a + 1, b - 1}}
	xmasUpRight := X_MAS{XY{a + 1, b - 1}, XY{a, b}, XY{a - 1, b + 1}}
	xmasUpLeft := X_MAS{XY{a + 1, b + 1}, XY{a, b}, XY{a - 1, b - 1}}
	// check if xmases points are in bounds and contain the right characters in the right order
	var count int
	diagonalXmases := []X_MAS{xmasDownRight, xmasDownLeft, xmasUpRight, xmasUpLeft}
	for _, x := range diagonalXmases {
		if x[0].x < 0 || x[0].x >= len(mat) {
			continue
		}
		if x[0].y < 0 || x[0].y >= len(mat[x[0].x]) {
			continue
		}
		if x[1].x < 0 || x[1].x >= len(mat) {
			continue
		}
		if x[1].y < 0 || x[1].y >= len(mat[x[1].x]) {
			continue
		}
		if x[2].x < 0 || x[2].x >= len(mat) {
			continue
		}
		if x[2].y < 0 || x[2].y >= len(mat[x[2].x]) {
			continue
		}
		if mat[x[0].x][x[0].y] != 'M' {
			continue
		}
		if mat[x[1].x][x[1].y] != 'A' {
			continue
		}
		if mat[x[2].x][x[2].y] != 'S' {
			continue
		}
		count++
	}
	count /= 2
	return count
}

func nextCharacter(mat [][]byte, char byte, a, b int) []XY {
	var chars []XY
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if a+i < 0 || a+i >= len(mat) {
				continue
			}
			if b+j < 0 || b+j >= len(mat[a+i]) {
				continue
			}
			if mat[a+i][b+j] != char {
				continue
			}
			chars = append(chars, XY{a + i, b + j})
		}
	}
	return chars
}

func xmasSingleDir(mat [][]byte) int {
	xmas := make(map[XMAS]bool)
	for i := range mat {
		for j := range mat[i] {
			// look for 'x'
			if mat[i][j] != 'X' {
				continue
			}
			xmases := xmasLine(mat, i, j)
			for _, x := range xmases {
				xmas[x] = true
			}
		}
	}
	return len(xmas)
}

func xmasAnydir(mat [][]byte) int {
	xmas := make(map[XMAS]bool)
	for i := range mat {
		for j := range mat[i] {
			// look for 'x'
			if mat[i][j] != 'X' {
				continue
			}
			x := XY{i, j}
			// look around 'x' for 'm'
			ms := nextCharacter(mat, 'M', x.x, x.y)
			if ms == nil {
				continue
			}
			for _, m := range ms {
				// look around 'm' for 'a'
				as := nextCharacter(mat, 'A', m.x, m.y)
				if as == nil {
					continue
				}
				for _, a := range as {
					// look around 'a' for 's'
					ss := nextCharacter(mat, 'S', a.x, a.y)
					if ss == nil {
						continue
					}
					for _, s := range ss {
						xmas[XMAS{
							XY{x.x, x.y},
							XY{m.x, m.y},
							XY{a.x, a.y},
							XY{s.x, s.y},
						}] = true
					}
				}
			}

		}
	}
	return len(xmas)
}

func x_masSingleDir(mat [][]byte) int {
	var count int
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] != 'A' {
				continue
			}
			count += x_masLine(mat, i, j)
		}
	}
	return count
}
