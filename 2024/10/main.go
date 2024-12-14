package main

import (
	"bytes"
	_ "embed"
	"log"
)

//go:embed sample
var in []byte

func main() {
	mat := bytes.Split(in, []byte("\n"))
	for i := range mat {
		log.Print(mat[i])
	}
}
