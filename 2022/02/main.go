package main

import (
	"bytes"
	"fmt"
	"os"
)

type rps int

const (
	rock rps = iota + 1
	paper
	scissors
)

type score int

const (
	w score = 6
	d score = 3
	l score = 0
)

func match(p1, p2 rps) score {
	switch p1 {
	case rock:
		switch p2 {
		case rock:
			return d
		case paper:
			return l
		case scissors:
			return w
		}
	case paper:
		switch p2 {
		case rock:
			return w
		case paper:
			return d
		case scissors:
			return l
		}
	case scissors:
		switch p2 {
		case rock:
			return l
		case paper:
			return w
		case scissors:
			return d
		}
	}
	return 0
}

func revMatch(opp rps, s score) rps {
	switch opp {
	case rock:
		switch s {
		case l:
			return scissors
		case d:
			return rock
		case w:
			return paper
		}
	case paper:
		switch s {
		case l:
			return rock
		case d:
			return paper
		case w:
			return scissors
		}
	case scissors:
		switch s {
		case l:
			return paper
		case d:
			return scissors
		case w:
			return rock
		}
	}
	return 0
}

var (
	p1Rps = map[byte]rps{'A': rock, 'B': paper, 'C': scissors}
	// p2Rps    = map[byte]rps{'X': rock, 'Y': paper, 'Z': scissors}
	scoreRps = map[byte]score{'X': l, 'Y': d, 'Z': w}
)

func main() {
	b, _ := os.ReadFile("input")
	l := bytes.Split(b, []byte{'\n'})
	var score int
	for i := range l {
		opp := p1Rps[l[i][0]]
		// me := p2Rps[l[i][2]]
		scr := scoreRps[l[i][2]]
		score += int(scr)
		// score += int(match(me, opp))
		score += int(revMatch(opp, scr))
	}
	fmt.Println(score)
}
