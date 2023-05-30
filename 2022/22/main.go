package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	field, moves := parseInput(f)
	h := start(field)
	hP2 := *h
	for i := range field {
		fmt.Printf("%q\n", field[i])
	}
	for i := range moves {
		singleInstr(field, h, moves[i], nil)
	}
	h.i++
	h.j++
	fmt.Println(moves, h)

	// part 1
	fmt.Println("part 1:", h.i*1000+h.j*4+int(h.d))

	// part 2
	wrapInfo := make(map[humn]humn)
	const _50 = 50
	for k := 0; k < _50; k++ {
		wrapInfo[humn{i: _50 * 0, j: _50*1 + k, d: up}] = humn{i: _50*3 + k, j: _50 * 0, d: right}
		wrapInfo[humn{i: _50 * 0, j: _50*2 + k, d: up}] = humn{i: _50*4 - 1, j: _50*0 + k, d: up}
		wrapInfo[humn{i: _50*1 - 1, j: _50*2 + k, d: down}] = humn{i: _50*1 + k, j: _50*2 - 1, d: left}
		wrapInfo[humn{i: _50 * 2, j: _50*0 + k, d: up}] = humn{i: _50*1 + k, j: _50, d: right}
		wrapInfo[humn{i: _50*3 - 1, j: _50*1 + k, d: down}] = humn{i: _50*3 + k, j: _50 - 1, d: left}
		wrapInfo[humn{i: _50*4 - 1, j: _50*0 + k, d: down}] = humn{i: _50 * 0, j: _50*2 + k, d: down}

		wrapInfo[humn{i: _50*3 - (k + 1), j: 0, d: left}] = humn{i: _50*0 + k, j: _50 * 1, d: right}
		wrapInfo[humn{i: _50*3 + k, j: 0, d: left}] = humn{i: _50 * 0, j: _50*1 + k, d: down}
		wrapInfo[humn{i: _50*3 + k, j: _50 - 1, d: right}] = humn{i: _50*3 - 1, j: _50*1 + k, d: up}
		wrapInfo[humn{i: _50*0 + k, j: _50, d: left}] = humn{i: _50*3 - (k + 1), j: _50 * 0, d: right}
		wrapInfo[humn{i: _50*1 + k, j: _50, d: left}] = humn{i: _50 * 2, j: _50*0 + k, d: down}
		wrapInfo[humn{i: _50*1 + k, j: _50*2 - 1, d: right}] = humn{i: _50*1 - 1, j: _50*2 + k, d: up}
		wrapInfo[humn{i: _50*2 + k, j: _50*2 - 1, d: right}] = humn{i: _50*1 - (k + 1), j: _50*3 - 1, d: left}
		wrapInfo[humn{i: _50*1 - (k + 1), j: _50*3 - 1, d: right}] = humn{i: _50*2 + k, j: _50*2 - 1, d: left}
	}

	for i := range moves {
		singleInstr(field, &hP2, moves[i], wrapInfo)
	}

	hP2.i++
	hP2.j++
	// fmt.Println(moves, h)
	fmt.Println("part 2:", hP2.i*1000+hP2.j*4+int(hP2.d))
}

func singleInstr(m [][]byte, h *humn, instr byte, wrapInfo map[humn]humn) {
	switch instr {
	case 'L':
		h.d = (h.d - 1) % 4
	case 'R':
		h.d = (h.d + 1) % 4
	default:
		for i := 0; i < int(instr); i++ {
			iNext, jNext, dNext := h.i, h.j, h.d
			switch h.d {
			case right:
				jNext++
			case down:
				iNext++
			case left:
				jNext--
			case up:
				iNext--
			}
			if iNext < 0 || jNext < 0 || iNext >= 200 || jNext >= 150 || m[iNext][jNext] == '~' {
				if wrapInfo != nil {
					next := wrapInfo[*h]
					iNext = next.i
					jNext = next.j
					dNext = next.d
				} else {
					switch h.d {
					case right:
						jNext = bytes.IndexFunc(m[iNext], func(r rune) bool { return r == '#' || r == '.' })
					case down:
						var col []byte
						for c := range m {
							col = append(col, m[c][jNext])
						}
						iNext = bytes.IndexFunc(col, func(r rune) bool { return r == '#' || r == '.' })
					case left:
						jNext = bytes.LastIndexFunc(m[iNext], func(r rune) bool { return r == '#' || r == '.' })
					case up:
						var col []byte
						for c := range m {
							col = append(col, m[c][jNext])
						}
						iNext = bytes.LastIndexFunc(col, func(r rune) bool { return r == '#' || r == '.' })
					}
				}
			}
			if m[iNext][jNext] == '#' {
				break
			}
			h.i = iNext
			h.j = jNext
			h.d = dNext
		}
	}
}

type humn struct {
	i, j int
	d    dir
}

type dir uint8

const (
	right dir = iota
	down
	left
	up
)

func start(m [][]byte) *humn {
	for i := range m {
		for j := range m[i] {
			if m[i][j] != '.' {
				continue
			}
			return &humn{i: i, j: j, d: right}
		}
	}
	return nil
}

func parseInput(r io.Reader) ([][]byte, []byte) {
	scanner := bufio.NewScanner(r)
	var maxL int
	var nodes [][]byte
	var instrLine string
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if unicode.In(rune(line[0]), unicode.Digit, unicode.Letter) {
			instrLine = line
			continue
		}
		if len(line) > maxL {
			maxL = len(line)
		}
		walledLine := bytes.ReplaceAll([]byte(line), []byte{' '}, []byte{'~'})
		nodes = append(nodes, walledLine)
	}
	for i := range nodes {
		if len(nodes[i]) == maxL {
			continue
		}
		nodes[i] = append(nodes[i], bytes.Repeat([]byte{'~'}, maxL-len(nodes[i]))...)
	}
	var instrs []byte
	var intermediate []byte
	for i, c := range instrLine {
		if !unicode.IsLetter(c) {
			intermediate = append(intermediate, instrLine[i])
			continue
		}
		if len(intermediate) > 0 {
			n, _ := strconv.Atoi(string(intermediate))
			instrs = append(instrs, byte(n))
			intermediate = nil
		}
		instrs = append(instrs, instrLine[i])
	}
	if len(intermediate) > 0 {
		n, _ := strconv.Atoi(string(intermediate))
		instrs = append(instrs, byte(n))
	}
	return nodes, instrs
}
