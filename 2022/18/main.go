package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	xyzs := parseInput(f)

	// part 1
	var b boulder
	for _, xyz := range xyzs {
		b.addRock(xyz)
	}

	fmt.Println("part 1: ", b.outSides)
	// part 2

	internal := internal(b.cubes)
	// fmt.Println("part 2: ", internal)
	fmt.Println("part 2: ", b.outSides-internal.outSides)
	fmt.Println("part 2: ", b.findExternal())
	// 2444 too low
}

type coord3D struct{ x, y, z int }

func (coord3D) sides() int { return 6 }

func parseInput(r io.Reader) []coord3D {
	scanner := bufio.NewScanner(r)
	var xyzs []coord3D
	for scanner.Scan() {
		line := scanner.Text()
		coordStr := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordStr[0])
		y, _ := strconv.Atoi(coordStr[1])
		z, _ := strconv.Atoi(coordStr[2])
		xyzs = append(xyzs, coord3D{x: x, y: y, z: z})
	}
	return xyzs
}

type boulder struct {
	outSides int
	cubes    []coord3D
}

func (b *boulder) addRock(c coord3D) int {
	sides := c.sides()
	for _, curr := range b.cubes {
		switch curr {
		case (coord3D{x: c.x + 1, y: c.y, z: c.z}),
			(coord3D{x: c.x - 1, y: c.y, z: c.z}),
			(coord3D{x: c.x, y: c.y + 1, z: c.z}),
			(coord3D{x: c.x, y: c.y - 1, z: c.z}),
			(coord3D{x: c.x, y: c.y, z: c.z + 1}),
			(coord3D{x: c.x, y: c.y, z: c.z - 1}):
			sides -= 2
		}
	}
	b.outSides += sides
	b.cubes = append(b.cubes, c)
	return b.outSides
}

func internal(cubes []coord3D) boulder {
	contains := func(xyzs []coord3D, xyz coord3D) bool {
		for i := range xyzs {
			if xyzs[i] != xyz {
				continue
			}
			return true
		}
		return false
	}

	var b boulder
	for x := 0; x <= 21; x++ {
		for y := 0; y <= 21; y++ {
			for z := 0; z <= 21; z++ {
				if contains(cubes, coord3D{x, y, z}) {
					continue
				}
				if !isInternal(cubes, coord3D{x, y, z}) {
					continue
				}
				b.addRock(coord3D{x, y, z})
			}
		}
	}
	return b
}

func isInternal(cubes []coord3D, c coord3D) bool {
	contains := func(xyzs []coord3D, xyz coord3D) bool {
		for i := range xyzs {
			if xyzs[i] != xyz {
				continue
			}
			return true
		}
		return false
	}

	var interiorCubeTest bool
	for i := 0; c.x+i <= 21; i++ {
		testedCube := coord3D{x: c.x + i, y: c.y, z: c.z}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	if !interiorCubeTest {
		return false
	}

	interiorCubeTest = false
	for i := 0; c.x-i >= 0; i++ {
		testedCube := coord3D{x: c.x - i, y: c.y, z: c.z}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	if !interiorCubeTest {
		return false
	}

	interiorCubeTest = false
	for i := 0; c.y+i <= 21; i++ {
		testedCube := coord3D{x: c.x, y: c.y + i, z: c.z}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	if !interiorCubeTest {
		return false
	}

	interiorCubeTest = false
	for i := 0; c.y-i >= 0; i++ {
		testedCube := coord3D{x: c.x, y: c.y - i, z: c.z}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	if !interiorCubeTest {
		return false
	}

	interiorCubeTest = false
	for i := 0; c.z+i <= 21; i++ {
		testedCube := coord3D{x: c.x, y: c.y, z: c.z + i}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	if !interiorCubeTest {
		return false
	}

	interiorCubeTest = false
	for i := 0; c.z-i >= 0; i++ {
		testedCube := coord3D{x: c.x, y: c.y, z: c.z - i}
		if !contains(cubes, testedCube) {
			continue
		}
		interiorCubeTest = true
		break
	}
	return interiorCubeTest
}

func (b *boulder) findExternal() int {
	var max coord3D
	for _, cube := range b.cubes {
		if cube.x > max.x {
			max.x = cube.x
		}
		if cube.y > max.y {
			max.y = cube.y
		}
		if cube.z > max.z {
			max.z = cube.z
		}
	}
	boulderCubes := make(map[coord3D]bool)
	for _, cube := range b.cubes {
		boulderCubes[cube] = true
	}
	return visit(coord3D{x: 0, y: 0, z: 0}, boulderCubes, make(map[coord3D]bool), max)

}

func visit(cube coord3D, boulderCubes map[coord3D]bool, exterior map[coord3D]bool, cubeSize coord3D) int {
	if exterior[cube] {
		return 0
	}
	if cube.x < -1 || cube.x > cubeSize.x+1 || cube.y < -1 || cube.y > cubeSize.y+1 || cube.z < -1 || cube.z > cubeSize.z+1 {
		return 0
	}
	if boulderCubes[cube] {
		return 1
	}
	exterior[cube] = true

	neighbors := func(c coord3D) []coord3D {
		return []coord3D{
			{x: c.x + 1, y: c.y, z: c.z},
			{x: c.x - 1, y: c.y, z: c.z},
			{x: c.x, y: c.y + 1, z: c.z},
			{x: c.x, y: c.y - 1, z: c.z},
			{x: c.x, y: c.y, z: c.z + 1},
			{x: c.x, y: c.y, z: c.z - 1},
		}
	}
	var count int
	for _, neighbor := range neighbors(cube) {
		count += visit(neighbor, boulderCubes, exterior, cubeSize)
	}
	return count

}
