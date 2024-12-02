package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input
var in string

func main() {
	lines := strings.Split(in, "\n")
	var safeCountP1, safeCountP2 int
	for _, line := range lines {
		var report []int
		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, n)
		}
		if isSafe(report) {
			safeCountP1++
			safeCountP2++
			continue
		}
		if problemDampener(report) {
			safeCountP2++
		}
	}
	log.Println(safeCountP1)
	log.Println(safeCountP2)
}

func isSafe(report []int) bool {
	dir := report[0] < report[1]
	f1 := func(a, b int) bool {
		if dir {
			return a < b
		}
		return a > b
	}
	f2 := func(a, b int) bool {
		switch a - b {
		case -3, -2, -1, 1, 2, 3:
			return true
		default:
			return false
		}
	}
	for i := 0; i < len(report)-1; i++ {
		if f1(report[i], report[i+1]) && f2(report[i], report[i+1]) {
			continue
		}
		return false
	}
	return true
}

func problemDampener(report []int) bool {
	for i := range report {
		var tempReport []int
		for j := range report {
			if i == j {
				continue
			}
			tempReport = append(tempReport, report[j])
		}
		if !isSafe(tempReport) {
			continue
		}
		return true
	}
	return false
}
