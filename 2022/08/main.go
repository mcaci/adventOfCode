package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

const zerosByte = 48

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var treesList []*tree
	var l int
	scanner := bufio.NewScanner(f)
	for lineN := 0; scanner.Scan(); lineN++ {
		line := scanner.Text()
		for i := range line {
			treesList = append(treesList, newTree(lineN, i, line[i]-zerosByte, func(i1, i2 int) bool {
				if i1 == 0 || i2 == 0 {
					return true
				}
				if i1 == len(line)-1 || i2 == len(line)-1 {
					return true
				}
				return false
			}))
		}
		l = len(line)
	}
	setNeighbours(treesList, l)
	// part 1
	visit(treesList[0], func(t *tree) bool {
		if visibleFrom(t, func(t *tree) *tree { return t.top }) {
			return true
		}
		if visibleFrom(t, func(t *tree) *tree { return t.left }) {
			return true
		}
		if visibleFrom(t, func(t *tree) *tree { return t.bottom }) {
			return true
		}
		if visibleFrom(t, func(t *tree) *tree { return t.right }) {
			return true
		}
		return false
	})
	var count int
	for i := range treesList {
		if !treesList[i].visible {
			continue
		}
		count++
	}
	fmt.Println("part 1", count)
	// part 2
	var score int
	for i := range treesList {
		s := scenicScore(treesList[i])
		if score > s {
			continue
		}
		count++
		score = s
	}
	fmt.Println("part 2", score)
	// image
	forest := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{l, l}})
	for _, t := range treesList {
		forest.Set(t.i, t.j, color.RGBA{G: 35 + uint8(175.0/float64(t.height+1)), A: 255})
	}
	scaledForest := scaleTo(forest, image.Rectangle{image.Point{0, 0}, image.Point{27 * l, 27 * l}}, draw.NearestNeighbor)
	fForest, _ := os.Create("forest.png")
	defer fForest.Close()
	png.Encode(fForest, scaledForest)
}

func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

type tree struct {
	i, j                     int
	height                   byte
	visible, visited         bool
	top, right, bottom, left *tree
	scenicScore              int
}

func (t tree) String() string {
	return fmt.Sprintf("i: %d, j:%d, h:%d, visible:%t\n", t.i, t.j, t.height, t.visible)
}

func newTree(i, j int, height byte, atBorder func(int, int) bool) *tree {
	t := &tree{i: i, j: j, height: height}
	if atBorder(i, j) {
		t.visible = true
	}
	return t
}

func setNeighbours(treesList []*tree, l int) {
	for i, t := range treesList {
		if i%l > 0 {
			t.left = treesList[i-1]
		}
		if i/l < l-1 {
			t.bottom = treesList[i+l]
		}
		if i%l < l-1 {
			t.right = treesList[i+1]
		}
		if i/l > 0 {
			t.top = treesList[i-l]
		}
	}
}

func visit(t *tree, visible func(*tree) bool) {
	if t == nil {
		return
	}
	if t.visited {
		return
	}
	t.visited = true
	if visible(t) {
		t.visible = true
	}
	visit(t.left, visible)
	visit(t.bottom, visible)
	visit(t.right, visible)
	visit(t.top, visible)
}

func visibleFrom(t *tree, assignNext func(t *tree) *tree) bool {
	t2 := assignNext(t)
	for t2 != nil {
		visible := t.height > t2.height
		if !visible {
			return false
		}
		t2 = assignNext(t2)
	}
	return true
}

func countFrom(t *tree, assignNext func(t *tree) *tree) int {
	var count int
	t2 := assignNext(t)
	for t2 != nil {
		count++
		visible := t.height > t2.height
		if !visible {
			return count
		}
		t2 = assignNext(t2)
	}
	return count
}

func scenicScore(t *tree) int {
	u := countFrom(t, func(t *tree) *tree { return t.top })
	l := countFrom(t, func(t *tree) *tree { return t.left })
	b := countFrom(t, func(t *tree) *tree { return t.bottom })
	r := countFrom(t, func(t *tree) *tree { return t.right })
	t.scenicScore = u * l * b * r
	return t.scenicScore
}
