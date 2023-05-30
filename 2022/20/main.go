package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := parseNList(f)
	actualBuf := make([]*nList, len(buf))
	copy(actualBuf, buf)
	// part 1
	fmt.Println("part 1: ", mix(buf, 1))

	// part 2
	const key = 811589153
	for i := range actualBuf {
		actualBuf[i].n *= key
	}
	fmt.Println("part 2: ", mix(actualBuf, 10))
}

func mix(buf []*nList, nTimes int) int {
	// fmt.Println(buf)
	for i := 0; i < nTimes; i++ {
		// fmt.Println("time", i+1)
		for k := range buf {
			buf[k].moved = false
		}
		var idCurr int
		for idCurr < len(buf) {
			var e *nList = buf[idCurr]
			var iCurr int
			for i := range buf {
				// fmt.Println(iCurr, buf[i])
				if buf[i].id != idCurr {
					continue
				}
				e = buf[i]
				iCurr = i
				break
			}
			if e.moved {
				idCurr++
				continue
			}
			// fmt.Println("moving ", e)
			if e.n == 0 {
				e.moved = true
				iCurr++
				continue
			}
			iNext := iCurr + e.n
			if iNext <= 0 {
				iNext %= len(buf) - 1
				iNext += len(buf) - 1
			}
			if iNext >= len(buf) {
				iNext %= len(buf) - 1
			}
			if iNext > iCurr {
				for i := iCurr; i < iNext; i++ {
					buf[i] = buf[i+1]
				}
			}
			if iNext < iCurr {
				for i := iCurr; i > iNext; i-- {
					buf[i] = buf[i-1]
				}
			}
			buf[iNext] = e
			// fmt.Println(buf)
			e.moved = true
		}
	}

	var zeroPos int
	for i, e := range buf {
		if e.n != 0 {
			continue
		}
		zeroPos = i
	}
	x, y, z := 1000+zeroPos, 2000+zeroPos, 3000+zeroPos

	fmt.Println(buf[x%(len(buf))], buf[y%(len(buf))], buf[z%(len(buf))])
	return buf[x%(len(buf))].n + buf[y%(len(buf))].n + buf[z%(len(buf))].n
}

type nList struct {
	id    int
	n     int
	moved bool
}

func (n nList) String() string { return fmt.Sprintf("(%d,%d,%t)", n.id, n.n, n.moved) }

func parseNList(r io.Reader) []*nList {
	scanner := bufio.NewScanner(r)
	var buffer []*nList
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		buffer = append(buffer, &nList{id: i, n: n})
	}
	return buffer
}
