package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var in string

func main() {
	lines := strings.Split(in, "\n")
	var c1, c2 []int
	for _, l := range lines {
		f := strings.Fields(l)
		n1, _ := strconv.Atoi(f[0])
		c1 = append(c1, n1)
		n2, _ := strconv.Atoi(f[1])
		c2 = append(c2, n2)
	}
	sort.Ints(c1)
	sort.Ints(c2)
	if len(c1) != len(c2) {
		panic("not expected")
	}
	var totaldistance int
	for i := range c1 {
		n := c1[i] - c2[i]
		if n < 0 {
			n = n * -1
		}
		totaldistance += n
	}
	fmt.Println(totaldistance)
	var similarityscore int
nextNumber:
	for _, n := range c1 {
		var hits int
		for _, m := range c2 {
			switch {
			case n > m:
				continue
			case n == m:
				hits++
				continue
			case n < m:
				similarityscore += n * hits
				continue nextNumber
			default:
				panic("not possible")
			}
		}
	}
	fmt.Println(similarityscore)
}
