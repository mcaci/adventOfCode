package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	b, _ := os.ReadFile("input")
	l := bytes.Split(b, []byte{'\n'})
	var cals []int
	var curr int
	for i := range l {
		if len(l[i]) != 0 {
			n, _ := strconv.Atoi(string(l[i]))
			curr += n
			continue
		}
		cals = append(cals, curr)
		curr = 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	fmt.Println(cals[0]+ cals[1]+ cals[2])
}
