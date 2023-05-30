package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sort"

	"golang.org/x/image/draw"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var steps []*hillStep
	var h, l int
	scanner := bufio.NewScanner(f)
	var start, end *hillStep
	for r := 0; scanner.Scan(); r++ {

		line := scanner.Text()
		for c := range line {
			var step *hillStep
			switch line[c] {
			case 'S':
				step = newStep(c, r, 'a'-1)
				step.distance = 0
				start = step
			case 'E':
				step = newStep(c, r, 'z'+1)
				end = step
			default:
				step = newStep(c, r, line[c])
			}
			steps = append(steps, step)
		}
		l = len(line)
		h++
	}
	setNeighbours(steps, l, h)
	// part 1
	bfs(start, steps)
	fmt.Println("part 1:", end.distance)
	// part 2
	endDists := []int{end.distance}
	for i := range steps {
		if steps[i].height != 'a' {
			continue
		}
		var nextIsB bool
		if steps[i].left != nil && steps[i].left.height == 'b' {
			nextIsB = true
		}
		if steps[i].top != nil && steps[i].top.height == 'b' {
			nextIsB = true
		}
		if steps[i].right != nil && steps[i].right.height == 'b' {
			nextIsB = true
		}
		if steps[i].bottom != nil && steps[i].bottom.height == 'b' {
			nextIsB = true
		}
		if !nextIsB {
			continue
		}
		for j := range steps {
			steps[j].visited = false
			steps[j].distance = math.MaxInt
		}
		steps[i].distance = 0
		log.Println(steps[i])
		bfs(steps[i], steps)
		endDists = append(endDists, end.distance)
	}
	sort.Ints(endDists)
	fmt.Println("part 2:", endDists, endDists[0])
	// image
	hill := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{l, h}})
	for _, s := range steps {
		hill.Set(s.i, s.j, color.RGBA{
			R: uint8(55 + 200*(float64(s.height)-('a'-1))/float64('z'+1-('a'-1))),
			G: uint8(50 + 150*(float64(s.height)-('a'-1))/float64('z'+1-('a'-1))),
			B: uint8(25 + 100*(float64(s.height)-('a'-1))/float64('z'+1-('a'-1))),
			A: 255})
	}
	scaledHill := scaleTo(hill, image.Rectangle{image.Point{0, 0}, image.Point{27 * l, 27 * h}}, draw.NearestNeighbor)
	fHill, _ := os.Create("hill.png")
	defer fHill.Close()
	png.Encode(fHill, scaledHill)
}

func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

type hillStep struct {
	i, j                     int
	height                   byte
	visited                  bool
	top, right, bottom, left *hillStep
	distance                 int
}

func (s hillStep) String() string {
	return fmt.Sprintf("i: %d, j:%d, h:%s, d: %d", s.i, s.j, string(s.height), s.distance)
}

func newStep(i, j int, height byte) *hillStep {
	return &hillStep{i: i, j: j, height: height, distance: math.MaxInt}
}

func setNeighbours(steps []*hillStep, l, h int) {
	for i, s := range steps {
		if s.i > 0 && steps[i-1].height <= s.height+1 {
			s.left = steps[i-1]
		}
		if s.j < h-1 && steps[i+l].height <= s.height+1 {
			s.bottom = steps[i+l]
		}
		if s.i < l-1 && steps[i+1].height <= s.height+1 {
			s.right = steps[i+1]
		}
		if s.j > 0 && steps[i-l].height <= s.height+1 {
			s.top = steps[i-l]
		}
	}
}

func find(steps []*hillStep, filter func(*hillStep) bool) *hillStep {
	for _, s := range steps {
		if !filter(s) {
			continue
		}
		return s
	}
	return nil
}

func setDist(s *hillStep, steps []*hillStep) {
	var dists []int
	if right := find(steps, func(hs *hillStep) bool { return hs.left == s }); right != nil {
		dists = append(dists, right.distance)
	}
	if left := find(steps, func(hs *hillStep) bool { return hs.right == s }); left != nil {
		dists = append(dists, left.distance)
	}
	if top := find(steps, func(hs *hillStep) bool { return hs.bottom == s }); top != nil {
		dists = append(dists, top.distance)
	}
	if bottom := find(steps, func(hs *hillStep) bool { return hs.top == s }); bottom != nil {
		dists = append(dists, bottom.distance)
	}
	sort.Ints(dists)
	s.distance = dists[0] + 1
}

func bfs(start *hillStep, steps []*hillStep) {
	var stack []*hillStep
	stack = append(stack, start)
	for i := 0; len(stack) > 0; i++ {
		e := stack[0]
		stack = stack[1:]
		if e.visited {
			continue
		}
		e.visited = true
		if e != start {
			setDist(e, steps)
		}
		if e.left != nil {
			stack = append(stack, e.left)
		}
		if e.top != nil {
			stack = append(stack, e.top)
		}
		if e.right != nil {
			stack = append(stack, e.right)
		}
		if e.bottom != nil {
			stack = append(stack, e.bottom)
		}
	}

}
