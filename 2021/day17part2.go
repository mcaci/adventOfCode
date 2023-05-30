package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	day17Part2()
}

type coord struct{ x, y float64 }

func sim(v, ul, br coord) bool {
	var start coord
	vx, vy := v.x, v.y
	for {
		start.x += vx
		start.y += vy
		if start.x >= ul.x && start.x <= br.x &&
			start.y <= ul.y && start.y >= br.y {
			return true
		}
		if start.x > br.x {
			return false
		}
		if start.x < ul.x && vx == 0 {
			return false
		}
		if start.y < br.y {
			return false
		}
		if vx > 0 {
			vx--
		}
		vy--
	}
	return false
}

func day17Part2() {
	l, err := os.ReadFile("day17")
	if err != nil {
		log.Fatal(err)
	}
	l = bytes.Map(func(r rune) rune {
		if !(unicode.IsDigit(r) || r == '-') {
			return ' '
		}
		return r
	}, l)
	fields := bytes.Fields(l)
	upLeftCoord := coord{x: byteToFloat64(fields[0]), y: byteToFloat64(fields[3])}
	bottomRightCoord := coord{x: byteToFloat64(fields[1]), y: byteToFloat64(fields[2])}
	var count int
	for x := 0.0; x <= 171.0; x++ {
		for y := 128.0; y >= -129.0; y-- {
			vect := coord{x, y}
			ok := sim(vect, upLeftCoord, bottomRightCoord)
			fmt.Printf("For vect %v simulation is %t\n", vect, ok)
			if !ok {
				continue
			}
			count++
		}
	}
	fmt.Println(count)
}

type stepSum struct {
	step, sum int
}

func day17Part2Fail() {
	l, err := os.ReadFile("day17")
	if err != nil {
		log.Fatal(err)
	}
	l = bytes.Map(func(r rune) rune {
		if !(unicode.IsDigit(r) || r == '-') {
			return ' '
		}
		return r
	}, l)
	fields := bytes.Fields(l)

	xStepsMap := make(map[int][]stepSum)
	// for i := 1; i <= 20; i++ {
	for i := 1; i <= int(byteToFloat64(fields[1])); i++ {
		sum := i
		if sum >= int(byteToFloat64(fields[0])) {
			xStepsMap[i] = append(xStepsMap[i], stepSum{1, sum})
			continue
		}
		j := i + 1
		// for sum+j <= 20 {
		for sum+j <= int(byteToFloat64(fields[1])) {
			sum += j
			if sum >= int(byteToFloat64(fields[0])) {
				xStepsMap[i] = append(xStepsMap[i], stepSum{j - i + 1, sum})
			}
			j++
		}
	}
	yStepsMap := make(map[int][]stepSum)
	for i := -1; i >= int(byteToFloat64(fields[2])); i-- {
		sum := i
		if sum <= int(byteToFloat64(fields[3])) {
			yStepsMap[i] = append(yStepsMap[i], stepSum{1, sum})
			continue
		}
		j := i - 1
		for sum+j >= int(byteToFloat64(fields[2])) {
			sum += j
			if sum <= int(byteToFloat64(fields[3])) {
				yStepsMap[i] = append(yStepsMap[i], stepSum{i - j + 1, sum})
			}
			j--
		}
	}
	{
		yUpStepsMap := make(map[int][]stepSum)
		for k, v := range yStepsMap {
			kNew := -(k + 1)
			vNew := make([]stepSum, len(v))
			for i, n := range v {
				// 4 + (i-1)*2 -1 is the number of steps needed for the object to go to x=0 (same height as start at step 0)
				vNew[i] = stepSum{n.step + (4 + (kNew-1)*2 - 1), n.sum}
			}
			yUpStepsMap[kNew] = vNew
		}
		for k, v := range yUpStepsMap {
			yStepsMap[k] = append(yStepsMap[k], v...)
		}
	}
	fmt.Println("steps:", xStepsMap)
	fmt.Println("steps:", yStepsMap)

	var count int
	for _, vx := range xStepsMap {
		for i := range vx {
			for _, vy := range yStepsMap {
				for j := range vy {
					if vx[i].step > vy[j].step {
						continue
					}
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func byteToFloat64(b []byte) float64 {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return float64(n)
}
