package main

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

type oasisReading []int

func nextValue(r oasisReading) int {
	nrs := []oasisReading{r}
	lr := r
	for !zero(lr) {
		var nr oasisReading
		for i := range lr {
			if i == 0 {
				continue
			}
			nr = append(nr, lr[i]-lr[i-1])
		}
		nrs = append(nrs, nr)
		lr = nr
	}
	// log.Print(nrs)
	nrs[len(nrs)-1] = append(nrs[len(nrs)-1], 0)
	for i := len(nrs) - 2; i >= 0; i-- {
		lastLower := nrs[i+1][len(nrs[i+1])-1]
		nrs[i] = append(nrs[i], nrs[i][len(nrs[i])-1]+lastLower)
	}
	// log.Print(nrs)
	return nrs[0][len(nrs[0])-1]
}

func previousValue(r oasisReading) int {
	nrs := []oasisReading{r}
	lr := r
	for !zero(lr) {
		var nr oasisReading
		for i := range lr {
			if i == 0 {
				continue
			}
			nr = append(nr, lr[i]-lr[i-1])
		}
		nrs = append(nrs, nr)
		lr = nr
	}
	// log.Print(nrs)
	nrs[len(nrs)-1] = append(oasisReading{0}, nrs[len(nrs)-1]...)
	for i := len(nrs) - 2; i >= 0; i-- {
		lastLower := nrs[i+1][0]
		nrs[i] = append(oasisReading{nrs[i][0] - lastLower}, nrs[i]...)
	}
	// log.Print(nrs)
	return nrs[0][0]
}

func zero(r oasisReading) bool {
	for i := range r {
		if r[i] != 0 {
			return false
		}
	}
	return true
}

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var ors []oasisReading
	for scanner.Scan() {
		line := scanner.Text()
		var or oasisReading
		f := strings.Fields(line)
		// log.Print(f)
		for i := range f {
			n, _ := strconv.Atoi(f[i])
			or = append(or, n)
		}
		ors = append(ors, or)
	}
	log.Print(run(ors, nextValue))
	log.Print(run(ors, previousValue))
}

func run(ors []oasisReading, nextValue func(oasisReading) int) int {
	var sum int
	for _, or := range ors {
		sum += nextValue(or)
	}
	return sum
}
