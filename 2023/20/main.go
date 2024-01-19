package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed sample
var input string

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var modules []module
	for scanner.Scan() {
		line := scanner.Text()
		l := line[:strings.Index(line, " ")]
		r := line[strings.Index(line, ">")+1:]
		var ds []string
		for _, d := range strings.Split(r, ",") {
			ds = append(ds, strings.Trim(d, " "))
		}
		switch l[0] {
		case 'b':
			modules = append(modules, module{
				kind: l,
				name: l,
				dest: ds,
			})
			modules = append(modules, module{
				kind: "button",
				name: "button",
				dest: []string{l},
			})
		case '%':
			modules = append(modules, module{
				kind: "flip-flop",
				name: l[1:],
				dest: ds,
			})
		case '&':
			modules = append(modules, module{
				kind: "conjunction",
				name: l[1:],
				dest: ds,
			})
		default:
		}
	}
	log.Print(modules)
}

type pulse bool

const (
	high = pulse(true)
	low  = pulse(false)
)

func (p pulse) String() string {
	switch p {
	case low:
		return "low"
	default:
		return "high"
	}
}

type module struct {
	kind, name string
	p          pulse
	dest       []string
}

func (m module) String() string { return fmt.Sprintf("{%s %s %s %v}", m.kind, m.name, m.p, m.dest) }

type moduler interface {
	process(pulse)
}
