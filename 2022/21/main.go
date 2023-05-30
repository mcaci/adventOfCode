package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// for i := range nodes {
	// 	fmt.Println(nodes[i])
	// }

	// part 1
	nodes := parseInput(f)
	subTreeHasHuman(getNode(nodes, "root"))
	fmt.Println("part 1: ", getNode(nodes, "root").f.op())

	// part 2
	walkDown(getNode(nodes, "root"))
	fmt.Println("part 2: ", getNode(nodes, "humn").e)
}

type opper interface{ op() int }

type value int

func (v value) op() int { return int(v) }

type operation struct {
	opName       byte
	lName, rName string
	l, r         *node
	f            func(int, int) int
}

func (o operation) op() int { return o.f(o.l.f.op(), o.r.f.op()) }
func add(a, b int) int      { return a + b }
func sub(a, b int) int      { return a - b }
func mul(a, b int) int      { return a * b }
func div(a, b int) int      { return a / b }

type node struct {
	name    string
	e       int
	f       opper
	l, r    *node
	hasHumn bool
}

func (n node) String() string {
	if v, ok := n.f.(value); n.l == nil && ok {
		return fmt.Sprintf("(name:%s, l:nil, r:nil, v:%d)", n.name, v.op())
	}
	if n.l == nil {
		return fmt.Sprintf("(name:%s, l:nil, r:nil)", n.name)
	}
	return fmt.Sprintf("(name:%s, l:%s, r:%s)", n.name, n.l.name, n.r.name)
}

func getNode(nodes []*node, key string) *node {
	for i := range nodes {
		if nodes[i].name != key {
			continue
		}
		return nodes[i]
	}
	return nil
}

func parseInput(r io.Reader) []*node {
	scanner := bufio.NewScanner(r)
	var nodes []*node
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		name := line[:4]
		node := node{name: name}
		rest := strings.TrimLeft(line[4:], ": ")
		if unicode.IsDigit(rune(rest[0])) {
			n, _ := strconv.Atoi(rest)
			node.f = value(n)
			nodes = append(nodes, &node)
			continue
		}
		o := operation{opName: rest[5], lName: rest[:4], rName: rest[len(rest)-4:]}
		switch rest[5] {
		case '+':
			o.f = add
		case '-':
			o.f = sub
		case '*':
			o.f = mul
		case '/':
			o.f = div
		}
		node.f = o
		nodes = append(nodes, &node)
	}
	for i := range nodes {
		switch f := nodes[i].f.(type) {
		case value:
			continue
		case operation:
			nodes[i].l = getNode(nodes, f.lName)
			nodes[i].r = getNode(nodes, f.rName)
			nodes[i].f = operation{opName: f.opName, lName: f.lName, rName: f.rName, l: nodes[i].l, r: nodes[i].r, f: f.f}
		}
	}
	return nodes
}

func subTreeHasHuman(n *node) bool {
	if n.name == "humn" {
		n.hasHumn = true
		return n.hasHumn
	}
	var l, r bool
	if n.l != nil {
		l = subTreeHasHuman(n.l)
	}
	if n.r != nil {
		r = subTreeHasHuman(n.r)
	}
	n.hasHumn = l || r
	return n.hasHumn
}

func walkDown(n *node) {
	switch n.name {
	case "humn":
		return
	case "root":
		l := n.l
		r := n.r
		if l.hasHumn {
			l.e = r.f.op()
			walkDown(l)
		}
		if r.hasHumn {
			r.e = l.f.op()
			walkDown(r)
		}
	default:
		l := n.l
		r := n.r
		if l.hasHumn {
			switch n.f.(operation).opName {
			case '+':
				l.e = n.e - r.f.op()
			case '-':
				l.e = n.e + r.f.op()
			case '*':
				l.e = n.e / r.f.op()
			case '/':
				l.e = n.e * r.f.op()
			}
			walkDown(l)
		}
		if r.hasHumn {
			switch n.f.(operation).opName {
			case '+':
				r.e = n.e - l.f.op()
			case '-':
				r.e = l.f.op() - n.e
			case '*':
				r.e = n.e / l.f.op()
			case '/':
				r.e = l.f.op() / n.e
			}
			walkDown(r)
		}
	}
}

func visit(n *node) {
	stack := []*node{n}
	var i int
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		blanks := strings.Repeat(" ", i)
		fmt.Printf("%s%d %s\n", blanks, i, curr)
		if curr.l != nil {
			i++
			stack = append(stack, curr.l)
		}
		if curr.r != nil {
			stack = append(stack, curr.r)
		}
	}
}
