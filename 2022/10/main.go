package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reg struct {
	cycle   int
	xDuring int
	xAfter  int
}

type addx struct {
	c, n int
}

func (a *addx) apply(r *reg) bool {
	r.xDuring = r.xAfter
	r.cycle++
	if a.c == 0 {
		a.c++
		return false
	}
	r.xAfter += a.n
	return true
}

type noop struct{}

func (noop) apply(r *reg) bool { r.xDuring = r.xAfter; r.cycle++; return true }

func newCommand(line string) interface{ apply(*reg) bool } {
	f := strings.Fields(line)
	switch f[0] {
	case "noop":
		return noop{}
	case "addx":
		n, _ := strconv.Atoi(f[1])
		return &addx{n: n}
	default:
		return nil
	}
}

func main() {
	f, err := os.Open("./10/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := reg{xAfter: 1, xDuring: 1}
	regState := []reg{r}
	display := []byte{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		cmd := newCommand(line)
		for !cmd.apply(&r) {
			regState = append(regState, r)
			displayPos := len(display) % 40
			centerPixel := r.xDuring
			pixel := byte(' ')
			switch displayPos {
			case centerPixel - 1, centerPixel, centerPixel + 1:
				pixel = '#'
			}
			display = append(display, pixel)
		}
		regState = append(regState, r)
		displayPos := len(display) % 40
		centerPixel := r.xDuring
		pixel := byte(' ')
		switch displayPos {
		case centerPixel - 1, centerPixel, centerPixel + 1:
			pixel = '#'
		}
		display = append(display, pixel)
	}
	// part 1
	var sum int
	c := []int{20, 60, 100, 140, 180, 220}
	for _, n := range c {
		log.Println(n, regState[n], n*regState[n].xDuring)
		sum += n * regState[n+1].xDuring
	}
	fmt.Println(sum)
	// part 2
	for i := 0; i < 240; i += 40 {
		fmt.Println(string(display[i : i+40]))
	}
}
