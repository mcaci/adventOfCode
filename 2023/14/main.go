package main

import (
	"bufio"
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
	var score int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			score += mirrorScore(f, 100)
			score += mirrorScore(transpose(f), 1)
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

func mirrorScore(f field, tFactor int) int {
	j := len(f) - 1
	for i := 0; i < len(f); i++ {
		if string(f[i]) != string(f[j]) {
			continue
		}
		if j-i == 1 {
			fmt.Println(i, j, tFactor, string(f[i]), string(f[j]))
			return j * tFactor
		}
		j--
	}
	j = 0
	for i := len(f) - 1; i >= 0; i-- {
		if string(f[i]) != string(f[j]) {
			continue
		}
		if i-j == 1 {
			fmt.Println(i, j, tFactor, string(f[i]), string(f[j]))
			return i * tFactor
		}
		j++
	}
	return 0
}
