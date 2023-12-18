package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var f field
	var score int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			hScore := mirrorScore(f) * 100
			vScore := mirrorScore(transpose(f))
			score += hScore + vScore
			f = nil
			continue
		}
		f = append(f, []byte(line))
	}
	log.Println(score)
}

type field [][]byte

func transpose(f field) field {
	o := make(field, len(f[0]))
	for i := range o {
		o[i] = make([]byte, len(f))
	}
	for i := range o {
		for j := range o[i] {
			o[i][j] = f[j][i]
		}
	}
	return o
}

func mirrorScore(f field) int {
	var score int
	var possibleMatchingRows []int
	for i := range f {
		if i == 0 {
			continue
		}
		if !bytes.Equal(f[i], f[i-1]) {
			continue
		}
		possibleMatchingRows = append(possibleMatchingRows, i)
	}
nextRow:
	for _, i := range possibleMatchingRows {
		for j := 0; ; j++ {
			l, r := i-(j+1), i+j
			if l < 0 || r >= len(f) {
				break
			}
			if !bytes.Equal(f[l], f[r]) {
				continue nextRow
			}
		}
		return i
	}
	return score
}
