package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	DARK_PIXEL  = '.'
	LIGHT_PIXEL = '#'
	INPUT_SIZE  = 100
)

func main() {
	day20(2)
	day20(50)
}

func day20(l int) {
	f, err := os.Open("day20")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iea, img := readInput(bufio.NewReader(f), 3*l+INPUT_SIZE)
	var out floorMap
	for i := 0; i < l; i++ {
		out = img.convert(iea)
		img = out
	}
	fmt.Println(out)
	var count int
	for i := range out {
		for j := range out[i] {
			if out[i][j] != LIGHT_PIXEL {
				continue
			}
			count++
		}
	}
	fmt.Println(count)
}

func readInput(r *bufio.Reader, size int) ([]byte, floorMap) {
	x := (size-INPUT_SIZE)/2 - 1
	var afterFirst, last bool
	var imgEnhAlgo []byte
	img := newFloorMap(size)
	for !last {
		l, err := r.ReadBytes('\n')
		switch err {
		case nil:
			l = l[:len(l)-1]
		case io.EOF:
			last = true
		default:
			log.Fatal(err)
		}
		if !afterFirst {
			imgEnhAlgo = l
			afterFirst = true
			continue
		}
		y := (size-INPUT_SIZE)/2 - 1
		for i := range l {
			y++
			if l[i] != LIGHT_PIXEL {
				continue
			}
			img[x][y] = LIGHT_PIXEL

		}
		x++
	}
	return imgEnhAlgo, img
}

type floorMap [][]byte

func newFloorMap(l int) floorMap {
	if l == 0 {
		l = 106 // I said so, let's see if it gets to errors
	}
	img := make([][]byte, l)
	for i := range img {
		img[i] = make([]byte, l)
		for j := range img[i] {
			img[i][j] = DARK_PIXEL
		}
	}
	return img
}

func (fm floorMap) convert(iea []byte) floorMap {
	out := newFloorMap(len(fm))
	for i := range fm {
		for j := range fm[i] {
			n := neighboursIndex(fm, i, j)
			out[i][j] = iea[n]
		}
	}
	return out
}

func capD(i int) int {
	if i <= 0 {
		return 0
	}
	return i - 1
}
func capI(i, max int) int {
	if i >= max {
		return max
	}
	return i + 1
}

func neighboursIndex(fm floorMap, i, j int) int64 {
	lim := len(fm) - 1
	tl, tc, tr := fm[capD(i)][capD(j)], fm[capD(i)][j], fm[capD(i)][capI(j, lim)]
	ml, mc, mr := fm[i][capD(j)], fm[i][j], fm[i][capI(j, lim)]
	bl, bc, br := fm[capI(i, lim)][capD(j)], fm[capI(i, lim)][j], fm[capI(i, lim)][capI(j, lim)]
	s := string([]byte{
		tl, tc, tr,
		ml, mc, mr,
		bl, bc, br,
	})
	r := strings.NewReplacer("#", "1", ".", "0")
	s = r.Replace(s)
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func (fm floorMap) String() string {
	var lines []string
	for i := range fm {
		lines = append(lines, string(fm[i]))
	}
	return strings.Join(lines, "\n")
}
