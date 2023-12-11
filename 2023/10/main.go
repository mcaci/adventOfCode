package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math"
	"sort"
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
	var startNode *graph.Vertex[pipe]
	for _, v := range vs {
		var s, e *graph.Vertex[pipe]
		switch v.E.dir {
		case start:
			startNode = v
			continue
		case horizontal:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusEast || w.E.dir == northPlusEast) {
					s = w
					continue
				}
				if v.E.x == w.E.x-1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusWest || w.E.dir == northPlusWest) {
					e = w
					continue
				}
			}
		case vertical:
			for _, w := range vs {
				if v.E.y == w.E.y+1 && v.E.x == w.E.x && (w.E.dir == start || w.E.dir == vertical || w.E.dir == southPlusWest || w.E.dir == southPlusEast) {
					s = w
					continue
				}
				if v.E.y == w.E.y-1 && v.E.x == w.E.x && (w.E.dir == start || w.E.dir == vertical || w.E.dir == northPlusWest || w.E.dir == northPlusEast) {
					e = w
					continue
				}
			}
		case northPlusEast:
			for _, w := range vs {
				if v.E.x == w.E.x-1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusWest || w.E.dir == northPlusWest) {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y+1 && (w.E.dir == start || w.E.dir == vertical || w.E.dir == southPlusWest || w.E.dir == southPlusEast) {
					e = w
					continue
				}
			}
		case northPlusWest:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusEast || w.E.dir == northPlusEast) {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y+1 && (w.E.dir == start || w.E.dir == vertical || w.E.dir == southPlusWest || w.E.dir == southPlusEast) {
					e = w
					continue
				}
			}
		case southPlusEast:
			for _, w := range vs {
				if v.E.x == w.E.x-1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusWest || w.E.dir == northPlusWest) {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y-1 && (w.E.dir == start || w.E.dir == vertical || w.E.dir == northPlusWest || w.E.dir == northPlusEast) {
					e = w
					continue
				}
			}
		case southPlusWest:
			for _, w := range vs {
				if v.E.x == w.E.x+1 && v.E.y == w.E.y && (w.E.dir == start || w.E.dir == horizontal || w.E.dir == southPlusEast || w.E.dir == northPlusEast) {
					s = w
					continue
				}
				if v.E.x == w.E.x && v.E.y == w.E.y-1 && (w.E.dir == start || w.E.dir == vertical || w.E.dir == northPlusWest || w.E.dir == northPlusEast) {
					e = w
					continue
				}
			}
		default:
			log.Fatal(string(v.E.dir), " nope")
		}
		if s != nil && e != nil {
			addEdges(g, v, s, e)
		}
	}
	d := path.BellmanFordDist(g, startNode)
	var keys []*graph.Vertex[pipe]
	for k := range d {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return d[keys[i]].Dist() < d[keys[j]].Dist()
	})
	for _, k := range keys {
		if d[k].Dist() >= math.MaxInt {
			break
		}
		log.Print(k, d[k])
	}
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
