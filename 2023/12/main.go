package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var recs []record
	for scanner.Scan() {
		line := scanner.Text()
		recs = append(recs, record(line))
	}

	var sum int
	for i := range recs {
		br := brokenBlocks(recs[i])
		seq := sequence(recs[i])
		r := replacements(recs[i])
		for j := range r {
			rep := replace(seq, r[j])
			ok := valid(rep, br)
			if ok {
				log.Print(string(r[j]))
			}
			if !ok {
				continue
			}
			sum++
		}
	}
	log.Println(sum)
}

type record string

func brokenBlocks(r record) []int {
	var recs []int
	i := strings.IndexByte(string(r), ' ')
	f := strings.Split(string(r)[i+1:], ",")
	for i := range f {
		n, _ := strconv.Atoi(f[i])
		recs = append(recs, n)
	}
	return recs
}

func sequence(r record) []byte {
	var recs []byte
	i := strings.IndexByte(string(r), ' ')
	for _, b := range string(r)[:i] {
		recs = append(recs, byte(b))
	}
	return recs
}

func replacements(r record) [][]byte {
	var count float64
	for i := range r {
		if r[i] != '?' {
			continue
		}
		count++
	}
	recs := make([][]byte, int(math.Exp2(count)))
	for i := range recs {
		n := strconv.FormatInt(int64(i), 2)
		if len(n) < int(count) {
			n = string(append(bytes.Repeat([]byte{'0'}, int(count)-len(n)), n...))
		}
		recs[i] = bytes.Map(func(r rune) rune {
			switch r {
			case '0':
				return '#'
			default:
				return '.'
			}
		}, []byte(n))

	}
	fmt.Println(recs)
	return recs
}

func replace(seq []byte, repl []byte) []byte {
	var replaced []byte
	var replId int
	for i := range seq {
		if seq[i] != '?' {
			replaced = append(replaced, seq[i])
			continue
		}
		replaced = append(replaced, repl[replId])
		replId++
	}
	return replaced
}

func valid(seq []byte, broken []int) bool {
	const (
		anyDots       = "\\.*"
		atLeastOneDot = "\\.+"
	)
	nBroken := func(i int) string { return fmt.Sprintf("#{%d}", i) }
	rex := "^" + anyDots
	for i := range broken {
		rex += nBroken(broken[i])
		switch i {
		case len(broken) - 1:
			rex += anyDots + "$"
		default:
			rex += atLeastOneDot
		}
	}
	r, err := regexp.Compile(rex)
	if err != nil {
		log.Panic("nope")
	}
	ok := r.Match(seq)
	if ok {
		log.Print(string(seq), " ooo ", rex, " ooo ", ok)
	}
	return ok
}
