package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"sort"

	"golang.org/x/image/draw"
)

func main() {
	f, err := os.Open("sample")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	field, l, h := parse(f)
	print(field, l, h)
	writePNG(fieldMap(field, l, h), "field.png", l, h)

	var caveImgSeq []*image.Paletted
	addToGif := gifAdder(&caveImgSeq, l, h)

	start := fieldPoint(field, l, h, func(c *cell) bool { return c.i > 0 })
	exit := fieldPoint(field, l, h, func(c *cell) bool { return c.i < h-1 })
	t := traverse(field, l, h, start, exit, addToGif)
	// part 1
	fmt.Println("part 1:", t)

	// part 2
	t1 := traverse(field, l, h, exit, start, addToGif)
	t2 := traverse(field, l, h, start, exit, addToGif)
	fmt.Println("part 2:", t+t1+t2)
	// gif
	writeGIF(caveImgSeq, "field.gif")
}

func traverse(field []*cell, l, h int, start, exit *struct{ i, j int }, addToGif func([][]byte)) int {
	expeditions := [][]struct{ i, j int }{{*start}}
	fmt.Println(start, exit, expeditions)
	var t int
exploration:
	for {
		t++
		var updatedExpds [][]struct{ i, j int }
		update(field, l, h, addToGif)
		m := fieldMap(field, l, h)
		freeSpots := dots(m)
		freeSpots = append(freeSpots, *start)
		freeSpots = append(freeSpots, *exit)
		// fmt.Println("turn", t+1)
		for n := range expeditions {
			nextMoves, _ := moveNext(expeditions[n], freeSpots)
			// if err != nil {
			// 	// log.Print(err)
			// }
			// fmt.Println("nextMoves", nextMoves)
			for _, move := range nextMoves {
				if move != *exit {
					continue
				}
				break exploration
			}
			switch len(nextMoves) {
			case 0:
				continue
			case 1:
				expeditions[n] = append(expeditions[n], nextMoves[0])
				updatedExpds = append(updatedExpds, expeditions[n])
			default:
				for i := range nextMoves {
					forkedExp := make([]struct{ i, j int }, len(expeditions[n]))
					copy(forkedExp, expeditions[n])
					forkedExp = append(forkedExp, nextMoves[i])
					updatedExpds = append(updatedExpds, forkedExp)
				}
			}
		}
		expMap := make(map[struct{ i, j int }]int)
		var filtered [][]struct{ i, j int }
		for i := range updatedExpds {
			endPoint := updatedExpds[i][len(updatedExpds[i])-1]
			if _, ok := expMap[endPoint]; ok {
				continue
			}
			expMap[endPoint]++
			filtered = append(filtered, updatedExpds[i])
		}
		expeditions = filtered
		// fmt.Println("exp", expeditions)
	}
	return t
}

func dots(m [][]byte) []struct{ i, j int } {
	var v []struct{ i, j int }
	for i := range m {
		if i == 0 || i == len(m)-1 {
			continue
		}
		for j := range m[i] {
			if m[i][j] != '.' {
				continue
			}
			v = append(v, struct{ i, j int }{i, j})
		}
	}
	return v
}

func fieldPoint(field []*cell, l, h int, filter func(*cell) bool) *struct{ i, j int } {
	tmpLine := bytes.Repeat([]byte{'.'}, l)
	var iLine int
	for i := range field {
		if filter(field[i]) {
			continue
		}
		tmpLine[field[i].j] = field[i].b
		iLine = field[i].i
	}
	return &struct{ i, j int }{i: iLine, j: bytes.Index(tmpLine, []byte{'.'})}
}

