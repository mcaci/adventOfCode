package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum int64
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		sum += snafuToDec(line)
		// fmt.Println(line, snafuToDec(line), decToSnafu(snafuToDec(line)))
	}

	// part 1
	fmt.Println("part 1:", decToSnafu(sum))

	// part 2
	fmt.Println("part 2:", 0)
}

func snafuToDec(snafu string) int64 {
	s2d := map[rune]int64{'2': 2, '1': 1, 0: '0', '-': -1, '=': -2}
	l := len(snafu)
	var dec int64
	for i, d := range snafu {
		dec += s2d[d] * int64(math.Pow(5, float64(l-(i+1))))
	}
	return dec
}

func decToSnafu(dec int64) string {
	base5 := strconv.FormatInt(dec, 5)
	if !strings.ContainsAny(base5, "34") {
		return base5
	}
	snafuDigits := make([]byte, len(base5))
	copy(snafuDigits, []byte(base5))
	for i := len(snafuDigits) - 1; i >= 0; i-- {
		// digit i...
		switch snafuDigits[i] {
		case '3':
			snafuDigits[i] = '='
		case '4':
			snafuDigits[i] = '-'
		case '5':
			snafuDigits[i] = '0'
		default:
			continue
		}
		// ...and carry on
		if i > 0 {
			snafuDigits[i-1]++
			continue
		}
		snafuDigits = append([]byte{'1'}, snafuDigits...)
	}
	return string(snafuDigits)
}
