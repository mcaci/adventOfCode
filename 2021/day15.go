package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	day15(1) // part1
	day15(5) // part2
}

func day15(times int) {
	f, err := os.Open("day15")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	maxRow := 100
	var cavern [][]cave
	for i := 0; i < maxRow*times; i++ {
		cavern = append(cavern, make([]cave, maxRow*times))
	}
	var distance [][]int
	for i := 0; i < maxRow*times; i++ {
		distance = append(distance, make([]int, maxRow*times))
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[0][0] = 0
	var lastLine bool
	for ln := 0; !lastLine; ln++ {
		l, err := r.ReadString('\n')
		switch err {
		case nil:
			l = l[:len(l)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		for i := range l {
			d, err := strconv.Atoi(string(l[i]))
			if err != nil {
				log.Fatal(err)
			}
			cavern[ln][i] = cave{i: ln, j: i, cost: d}
		}
	}
	if times > 1 {
		for i := range cavern {
			for j := range cavern[i] {
				if i < 100 && j < 100 {
					continue
				}
				cellToCopy := cavern[i%maxRow][j%maxRow]
				offset := i/maxRow + j/maxRow
				cost := (cellToCopy.cost+offset-1)%9 + 1
				cavern[i][j] = cave{i: i, j: j, cost: cost}
			}
		}
	}
	// start search
	queue := []cave{cavern[0][0]}
	for len(queue) > 0 {
		cave := queue[0]
		queue = queue[1:]
		for _, n := range cave.neighbor(cavern) {
			neighborCost := distance[n.i][n.j]
			costToNeighbor := distance[cave.i][cave.j] + int(n.cost)
			if neighborCost > costToNeighbor {
				distance[n.i][n.j] = costToNeighbor
				queue = append(queue, n)
			}
		}
	}
	fmt.Println(distance[(maxRow*times)-1][(maxRow*times)-1])
}

type cave struct {
	i, j int
	cost int
}

func (c cave) neighbor(cavern [][]cave) []cave {
	const max_val = 99
	var n []cave
	if c.i < max_val {
		n = append(n, cavern[c.i+1][c.j])
	}
	if c.i > 0 {
		n = append(n, cavern[c.i-1][c.j])
	}
	if c.j < max_val {
		n = append(n, cavern[c.i][c.j+1])
	}
	if c.j > 0 {
		n = append(n, cavern[c.i][c.j-1])
	}
	return n
}
