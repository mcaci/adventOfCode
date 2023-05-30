package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count = p2(scanner)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func p1(scanner *bufio.Scanner) int {
	aContainsB := func(aS, aE, bS, bE int) bool {
		if aS > bS {
			return false
		}
		if aE < bE {
			return false
		}
		return true
	}
	var count int

	for scanner.Scan() {
		line := scanner.Text()
		info := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '-' })
		aS, _ := strconv.Atoi(info[0])
		aE, _ := strconv.Atoi(info[1])
		bS, _ := strconv.Atoi(info[2])
		bE, _ := strconv.Atoi(info[3])
		if aContainsB(aS, aE, bS, bE) {
			count++
			continue
		}
		if aContainsB(bS, bE, aS, aE) {
			count++
			continue
		}
	}
	return count
}

func p2(scanner *bufio.Scanner) int {
	aContainsBsExtreme := func(aS, aE, bx int) bool {
		if aS > bx {
			return false
		}
		if aE < bx {
			return false
		}
		return true
	}
	var count int

	for scanner.Scan() {
		line := scanner.Text()
		info := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '-' })
		aS, _ := strconv.Atoi(info[0])
		aE, _ := strconv.Atoi(info[1])
		bS, _ := strconv.Atoi(info[2])
		bE, _ := strconv.Atoi(info[3])
		if aContainsBsExtreme(aS, aE, bS) {
			count++
			continue
		}
		if aContainsBsExtreme(aS, aE, bE) {
			count++
			continue
		}
		if aContainsBsExtreme(bS, bE, aS) {
			count++
			continue
		}
		if aContainsBsExtreme(bS, bE, aE) {
			count++
			continue
		}
	}
	return count
}