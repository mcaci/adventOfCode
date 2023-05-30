package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mcaci/graphgo/graph"
	"github.com/mcaci/graphgo/path"
)

type valve struct {
	name     string
	pressure int
}

func (c valve) String() string { return fmt.Sprintf("(%q, %d)", c.name, c.pressure) }

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	valves := parseInput(f)

	// part 1
	fmt.Println("Part 1: ", part1(valves))
	// part 2
}

func parseInput(r io.Reader) graph.Graph[valve] {
	scanner := bufio.NewScanner(r)
	valves := graph.New[valve](graph.ArcsListType)
	var cs []struct{ id, next string }
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()
		id := line[6:8]
		next := strings.Trim(line[49:], " ")
		cs = append(cs, struct {
			id   string
			next string
		}{id: id, next: next})
		rStr := line[23:strings.Index(line, ";")]
		rate, _ := strconv.Atoi(rStr)
		valves.AddVertex(&graph.Vertex[valve]{E: valve{name: id, pressure: rate}})
	}
	for _, c := range cs {
		_, x, _ := getCave(valves, c.id)
		ns := strings.Split(c.next, ", ")
		for _, n := range ns {
			_, y, _ := getCave(valves, n)
			valves.AddEdge(&graph.Edge[valve]{X: x, Y: y, P: graph.EdgeProperty{W: 1}})
		}
	}
	return valves
}

func part1(valves graph.Graph[valve]) int {
	const maxMinutes = 30
	var workingValves []*graph.Vertex[valve]
	for _, v := range valves.Vertices() {
		if v.E.pressure <= 0 {
			continue
		}
		workingValves = append(workingValves, v)
	}

	distances := updateDistances(valves)
	return (&state{tunnelID: "AA"}).update(distances, workingValves, maxMinutes)
}

func updateDistances(valves graph.Graph[valve]) map[string]map[*graph.Vertex[valve]]*path.Distance[valve] {
	distances := make(map[string]map[*graph.Vertex[valve]]*path.Distance[valve])
	verts := valves.Vertices()
	for _, v := range verts {
		distances[v.E.name] = path.BellmanFordDist(valves, v)
	}
	return distances
}

type state struct {
	t, p, f  int
	tunnelID string
}

func (st *state) update(distances map[string]map[*graph.Vertex[valve]]*path.Distance[valve], workingValves []*graph.Vertex[valve], maxMinutes int) int {
	currentScore := st.p + (maxMinutes-st.t)*st.f
	maxScore := currentScore

	for i, valve := range workingValves {
		fmt.Println(distances[st.tunnelID])
		fmt.Println(valve)
		dTime := distances[st.tunnelID][valve].Dist()
		openTime := dTime + 1
		if st.t+openTime >= maxMinutes {
			continue
		}
		nextState := &state{t: st.t + openTime, p: st.p + openTime*st.f, f: st.f + valve.E.pressure}
		candidateScore := nextState.update(distances, append(workingValves[:i], workingValves[i+1:]...), maxMinutes)
		if candidateScore > maxScore {
			maxScore = candidateScore
		}
	}
	return maxScore
}

func getCave(g graph.Graph[valve], name string) (int, *graph.Vertex[valve], error) {
	for i, u := range g.Vertices() {
		if u.E.name != name {
			continue
		}
		return i, u, nil
	}
	return 0, nil, errors.New("vertex not found")
}
