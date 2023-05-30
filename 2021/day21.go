package main

import "fmt"

// Day21 Input
// Player 1 starting position: 7
// Player 2 starting position: 3

const (
	LEN_BOARD = 10
	P1_POS    = 7
	P2_POS    = 3
)

func main() {
	day21Part1()
	day21Part2()
}

type player struct {
	pos, score int
}

func (p *player) action(d interface{ roll() int }) int {
	d1, d2, d3 := d.roll(), d.roll(), d.roll()
	p.pos = (p.pos + d1 + d2 + d3) % LEN_BOARD
	if p.pos == 0 {
		p.pos = 10
	}
	return p.pos
}

type deterministicDie int

func (d *deterministicDie) roll() int {
	v := int(*d)
	(*d)++
	if *d > 100 {
		*d = 1
	}
	return v
}

func day21Part1() {
	p1 := player{pos: P1_POS}
	p2 := player{pos: P2_POS}
	d := deterministicDie(1)
	var countRolls int
	for p1.score < 1000 || p2.score < 1000 {
		p1.score += p1.action(&d)
		countRolls += 3
		if p1.score >= 1000 {
			break
		}
		p2.score += p2.action(&d)
		countRolls += 3
		if p2.score >= 1000 {
			break
		}
	}
	fmt.Println(countRolls, p1, p2, countRolls*p1.score, countRolls*p2.score)
}

type State struct {
	p1, p2           int
	p1score, p2score int
}

func Advance(pos int, roll int) int {
	pos += roll
	for pos > 10 {
		pos -= 10
	}
	return pos
}

func day21Part2() {
	states := map[State]int64{{P1_POS, P2_POS, 0, 0}: 1}
	rolls := map[int]int64{}
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				rolls[r1+r2+r3]++
			}
		}
	}
	fmt.Printf("die: %#v\n", rolls)

	var p1wins, p2wins int64 = 0, 0

	for step := 1; step <= 100; step++ {
		for turn := 1; turn <= 2; turn++ {
			nextStates := map[State]int64{}
			fmt.Printf("step %d turn %d size: %d\n", step, turn, len(states))
			for state, count := range states {
				for roll, rCount := range rolls {
					p1, p2 := state.p1, state.p2
					p1score, p2score := state.p1score, state.p2score
					n := count * rCount

					if turn == 1 {
						p1 = Advance(p1, roll)
						p1score += p1
						if p1score >= 21 {
							p1wins += n
						} else {
							nextStates[State{p1, p2, p1score, p2score}] += n
						}
					} else {
						p2 = Advance(p2, roll)
						p2score += p2
						if p2score >= 21 {
							p2wins += n
						} else {
							nextStates[State{p1, p2, p1score, p2score}] += n
						}
					}
				}
			}
			states = nextStates
			fmt.Printf("p1wins: %d, p2wins: %d\n", p1wins, p2wins)
			if len(states) == 0 {
				return
			}
		}
	}
}
