package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/mcaci/graphgo/graph"
)

//go:embed sample1
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	g := graph.New[pipe](graph.ArcsListType)
	var j int
	for scanner.Scan() {
		line := scanner.Text()
		for i, d := range line {
			if d == '.' {
				continue
			}
			g.AddVertex(&graph.Vertex[pipe]{E: pipe{x: i, y: j, dir: direction(line[i])}})
		}
		j++
	}
	vs := g.Vertices()
	for i, v := range vs {
		var s, e *graph.Vertex[pipe]
		log.Print(v)
		switch v.E.dir {
		case start:
			continue
		case horizontal:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y {
					s = w
					continue
				}
				if v.E.x == w.E.x-1 && v.E.y == w.E.y {
					e = w
					continue
				}

			}
		case vertical:
			for j, w := range vs {
				if i == j {
					continue
				}
				if v.E.y == w.E.y+1 && v.E.x == w.E.x {
					s = w
					continue
				}
				if v.E.y == w.E.y-1 && v.E.x == w.E.x {
					e = w
					continue
				}
				if s != nil && e != nil {
					log.Print("adding", v)
					addEdges(g, v, s, e)
					break
				}
			}
		case northPlusEast:
		case northPlusWest:
		case southPlusEast:
		case southPlusWest:
		}
		if s != nil && e != nil {
			log.Print("adding", v)
			addEdges(g, v, s, e)
		}
	}
	log.Print(g, run())
}

func run() int {
	return 0
}

type pipe struct {
	x, y int
	dir  direction
}

func (p pipe) String() string {
	return fmt.Sprintf("(%d,%d):%s", p.x, p.y, string(p.dir))
}

type direction byte

const (
	start         = direction('S')
	horizontal    = direction('-')
	vertical      = direction('|')
	northPlusEast = direction('L')
	northPlusWest = direction('J')
	southPlusWest = direction('7')
	southPlusEast = direction('F')
)

func addEdges(g graph.Graph[pipe], v, s, e *graph.Vertex[pipe]) {
	sE := &graph.Edge[pipe]{X: v, Y: s}
	if !g.ContainsEdge(sE) {
		g.AddEdge(sE)
	}
	eE := &graph.Edge[pipe]{X: v, Y: e}
	if !g.ContainsEdge(eE) {
		g.AddEdge(eE)
	}
}
