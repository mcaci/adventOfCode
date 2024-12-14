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
	var stones []int
	for _, sn := range strings.Fields(in) {
		s, _ := strconv.Atoi(sn)
		stones = append(stones, s)
	}
	log.Print(stones)
}
