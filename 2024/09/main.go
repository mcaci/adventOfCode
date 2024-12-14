package main

import (
	_ "embed"
	"log"
)

//go:embed sample
var in string

func main() {
	log.Print(in)
}
