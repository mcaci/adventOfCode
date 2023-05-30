package main

import (
	"fmt"
	"os"
)

func main() {
	b, _ := os.ReadFile("input")
	// Part 1
	// const start = 4
	// Part 2
	const start = 14
nextSeq:
	for i := start; i <= len(b); i++ {
		m := make(map[byte]bool)
		for _, c := range b[i-start : i] {
			if !m[c] {
				m[c] = true
				continue
			}
			fmt.Println(i, string(b[i]), string(b[i-start:i]))
			continue nextSeq
		}
		fmt.Println(i, string(b[i]), string(b[i-start:i]))
		fmt.Println(string(b[:i]))
		break
	}
}