func update(field []*cell, l, h int, addToGif func([][]byte)) {
	for i := range field {
		iNext, jNext := field[i].i, field[i].j
		switch field[i].b {
		case '>':
			jNext = field[i].j + 1
			if jNext >= l-1 {
				jNext = 1
			}
		case '<':
			jNext = field[i].j - 1
			if jNext <= 0 {
				jNext = l - 2
			}
		case '^':
			iNext = field[i].i - 1
			if iNext <= 0 {
				iNext = h - 2
			}
		case 'v':
			iNext = field[i].i + 1
			if iNext >= h-1 {
				iNext = 1
			}
		}
		field[i].i = iNext
		field[i].j = jNext
	}
	addToGif(fieldMap(field, l, h))
}

type cell struct {
	i, j int
	b    byte
}

func parse(r io.Reader) ([]*cell, int, int) {
	scanner := bufio.NewScanner(r)
	var field []*cell
	var l, h int
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := range line {
			switch line[j] {
			case '.':
			default:
				field = append(field, &cell{i: i, j: j, b: line[j]})
			}
		}
		h = i + 1
		l = len(line)
	}
	return field, l, h
}

func moveNext(path []struct{ i, j int }, dots []struct{ i, j int }) ([]struct{ i, j int }, error) {
	expedition := path[len(path)-1]
	i, j := expedition.i, expedition.j
	candidates := []struct{ i, j int }{{i, j}, {i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	var avail []struct{ i, j int }
nextDot:
	for _, dot := range dots {
		for _, cand := range candidates {
			if dot != cand {
				continue
			}
			avail = append(avail, dot)
			continue nextDot
		}
	}
	if len(avail) == 0 {
		return nil, errors.New("no next cell found")
	}
	dist := func(a, b struct{ i, j int }) float64 {
		return math.Sqrt(float64((a.i-b.i)*(a.i-b.i) + (a.j-b.j)*(a.j-b.j)))
	}
	sort.Slice(avail, func(i, j int) bool { return dist(avail[i], avail[j]) < 0 })
	return avail, nil
}

func fieldMap(field []*cell, l, h int) [][]byte {
	b := make([][]byte, h)
	for i := range b {
		b[i] = bytes.Repeat([]byte{'.'}, l)
	}
	for i := range field {
		b[field[i].i][field[i].j] = field[i].b
	}
	return b
}

func print(field []*cell, l, h int) {
	b := fieldMap(field, l, h)
	for i := range b {
		fmt.Println(string(b[i]))
	}
}

func writePNG(b [][]byte, name string, l, h int) {
	f, _ := os.Create(name)
	defer f.Close()
	img := byteToImage(b, l, h)
	scaledImg := scaleTo(img, image.Rectangle{img.Bounds().Min, img.Bounds().Max.Mul(4)}, draw.NearestNeighbor)
	png.Encode(f, scaledImg)
}

func byteToImage(cave [][]byte, l, h int) image.Image {
	caveImg := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{l, h}}, palette.Plan9)
	for i := range cave {
		for j := range cave[i] {
			var r, g, b uint8
			switch cave[i][j] {
			case '#':
				r, g, b = 25, 10, 0
			case '.':
				r, g, b = 50, 0, 255
			case '2':
				r, g, b = 0, 160, 255
			case '3':
				r, g, b = 0, 170, 255
			case '4':
				r, g, b = 0, 180, 255
			case '5':
				r, g, b = 0, 190, 255
			case '6':
				r, g, b = 0, 200, 255
			case '7':
				r, g, b = 0, 210, 255
			case '8':
				r, g, b = 0, 220, 255
			case '9':
				r, g, b = 0, 230, 255
			default:
				r, g, b = 0, 250, 255
			}
			caveImg.Set(j, i, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return caveImg
}

func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func writeGIF(images []*image.Paletted, name string) {
	f, _ := os.Create(name)
	defer f.Close()
	delays := make([]int, len(images))
	for i := range delays {
		delays[i] = 50
	}
	err := gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func gifAdder(images *[]*image.Paletted, l, h int) func([][]byte) {
	return func(b [][]byte) {
		img := byteToImage(b, l, h)
		if i, ok := img.(*image.Paletted); ok {
			*images = append(*images, i)
		}
	}
}
