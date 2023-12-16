package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var f field
	for scanner.Scan() {
		line := scanner.Text()
		f = append(f, []byte(line))
		fmt.Println(line)
	}
	fmt.Println()
	tiltNorth(f)
	for i := range f {
		fmt.Println(string(f[i]))
	}
	log.Println(score(f))
}

type field [][]byte

func tiltNorth(f field) {
	for {
		var rolled bool
		for i := range f {
			if i == 0 {
				continue
			}
			for j := range f[i] {
				switch f[i][j] {
				case 'O':
					if f[i-1][j] != '.' {
						continue
					}
					f[i-1][j] = 'O'
					f[i][j] = '.'
					rolled = true
				default:
					continue
				}
			}
		}
		if !rolled {
			return
		}
	}
}

func score(f field) int {
	var score int
	for i := range f {
		rocks := bytes.Count(f[i], []byte{'O'})
		s := rocks * (len(f) - i)
		fmt.Println(rocks, len(f)-i, s)
		score += s
	}
	return score
}
