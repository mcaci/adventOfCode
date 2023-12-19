package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log"
	"slices"
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
	}
	// tilt(f, N)
	// log.Println("part 1: ", score(f))
	var seq []int
	buckets := make(map[int]int)
	for i := 0; i < 1000000000; i++ {
		fCopy := make(field, len(f))
		for i := range f {
			fCopy[i] = make([]byte, len(f[i]))
			copy(fCopy[i], f[i])
		}
		cycle(f)

		s := score(f)
		seq = append(seq, s)
		buckets[s]++
		var count, first int
		for k, v := range buckets {
			switch v {
			case 2:
				count++
			case 3:
				count++
				first = k
			default:
			}
		}
		iFirst := slices.Index(seq, first)
		if iFirst < 0 {
			continue
		}
		log.Println(iFirst, "/", first, "/", seq)
		break
	}
	for i := range seq {
		var seq1, seq2 []int
		for j := range seq[i+1:] {
			seq1 = append(seq1, seq[j-(i+1)])
			seq2 = append(seq2, seq[j])
			if !slices.Equal(seq1, seq2) {
				continue
			}
			log.Println(seq1, seq2)
		}
	}
	log.Println("part 2: ", score(f))
}

type field [][]byte

func cycle(f field) {
	tilt(f, N)
	tilt(f, W)
	tilt(f, S)
	tilt(f, E)
}

func tilt(f field, d direction) {
	for {
		var rolled bool
	nextLine:
		for i := range f {
			for j := range f[i] {
				switch f[i][j] {
				case 'O':
					switch d {
					case N:
						if i == 0 {
							continue nextLine
						}
						if f[i-1][j] != '.' {
							continue
						}
						f[i-1][j] = 'O'
					case E:
						if j == len(f[i])-1 {
							continue
						}
						if f[i][j+1] != '.' {
							continue
						}
						f[i][j+1] = 'O'
					case S:
						if i == len(f)-1 {
							continue nextLine
						}
						if f[i+1][j] != '.' {
							continue
						}
						f[i+1][j] = 'O'
					case W:
						if j == 0 {
							continue
						}
						if f[i][j-1] != '.' {
							continue
						}
						f[i][j-1] = 'O'
					}
				default:
					continue
				}
				f[i][j] = '.'
				rolled = true
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
		score += s
	}
	return score
}

type direction byte

func (d direction) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	default:
		return "#"
	}
}

const (
	N = direction(iota)
	E
	S
	W
)
