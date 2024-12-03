package main

import (
	_ "embed"
	"flag"
	"log"
	"strconv"
	"strings"
)

//go:embed input
var in string

func main() {
	p2 := flag.Bool("p2", false, "part 2")
	flag.Parse()
	var skip bool
	var sum int
	for i := range in {
		if *p2 {
			switch {
			case do(in, i):
				skip = false
			case dont(in, i):
				skip = true
			}
			if skip {
				continue
			}
		}

		if !mul(in, i) {
			continue
		}
		start := i + 4
		f := strings.Split(operands(in, start), ",")
		if len(f) != 2 {
			continue
		}
		a, err := strconv.Atoi(f[0])
		if err != nil {
			continue
		}
		b, err := strconv.Atoi(f[1])
		if err != nil {
			continue
		}
		sum += a * b
	}
	log.Println(sum)
}

func mul(in string, i int) bool {
	if in[i] != 'm' {
		return false
	}
	if in[i+1] != 'u' {
		return false
	}
	if in[i+2] != 'l' {
		return false
	}
	if in[i+3] != '(' {
		return false
	}
	return true
}

func operands(in string, start int) string {
	for i := range in[start:] {
		if in[start+i] != ')' {
			continue
		}
		return in[start : start+i]
	}
	return ""
}

func do(in string, i int) bool {
	if in[i] != 'd' {
		return false
	}
	if in[i+1] != 'o' {
		return false
	}
	if in[i+2] != '(' {
		return false
	}
	if in[i+3] != ')' {
		return false
	}
	return true
}

func dont(in string, i int) bool {
	if in[i] != 'd' {
		return false
	}
	if in[i+1] != 'o' {
		return false
	}
	if in[i+2] != 'n' {
		return false
	}
	if in[i+3] != '\'' {
		return false
	}
	if in[i+4] != 't' {
		return false
	}
	if in[i+5] != '(' {
		return false
	}
	if in[i+6] != ')' {
		return false
	}
	return true
}
