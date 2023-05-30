package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	l               = 7
	h               = 5
	rockIcon        = '#'
	empty           = '.'
	p1Rocks         = 2022
	p2Rocks         = 1000000000000
	fullCycleRocks  = 3242 // taken from map patternSearch
	cycles          = p2Rocks / fullCycleRocks
	remainderRocks  = p2Rocks % fullCycleRocks
	hPerCycle       = 5176 // line 42, use fullCycleRocks as limit
	hRemainderRocks = 540  // line 42, use remainderRocks as limit
)

func main() {
	jetPattern, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	chamber := make([][]byte, h)
	for i := range chamber {
		chamber[i] = make([]byte, l)
		for j := range chamber[i] {
			chamber[i][j] = empty
		}
	}

	var rockID int
	var jetID int

	for round := 1; round <= fullCycleRocks; round++ {
		rock := rocks[rockID]
		var topRow int
		for row := range chamber {
			if bytes.Contains(chamber[row], []byte{rockIcon}) {
				continue
			}
			topRow = row
			break
		}
		if len(chamber)-topRow < 10 {
			for i := 0; i < 10; i++ {
				row := make([]byte, l)
				for j := range row {
					row[j] = empty
				}
				chamber = append(chamber, row)
			}
		}
		rockCoord := spawn(rock, topRow, chamber)
		var stop bool
		for !stop {
			// fmt.Println("push", rockCoord, "to", string(jetPattern[jetID]), ":", jetID)
			jetPush(rockCoord, jetPattern[jetID%len(jetPattern)], chamber)
			jetID++
			stop = dropDown(rockCoord, chamber)
			// if jetID == len(rocks)*len(jetPattern) {
			// break
			// }
		}
		drawCoords(rockCoord, chamber)
		rockID = (rockID + 1) % len(rocks)
	}

	// for i := 10; i >= 0; i-- {
	// 	fmt.Println(string(chamber[i]))
	// }

	patternSearch := make(map[string]int)
	for i := range chamber {
		if string(chamber[i]) == strings.Repeat(".", l) {
			break
		}
		if string(chamber[i]) != ".#####." {
			continue
		}
		if string(chamber[i+1]) != ".#.#..." {
			continue
		}
		if i > 0 {
			fmt.Println(i, string(chamber[i+1]), string(chamber[i]), string(chamber[i-1]))
			if patternSearch[string(chamber[i-1])] > 0 {
				continue
			}
			patternSearch[string(chamber[i-1])] = i
		}
	}
	fmt.Println(patternSearch)

	// part 1
	var topRow int
	for i := len(chamber) - 1; i >= 0; i-- {
		if !bytes.Contains(chamber[i], []byte{rockIcon}) {
			continue
		}
		topRow = i
		break
	}
	fmt.Println("total: ", topRow+1)
	// part 2
	fmt.Println(len(jetPattern), len(jetPattern)*5, p2Rocks/fullCycleRocks, p2Rocks%fullCycleRocks)
	fmt.Println("total: ", (hPerCycle*cycles)+hRemainderRocks+1)
	// too low  1594787434351
	// too high 1596545342389
}

type coord struct{ x, y int }

func spawn(r rock, topRow int, chamber [][]byte) []*coord {
	const (
		bottomPosOrXPos = 3
		leftPosOrYPos   = 2
	)
	row := bottomPosOrXPos + topRow
	col := leftPosOrYPos
	var coords []*coord
	for _, f := range r.spawn() {
		r, c := f(row, col)
		coords = append(coords, &coord{r, c})
		// fmt.Print("(", r, ",", c, ")")
	}
	return coords
}

func jetPush(xys []*coord, jet byte, chamber [][]byte) {
	switch jet {
	case '<':
		for _, xy := range xys {
			y1 := xy.y - 1
			if y1 < 0 {
				// fmt.Println("wall hit")
				return
			}
			if chamber[xy.x][y1] == rockIcon {
				// fmt.Println("rock hit", xy.x, y1)
				return
			}
		}
		for i := range xys {
			xys[i].y--
		}
	case '>':
		for _, xy := range xys {
			y1 := xy.y + 1
			if y1 > l-1 {
				// fmt.Println("wall hit")
				return
			}
			if chamber[xy.x][y1] == rockIcon {
				// fmt.Println("rock hit", xy.x, y1)
				return
			}
		}
		for i := range xys {
			xys[i].y++
		}
	}
}

func dropDown(xys []*coord, chamber [][]byte) bool {
	for _, xy := range xys {
		lowerX := xy.x - 1
		if lowerX < 0 {
			return true
		}
		if chamber[lowerX][xy.y] == rockIcon {
			// fmt.Println("rock hit", lowerX, xy.y)
			return true
		}
	}

	for i := range xys {
		xys[i].x--
	}
	return false
}

func drawCoords(xys []*coord, chamber [][]byte) {
	for _, xy := range xys {
		r, c := xy.x, xy.y
		chamber[r][c] = rockIcon
	}
}

type rock string

const (
	rockA = rock("####")
	rockB = rock(" # \n###\n # ")
	rockC = rock("  #\n  #\n###")
	rockD = rock("#\n#\n#\n#")
	rockE = rock("##\n##")
)

var (
	rocks = []rock{rockA, rockB, rockC, rockD, rockE}
)

func (r rock) spawn() []func(int, int) (int, int) {
	switch r {
	case rockA:
		return []func(int, int) (int, int){
			func(i1, i2 int) (int, int) { return i1, i2 },
			func(i1, i2 int) (int, int) { return i1, i2 + 1 },
			func(i1, i2 int) (int, int) { return i1, i2 + 2 },
			func(i1, i2 int) (int, int) { return i1, i2 + 3 },
		}
	case rockB:
		return []func(int, int) (int, int){
			func(i1, i2 int) (int, int) { return i1, i2 + 1 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 + 1 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 + 2 },
			func(i1, i2 int) (int, int) { return i1 + 2, i2 + 1 },
		}
	case rockC:
		return []func(int, int) (int, int){
			func(i1, i2 int) (int, int) { return i1 + 2, i2 + 2 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 + 2 },
			func(i1, i2 int) (int, int) { return i1, i2 },
			func(i1, i2 int) (int, int) { return i1, i2 + 1 },
			func(i1, i2 int) (int, int) { return i1, i2 + 2 },
		}
	case rockD:
		return []func(int, int) (int, int){
			func(i1, i2 int) (int, int) { return i1, i2 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 },
			func(i1, i2 int) (int, int) { return i1 + 2, i2 },
			func(i1, i2 int) (int, int) { return i1 + 3, i2 },
		}
	case rockE:
		return []func(int, int) (int, int){
			func(i1, i2 int) (int, int) { return i1, i2 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 },
			func(i1, i2 int) (int, int) { return i1, i2 + 1 },
			func(i1, i2 int) (int, int) { return i1 + 1, i2 + 1 },
		}
	}

	return nil
}
