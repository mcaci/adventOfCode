package main

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"
	"unicode"
)

//go:embed sample
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var parts []machinePart
	var flows []workflow
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case len(line) == 0:
			continue
		case line[0] == '{':
			// parse machine part
			l := strings.Map(func(r rune) rune {
				if !unicode.IsDigit(r) {
					return ' '
				}
				return r
			}, line)
			f := strings.Split(l, " ")
			var xmas []int
			for i := range f {
				if f[i] == "" {
					continue
				}
				n, _ := strconv.Atoi(f[i])
				xmas = append(xmas, n)
			}
			parts = append(parts, machinePart{x: xmas[0], m: xmas[1], a: xmas[2], s: xmas[3]})
		default:
			// parse condition
			var w workflow
			stepsLine := line[strings.IndexByte(line, '{')+1 : strings.IndexByte(line, '}')]
			f := strings.Split(stepsLine, ",")
			var steps []step
			for i := range f {
				var s step
				switch strings.ContainsRune(f[i], ':') {
				case true:
					sep := strings.IndexRune(f[i], ':')
					c := f[i][:sep]
					switch {
					case strings.ContainsRune(c, '<'):
						sign := strings.IndexRune(c, '<')
						l := c[:sign]
						r, _ := strconv.Atoi(c[sign+1:])
						switch l {
						case "x":
							s.cond = func(mp machinePart) bool { return mp.x < r }
						case "m":
							s.cond = func(mp machinePart) bool { return mp.m < r }
						case "a":
							s.cond = func(mp machinePart) bool { return mp.a < r }
						case "s":
							s.cond = func(mp machinePart) bool { return mp.s < r }
						default:
							log.Fatal("unexpected left side in ", line)
						}
					case strings.ContainsRune(c, '>'):
						sign := strings.IndexRune(c, '>')
						l := c[:sign]
						r, _ := strconv.Atoi(c[sign+1:])
						switch l {
						case "x":
							s.cond = func(mp machinePart) bool { return mp.x > r }
						case "m":
							s.cond = func(mp machinePart) bool { return mp.m > r }
						case "a":
							s.cond = func(mp machinePart) bool { return mp.a > r }
						case "s":
							s.cond = func(mp machinePart) bool { return mp.s > r }
						default:
							log.Fatal("unexpected right side in ", line)
						}
					default:
						log.Fatal("unexpected condition in ", line)
					}
					s.next = f[i][sep+1:]
				default:
					s.cond = always
					s.next = f[i]
				}
				steps = append(steps, s)
			}
			w.n = line[:strings.IndexByte(line, '{')]
			w.s = steps
			flows = append(flows, w)
		}
	}
	log.Print(parts)
	log.Print(flows)
}

type machinePart struct{ x, m, a, s int }
type workflow struct {
	n string
	s []step
}
type step struct {
	cond func(machinePart) bool
	next string
}

func always(machinePart) bool { return true }
