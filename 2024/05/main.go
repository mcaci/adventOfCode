package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed sample
var in string

func main() {
	var switchUpdate bool
	var rules []PageOrderingRule
	var updates []Update
nextLine:
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			switchUpdate = true
		}
		if !switchUpdate {
			line := strings.Split(line, "|")
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			for i := range rules {
				if *rules[i].element != x {
					continue
				}
				rules[i].children = append(rules[i].children, &Tree[int]{element: &y})
				continue nextLine
			}
			rules = append(rules, PageOrderingRule(Tree[int]{element: &x, children: []*Tree[int]{{element: &y}}}))
			continue
		}
		l := strings.Split(line, ",")
		var u Update
		for _, v := range l {
			n, _ := strconv.Atoi(v)
			u = append(u, n)
		}
		updates = append(updates, u)
	}
	for len(rules) > 1 {
	}
}

type PageOrderingRule Tree[int]
type Update []int

type Tree[T comparable] struct {
	element  *T
	children []*Tree[T]
}

func (t *Tree[T]) Size() int {
	switch {
	case t == nil, t.element == nil:
		return 0
	case t.children == nil:
		return 1
	default:
		size := 1
		for i := range t.children {
			size += t.children[i].Size()
		}
		return size
	}
}

func (t *Tree[T]) Find(e *T) *Tree[T] {
	switch {
	case t == nil, t.element == nil:
		return nil
	case *t.element == *e:
		return t
	case t.children == nil:
		return nil
	default:
		i := slices.IndexFunc(t.children, func(t *Tree[T]) bool { return t.Find(e) != nil })
		if i < 0 {
			return nil
		}
		return t.children[i]
	}
}

func (t *Tree[T]) Contains(e *T) bool { return t.Find(e) != nil }

func (t *Tree[T]) String() string {
	switch {
	case t == nil, t.element == nil:
		return "()"
	case t.children == nil:
		return fmt.Sprintf("(element:%v)", *t.element)
	default:
		var children []string
		for i := range t.children {
			children = append(children, t.children[i].String())
		}
		return fmt.Sprintf("(element:%v,\nchildren:%v)", *t.element, strings.Join(children, ","))
	}
}
