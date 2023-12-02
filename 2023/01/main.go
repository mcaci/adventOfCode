package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	b, _ := os.ReadFile("input")
	lines := bytes.Split(b, []byte{'\n'})
	var sum int
	for i := range lines {
		f := digitPart2(lines[i])
		l := lastDigitPart2(lines[i])
		n := fmt.Sprintf("%d%d", f, l)
		// fmt.Printf("%q %d %d %q\n", lines[i], f, l, n)
		s, _ := strconv.Atoi(n)
		sum += s
	}
	fmt.Println(sum)
}

func digitPart1(b []byte) int     { return int(b[bytes.IndexFunc(b, unicode.IsDigit)] - 48) }
func lastDigitPart1(b []byte) int { return int(b[bytes.LastIndexFunc(b, unicode.IsDigit)] - 48) }

func digitPart2(b []byte) int {
	seq := make([]int, 10)
	seq[0] = bytes.IndexFunc(b, unicode.IsDigit)
	seq[1] = bytes.Index(b, []byte("one"))
	seq[2] = bytes.Index(b, []byte("two"))
	seq[3] = bytes.Index(b, []byte("three"))
	seq[4] = bytes.Index(b, []byte("four"))
	seq[5] = bytes.Index(b, []byte("five"))
	seq[6] = bytes.Index(b, []byte("six"))
	seq[7] = bytes.Index(b, []byte("seven"))
	seq[8] = bytes.Index(b, []byte("eight"))
	seq[9] = bytes.Index(b, []byte("nine"))
	// fmt.Println(seq)

	id, min := 0, math.MaxInt
	for i := range seq {
		if seq[i] > min || seq[i] == -1 {
			continue
		}
		id = i
		min = seq[i]
	}
	if id == 0 {
		return int(b[seq[0]] - 48)
	}
	return id
}
func lastDigitPart2(b []byte) int {
	seq := make([]int, 10)
	seq[0] = bytes.LastIndexFunc(b, unicode.IsDigit)
	seq[1] = bytes.LastIndex(b, []byte("one"))
	seq[2] = bytes.LastIndex(b, []byte("two"))
	seq[3] = bytes.LastIndex(b, []byte("three"))
	seq[4] = bytes.LastIndex(b, []byte("four"))
	seq[5] = bytes.LastIndex(b, []byte("five"))
	seq[6] = bytes.LastIndex(b, []byte("six"))
	seq[7] = bytes.LastIndex(b, []byte("seven"))
	seq[8] = bytes.LastIndex(b, []byte("eight"))
	seq[9] = bytes.LastIndex(b, []byte("nine"))

	id, max := 0, 0
	for i := range seq {
		if seq[i] < max || seq[i] == -1 {
			continue
		}
		id = i
		max = seq[i]
	}
	if id == 0 {
		return int(b[seq[0]] - 48)
	}
	return id
}
