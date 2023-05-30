package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type monkey[T constraints.Integer] struct {
	items          []T
	operation      func(T) T
	test           func(T) bool
	trueM, falseM  int
	inspectedItems T
}

func parse[T constraints.Integer](line string, monkeys []*monkey[T], monkeyIndex byte) byte {
	infos := strings.Fields(line)
	switch infos[0] {
	case "Monkey":
		monkeyIndex = infos[1][0] - 48
	case "Starting":
		newLine := line[strings.Index(line, ":")+1:]
		items := strings.Fields(newLine)
		for i := range items {
			item, _ := strconv.Atoi(strings.Trim(items[i], ","))
			monkeys[monkeyIndex].items = append(monkeys[monkeyIndex].items, T(item))
		}
	case "Operation:":
		op := infos[len(infos)-2]
		n, err := strconv.Atoi(infos[len(infos)-1])
		monkeys[monkeyIndex].operation = func(i T) T {
			switch op {
			case "+":
				return i + T(n)
			case "*":
				if err != nil {
					return i * i
				}
				return i * T(n)
			}
			return 0
		}
	case "Test:":
		n, _ := strconv.Atoi(infos[len(infos)-1])
		monkeys[monkeyIndex].test = func(i T) bool { return i%T(n) == 0 }
	case "If":
		tf := infos[1][:len(infos[1])-1]
		b, _ := strconv.ParseBool(tf)
		n, _ := strconv.Atoi(infos[len(infos)-1])
		switch b {
		case true:
			monkeys[monkeyIndex].trueM = n
		case false:
			monkeys[monkeyIndex].falseM = n
		}
	}
	return monkeyIndex
}

func relaxP1[T constraints.Integer](i T) T { return i / 3 }

// modulo 2 * 7 * 13 * 3 * 19 * 5 * 17 * 11 maintains the divisibility for all these prime numbers
func relaxP2[T int | uint64](i T) T { return i % T(2*7*13*3*19*5*17*11) }

func main() {
	f, err := os.Open("./11/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	monkeys := make([]*monkey[uint64], 8)
	for i := range monkeys {
		monkeys[i] = &monkey[uint64]{}
	}
	s := bufio.NewScanner(f)
	var monkeyIndex byte
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		monkeyIndex = parse(line, monkeys, monkeyIndex)
	}
	// part 1
	// const maxRound = 20
	// part 2
	const maxRound = 10000
	for round := 0; round < maxRound; round++ {
		for i, m := range monkeys {
			for len(m.items) > 0 {
				item := m.items[0]
				m.inspectedItems++
				item = m.operation(item)
				// item = relaxP1(item)
				item = relaxP2(item)
				throwTo := m.falseM
				if m.test(item) {
					throwTo = m.trueM
				}
				monkeys[throwTo].items = append(monkeys[throwTo].items, item)
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
	}
	for _, m := range monkeys {
		log.Println(m)
	}
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspectedItems < monkeys[j].inspectedItems })
	fmt.Println(monkeys[len(monkeys)-1].inspectedItems * monkeys[len(monkeys)-2].inspectedItems)
}
