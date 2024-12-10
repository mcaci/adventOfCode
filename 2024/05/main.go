package main

import (
	_ "embed"
	"fmt"
	"log"
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
			rules = append(rules, PageOrderingRule(&Tree[int]{element: &x, children: []*Tree[int]{{element: &y}}}))
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
	var movedItems []int
	for i, r1 := range rules {
		for j, r2 := range rules {
			if i == j {
				continue
			}
			for k, c2 := range r2.children {
				if *c2.element != *r1.element {
					continue
				}
				r2.children[k] = r1
				movedItems = append(movedItems, *r1.element)
			}
		}
	}
	var rule PageOrderingRule
	for i, r := range rules {
		if slices.Contains(movedItems, *r.element) {
			continue
		}
		// log.Print(((*Tree[int])(r)).String())
		rule = rules[i]
	}
	// for _, u := range updates {
	// 	if !existsPath(u, rule) {
	// 		continue
	// 	}
	// 	log.Print(u)
	// }
	(*Tree[int])(rule).Traverse()
	log.Print((*Tree[int])(rule).RecordTraversal())
}

type PageOrderingRule *Tree[int]

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

func existsPath(u Update, r PageOrderingRule) bool {
	// var rs [][]int
	// var stack []int
	// for i := range r {
	// 	if r.element ==
	// }
	return true
}

func (t *Tree[T]) Traverse() {
	log.Print(*t.element)
	for i := range t.children {
		t.children[i].Traverse()
	}
}

func (t *Tree[T]) RecordTraversal() [][]T {
	if t == nil || t.element == nil {
		return nil
	}
	var record [][]T
	tStack := []*Tree[T]{t}
	var path []T
	for len(tStack) > 0 {
		var newT *Tree[T]
		newT, tStack = tStack[0], tStack[1:]
		if newT == nil || newT.element == nil {
			record = append(record, slices.Clone(path))
			path = path[:len(path)-1]
			continue
		}
		path = append(path, *newT.element)
		if newT.children == nil {
			record = append(record, slices.Clone(path))
			path = path[:len(path)-1]
			continue
		}
		tStack = append(tStack, newT.children...)
	}
	return record
}
