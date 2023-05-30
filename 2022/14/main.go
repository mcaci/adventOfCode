package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/draw"
)

const (
	d     = 300
	l     = 400
	h     = 180
	sandY = 500 - d
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var caveImgSeq []*image.Paletted
	addToGif := gifAdder(&caveImgSeq)

	cave := make([][]byte, h)
	for i := range cave {
		cave[i] = make([]byte, l)
		// for j := range cave[i] {
		// cave[i][j] = ' '
		// }
	}
	scanner := bufio.NewScanner(f)
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()
		rockLines := strings.Split(line, " -> ")
		for i := range rockLines {
			if i == 0 {
				continue
			}
			xy := func(s string) (int, int) {
				y, _ := strconv.Atoi(s[:strings.Index(s, ",")])
				x, _ := strconv.Atoi(s[strings.Index(s, ",")+1:])
				return x, y - d
			}
			s, e := rockLines[i-1], rockLines[i]
			x1, y1 := xy(s)
			x2, y2 := xy(e)
			switch {
			case y1 == y2:
				for i := int(math.Min(float64(x1), float64(x2))); i <= int(math.Max(float64(x1), float64(x2))); i++ {
					cave[i][y1] = '#'
				}
			case x1 == x2:
				for i := int(math.Min(float64(y1), float64(y2))); i <= int(math.Max(float64(y1), float64(y2))); i++ {
					cave[x1][i] = '#'
				}
			}
		}
	}
	var actualH int
	for i := len(cave) - 1; i >= 0; i-- {
		if !bytes.Contains(cave[i], []byte{'#'}) {
			continue
		}
		actualH = i + 3
		break
	}
	actualCave := make([][]byte, actualH)
	for i := range actualCave {
		actualCave[i] = make([]byte, l)
		copy(actualCave[i], cave[i])
	}
	for i := range actualCave[len(actualCave)-1] {
		actualCave[len(actualCave)-1][i] = '#'
	}
	// image
	// writePNG(cave, "cave.png")
	if addToGif != nil {
		addToGif(cave)
	}
	// part 1
	var count int
	for sandFall(cave, h, sandY, nil) {
		count++
	}
	// writePNG(cave, "cave-with-sand.png")
	fmt.Println("part 1", count)
	// part 2
	count = 0
	for sandFallActualCave(actualCave, actualH, sandY, nil) {
		count++
	}
	fmt.Println("part 2", count)
	// writePNG(actualCave, "actual-cave-with-sand.png")
	// gif
	writeGIF(caveImgSeq)

}

func sandFall(cave [][]byte, h, sandY int, addToGif func([][]byte)) bool {
	var x, y, x1, y1 int
	y = sandY
	cave[x][y] = '*'
	blocked := func(b byte) bool {
		switch b {
		case '*', '#':
			return true
		default:
			return false
		}
	}
	for {
		x1, y1 = x, y
		if x+1 >= h {
			return false
		}
		if !blocked(cave[x+1][y]) {
			x++
			cave[x][y] = '*'
			cave[x1][y] = 0
			continue
		}
		if !blocked(cave[x+1][y-1]) {
			x++
			y--
			cave[x][y] = '*'
			cave[x1][y1] = 0
			continue
		}
		if !blocked(cave[x+1][y+1]) {
			x++
			y++
			cave[x][y] = '*'
			cave[x1][y1] = 0
			continue
		}
		if addToGif != nil {
			addToGif(cave)
		}
		return true
	}
}

func sandFallActualCave(cave [][]byte, h, sandY int, addToGif func([][]byte)) bool {
	var x, y, x1, y1 int
	y = sandY
	if cave[x][y] == '*' {
		return false
	}
	cave[x][y] = '*'
	blocked := func(b byte) bool {
		switch b {
		case '*', '#':
			return true
		default:
			return false
		}
	}
	for {
		x1, y1 = x, y
		if x+1 >= h {
			return false
		}
		if !blocked(cave[x+1][y]) {
			x++
			cave[x][y] = '*'
			cave[x1][y] = 0
			continue
		}
		if !blocked(cave[x+1][y-1]) {
			x++
			y--
			cave[x][y] = '*'
			cave[x1][y1] = 0
			continue
		}
		if !blocked(cave[x+1][y+1]) {
			x++
			y++
			cave[x][y] = '*'
			cave[x1][y1] = 0
			continue
		}
		if addToGif != nil {
			addToGif(cave)
		}
		return true
	}
}

func caveSnapshot(cave [][]byte) image.Image {
	caveImg := image.NewPaletted(image.Rectangle{image.Point{0, 0}, image.Point{l, h}}, palette.Plan9)
	for i := range cave {
		for j := range cave[i] {
			var r, g, b uint8
			switch cave[i][j] {
			case 0:
				r, g, b = 25, 10, 0
			case '*':
				r, g, b = 150, 150, 0
			default:
				r, g, b = 200, 100, 0
			}
			caveImg.Set(j, i, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return caveImg
}

func writePNG(b [][]byte, name string) {
	f, _ := os.Create(name)
	defer f.Close()
	img := caveSnapshot(b)
	scaledImg := scaleTo(img, image.Rectangle{img.Bounds().Min, img.Bounds().Max.Mul(4)}, draw.NearestNeighbor)
	png.Encode(f, scaledImg)
}

func writeGIF(images []*image.Paletted) {
	const name = "cave.gif"
	f, _ := os.Create(name)
	defer f.Close()
	delays := make([]int, len(images))
	for i := range delays {
		delays[i] = 5
	}
	err := gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func gifAdder(images *[]*image.Paletted) func([][]byte) {
	return func(b [][]byte) {
		img := caveSnapshot(b)
		if i, ok := img.(*image.Paletted); ok {
			*images = append(*images, i)
		}
	}
}

func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}
