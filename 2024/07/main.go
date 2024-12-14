package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed sample
var in string

func main() {
	var ops []operation
	for _, line := range strings.Split(in, "\n") {
		i := strings.Index(line, ":")
		r, op := line[:i], line[i+2:]
		nr, _ := strconv.Atoi(r)
		var nop []int
		for _, sop := range strings.Fields(op) {
			n, _ := strconv.Atoi(sop)
			nop = append(nop, n)
		}
		ops = append(ops, operation{result: nr, operands: nop})
	}
	log.Print(ops)
}

type operation struct {
	result   int
	operands []int
}
