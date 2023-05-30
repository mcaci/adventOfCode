package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main_day17part1() {
	day17Part1()
}

func day17Part1() {
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
	var allowedTargetYs []float64
	for i := bytesToInt(fields[3]); i >= bytesToInt(fields[2]); i-- {
		allowedTargetYs = append(allowedTargetYs, float64(i))
	}
	var v0s []float64
	for _, y := range allowedTargetYs {
		v0 := math.Sqrt(-y)
		if v0-math.Round(v0) != 0 {
			continue
		}
		v0s = append(v0s, v0)
	}
	fmt.Println(len(allowedTargetYs), allowedTargetYs)
	fmt.Println(v0s)
	v := 9.0
	for y := float64(0); y+v > -129; {
		y += v
		fmt.Print(y, " ")
		v--
	}
	fmt.Println()
	v = 10.0
	for y := float64(0); y+v > -129; {
		y += v
		fmt.Print(y, " ")
		v--
	}
	fmt.Println()
	v = 11.0
	for y := float64(0); y+v > -129; {
		y += v
		fmt.Print(y, " ")
		v--
	}
	fmt.Println()
	v = 128.0
	for y := float64(0); y+v >= -129; {
		y += v
		fmt.Print(y, " ")
		v--
	}
	fmt.Println()
}

func bytesToInt(b []byte) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return n
}
