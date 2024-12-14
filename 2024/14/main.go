package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed sample
var in string

func main() {
	var rs []robot
	for _, line := range strings.Split(in, "\n") {
		ix := strings.Index(line, "=")
		ixend := strings.Index(line, " ")
		iy := strings.LastIndex(line, "=")
		p := line[ix+1 : ixend]
		sp := strings.Split(p, ",")
		spx, _ := strconv.Atoi(sp[0])
		spy, _ := strconv.Atoi(sp[1])
		v := line[iy+1:]
		sv := strings.Split(v, ",")
		svx, _ := strconv.Atoi(sv[0])
		svy, _ := strconv.Atoi(sv[1])
		rs = append(rs, robot{p: coord{spx, spy}, v: coord{svx, svy}})
	}
	log.Print(rs)
}

type coord struct{ x, y int }

type robot struct{ p, v coord }
