package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"
)

type direction byte

func (d direction) String() string { return string(d) }

const (
	L = direction('L')
	R = direction('R')
)

type node struct {
	key  string
	l, r *node
}

func (n node) String() string { return fmt.Sprintf("%s-><%s-%s>", n.key, n.l.key, n.r.key) }

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var i int
	var dirs []direction
	var nodes []*node
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			i++
			for i := range line {
				dirs = append(dirs, direction(line[i]))
			}
			continue
		}
		if len(line) == 0 {
			continue
		}
		var n, nl, nr *node
		k, kl, kr := line[:3], line[7:10], line[12:15]
		ik := slices.IndexFunc(nodes, func(o *node) bool { return o.key == k })
		switch ik {
		case -1:
			n = &node{key: k}
			nodes = append(nodes, n)
		default:
			n = nodes[ik]
		}
		ikl := slices.IndexFunc(nodes, func(o *node) bool { return o.key == kl })
		switch ikl {
		case -1:
			nl = &node{key: kl}
			nodes = append(nodes, nl)
		default:
			nl = nodes[ikl]
		}
		ikr := slices.IndexFunc(nodes, func(o *node) bool { return o.key == kr })
		switch ikr {
		case -1:
			nr = &node{key: kr}
			nodes = append(nodes, nr)
		default:
			nr = nodes[ikr]
		}
		n.l = nl
		n.r = nr
		// log.Print(line)
	}
	// log.Print(dirs, nodes)
	log.Print("Part 1:", runPart1(dirs, nodes))
	log.Print("Part 2:", runPart2(dirs, nodes))

}

func runPart1(dirs []direction, nodes []*node) int {
	n := nodes[slices.IndexFunc(nodes, func(o *node) bool { return o.key == "AAA" })]
	for i := 0; true; i++ {
		if n.key == "ZZZ" {
			return i
		}
		dir := dirs[i%len(dirs)]
		switch dir {
		case L:
			n = n.l
		case R:
			n = n.r
		default:
			log.Fatal("nope")
		}
	}
	return 0
}

func runPart2(dirs []direction, nodes []*node) int {
	var startingNodes []*node
	var endingNodes []*node
	for i := range nodes {
		switch nodes[i].key[2] {
		case 'A':
			startingNodes = append(startingNodes, nodes[i])
		case 'Z':
			endingNodes = append(endingNodes, nodes[i])
		default:
		}
	}

	// bruteforce not working... putting limit and find lcm of first occurrence of each Z entry
	for i := 0; i < 30000; i++ {
		count := make(map[string]uint)
		for _, e := range endingNodes {
			for _, s := range startingNodes {
				if s.key != e.key {
					continue
				}
				count[e.key]++
			}
		}
		if len(count) == len(endingNodes) {
			return i
		}
		if len(count) > 0 {
			log.Print(i, startingNodes)
		}
		dir := dirs[i%len(dirs)]
		for i := range startingNodes {
			switch dir {
			case L:
				startingNodes[i] = startingNodes[i].l
			case R:
				startingNodes[i] = startingNodes[i].r
			default:
				log.Fatal("nope")
			}
		}
	}
	// for part 2 after checking the log of the first occurrences of each Z entry
	return LCM(22357, 17263, 14999, 16697, 13301, 20659)
}

func GCD(a, b int) int {
	for b > 0 {
		b, a = a%b, b
	}
	return a
}

func LCM(ns ...int) int {
	switch len(ns) {
	case 0:
		return 0
	case 1:
		return ns[0]
	default:
		lcm := ns[0]
		for i := 1; i < len(ns); i++ {
			lcm = lcm * ns[i] / GCD(lcm, ns[i])
		}
		return lcm
	}
}
