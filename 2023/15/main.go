package main

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed sample
var input string

func main() {
	f := strings.Split(input, ",")
	var score int
	for _, s := range f {
		score += holidayASCIIStringHelper(s)
	}
	log.Println(score)
}

func holidayASCIIStringHelper(s string) int {
	var hash int
	for i := range s {
		hash += int(s[i])
		hash *= 17
		hash %= 256
	}
	return hash
}
