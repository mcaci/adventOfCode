package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main_readconc() {
	f, err := os.Open("day1")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	var lastLine bool
	var lineCount int
	for !lastLine {
		s, err := r.ReadString('\n')
		switch err {
		case nil:
			s = s[:len(s)-1]
		case io.EOF:
			lastLine = true
		default:
			log.Fatal(err)
		}
		fmt.Println(s)
		lineCount++
	}
	log.Println(lineCount)
}
