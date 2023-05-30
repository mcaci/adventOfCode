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

type baseMaterial struct {
	id  materialRobotType
	ore int
}

type advMaterial struct {
	id       materialRobotType
	ore, oth int
}

func (m baseMaterial) canBuild(aOre int) bool      { return aOre-m.ore >= 0 }
func (m advMaterial) canBuild(aOre, aOth int) bool { return aOre-m.ore >= 0 && aOth-m.oth >= 0 }

type materialRobotType int

const (
	geode = materialRobotType(iota)
	obsidian
	clay
	ore
	none
)

type state struct {
	minute                         int
	ore, clay, obsidian, geode     int
	oreR, clayR, obsidianR, geodeR int
}

func next(s state) state {
	return state{
		minute:    s.minute + 1,
		ore:       s.ore + s.oreR,
		clay:      s.clay + s.clayR,
		obsidian:  s.obsidian + s.obsidianR,
		geode:     s.geode + s.geodeR,
		oreR:      s.oreR,
		clayR:     s.clayR,
		obsidianR: s.obsidianR,
		geodeR:    s.geodeR,
	}
}

func (s state) String() string {
	return fmt.Sprintf("minute: %d, resources: %v, robots: %v", s.minute, []int{s.ore, s.clay, s.obsidian, s.geode}, []int{s.oreR, s.clayR, s.obsidianR, s.geodeR})
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	blueprints := parseInput(f)
	for i := range blueprints {
		fmt.Println(i, blueprints[i])
	}

	var sum int
	fmt.Println(geodes(blueprints[1], 24))
	// for _, b := range blueprints {
	// 	g := geodes(b, 24)
	// 	sum += b.id * g
	// 	fmt.Println(g)
	// }
	// part 1

	fmt.Println("part 1: ", sum)
	// part 2

	// fmt.Println("part 2: ", internal)
	fmt.Println("part 2: ", 0)
}

type blueprint struct {
	id       int
	ore      baseMaterial
	clay     baseMaterial
	obsidian advMaterial
	geode    advMaterial
}

func geodes(b blueprint, maxT int) int {
	var clayUnlocked, obsidianUnlocked, geodeUnlocked bool
	var geodesMax int
	states := []state{{oreR: 1}}
	for len(states) > 0 {
		candidate := states[0]
		states = states[1:]

		if clayUnlocked && candidate.clayR == 0 {
			continue
		}
		if obsidianUnlocked && candidate.obsidianR == 0 {
			continue
		}
		if geodeUnlocked && candidate.geodeR == 0 {
			continue
		}

		if candidate.minute > maxT {
			continue
		}

		if candidate.geode > geodesMax {
			geodesMax = candidate.geode
		}

		fmt.Println(candidate)

		if b.geode.canBuild(candidate.ore, candidate.obsidian) {
			s := next(candidate)
			s.ore -= b.geode.ore
			s.obsidian -= b.geode.oth
			s.geodeR++
			states = append(states, s)
			geodeUnlocked = true
		}
		if b.obsidian.canBuild(candidate.ore, candidate.clay) {
			if !geodeUnlocked || (geodeUnlocked && candidate.geodeR > 0) {
				s := next(candidate)
				s.ore -= b.obsidian.ore
				s.clay -= b.obsidian.oth
				s.obsidianR++
				states = append(states, s)
				obsidianUnlocked = true
			}
		}
		if b.clay.canBuild(candidate.ore) {
			if !obsidianUnlocked || (obsidianUnlocked && candidate.obsidianR > 0) {
				s := next(candidate)
				s.ore -= b.clay.ore
				s.clayR++
				states = append(states, s)

				clayUnlocked = true
			}
		}
		if b.ore.canBuild(candidate.ore) {
			if !clayUnlocked || (clayUnlocked && candidate.clayR > 0) {
				s := next(candidate)
				s.ore -= b.ore.ore
				s.oreR++
				states = append(states, s)
			}
		}

		if !clayUnlocked || (clayUnlocked && candidate.clayR > 0) {
			if !obsidianUnlocked || (obsidianUnlocked && candidate.obsidianR > 0) {
				if !geodeUnlocked || (geodeUnlocked && candidate.geodeR > 0) {
					states = append(states, next(candidate))
				}
			}
		}
	}
	return geodesMax
}

func parseInput(r io.Reader) []blueprint {
	scanner := bufio.NewScanner(r)
	var blueprints []blueprint
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()

		coordStr := strings.FieldsFunc(line, func(r rune) bool {
			return r != '-' && !unicode.IsDigit(r)
		})
		id, _ := strconv.Atoi(coordStr[0])
		oreBuildPrice, _ := strconv.Atoi(coordStr[1])
		clayBuildPrice, _ := strconv.Atoi(coordStr[2])
		obsOreBuildPrice, _ := strconv.Atoi(coordStr[3])
		obsClayBuildPrice, _ := strconv.Atoi(coordStr[4])
		geoOreBuildPrice, _ := strconv.Atoi(coordStr[5])
		geoObsBuildPrice, _ := strconv.Atoi(coordStr[6])
		blueprints = append(blueprints, blueprint{
			id:       id,
			ore:      baseMaterial{id: ore, ore: oreBuildPrice},
			clay:     baseMaterial{id: clay, ore: clayBuildPrice},
			obsidian: advMaterial{id: obsidian, ore: obsOreBuildPrice, oth: obsClayBuildPrice},
			geode:    advMaterial{id: geode, ore: geoOreBuildPrice, oth: geoObsBuildPrice},
		})
	}
	return blueprints
}
