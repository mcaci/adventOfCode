package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/path"
)

//go:embed input
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
	for _, v := range vs {
		var s, e *graph.Vertex[pipe]
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
			for _, w := range vs {
				if v.E.y == w.E.y+1 && v.E.x == w.E.x {
					s = w
					continue
				}
				if v.E.y == w.E.y-1 && v.E.x == w.E.x {
					e = w
					continue
				}
			}
		case northPlusEast:
			for _, w := range vs {
				if v.E.x == w.E.x-1 && v.E.y == w.E.y {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y+1 {
					e = w
					continue
				}
			}
		case northPlusWest:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y+1 {
					e = w
					continue
				}
			}
		case southPlusEast:
			for _, w := range vs {
				if v.E.x == w.E.x-1 && v.E.y == w.E.y {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y-1 {
					e = w
					continue
				}
			}
		case southPlusWest:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y-1 {
					e = w
					continue
				}
			}
		default:
			log.Fatal("nope")
		}
		if s != nil && e != nil {
			addEdges(g, v, s, e)
		}
	}
	log.Print(run(g))
}

func run(g graph.Graph[pipe]) int {
	es := g.Edges()
	for i := range es {
		log.Print(es[i])
	}
	vs := g.Vertices()
	var s *graph.Vertex[pipe]
	for i := range vs {
		if vs[i].E.dir != start {
			continue
		}
		s = vs[i]
		break
	}
	m := path.BellmanFordDist[pipe](g, s)
	var maxDist int
	for _, d := range m {
		if d.Dist() == math.MaxInt {
			continue
		}
		log.Print(d)
		if d.Dist() < maxDist {
			continue
		}
		maxDist = d.Dist()
	}
	return maxDist
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
	sE := &graph.Edge[pipe]{X: v, Y: s, P: graph.EdgeProperty{W: 1}}
	if !g.ContainsEdge(sE) {
		g.AddEdge(sE)
	}
	eE := &graph.Edge[pipe]{X: v, Y: e, P: graph.EdgeProperty{W: 1}}
	if !g.ContainsEdge(eE) {
		g.AddEdge(eE)
	}
}
