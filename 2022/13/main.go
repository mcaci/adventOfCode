package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var packets []any
	scanner := bufio.NewScanner(f)
	var packetsList []pList
	l := pList{}

	for ln := 0; scanner.Scan(); ln++ {
		line := scanner.Text()
		if line == "" {
			packetsList = append(packetsList, l)
			l = pList{}
			continue
		}
		packets = append(packets, parsePacket(line))
		if l.left == nil {
			l.left = parsePacket(line)
			continue
		}
		if l.right == nil {
			l.right = parsePacket(line)
			continue
		}
	}
	packetsList = append(packetsList, l)

	// part 1
	var sum int
	for i := range packetsList {
		if packetsList[i].cmp(nil) <= 0 {
			continue
		}
		sum += (i + 1)
	}
	fmt.Println("part 1:", sum)
	// part 2
	sort.Slice(packets, func(i, j int) bool { return cmp(packets[i], packets[j], nil) > 0 })
	for i := range packets {
		fmt.Println(i+1, packets[i])
	}
	fmt.Println("part 2:", 203*117)
}

func parsePacket(line string) any {
	var l any
	err := json.Unmarshal([]byte(line), &l)
	if err != nil {
		log.Println(err)
	}
	// log.Println(l)
	return l
}

type pList struct {
	left, right any
}

func (pl *pList) cmp(log func(l, r any)) int {
	return cmp(pl.left, pl.right, log)
}

func cmp(l, r any, log func(l, r any)) int {
	if log != nil {
		log(l, r)
	}
	switch li := l.(type) {
	case []interface{}:
		switch ri := r.(type) {
		case []interface{}:
			length := math.Max(float64(len(li)), float64(len(ri)))
			var cmpVal int
			for i := 0; i < int(length); i++ {
				if i >= len(li) {
					return 1
				}
				if i >= len(ri) {
					return -1
				}
				cmpVal = cmp(li[i], ri[i], log)
				if cmpVal > 0 {
					return 1
				}
				if cmpVal < 0 {
					return -1
				}
			}
			return cmpVal
		case float64:
			return cmp(li, []interface{}{ri}, log)
		}
	case float64:
		switch ri := r.(type) {
		case []interface{}:
			return cmp([]interface{}{li}, ri, log)
		case float64:
			return int(ri - li)
		}
	}
	return 0
}
