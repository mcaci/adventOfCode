package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	myimage "github.com/mcaci/adventOfCode/image"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []game
	var cardID int
	for scanner.Scan() {
		cardID++
		line := scanner.Text()
		colonId := strings.Index(line, ":")
		line = line[colonId+1:]
		sections := strings.Split(line, "|")
		var win []int
		w := strings.Trim(sections[0], " ")
		for _, s := range strings.Split(w, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			win = append(win, n)
		}
		var numbers []int
		ns := strings.Trim(sections[1], " ")
		for _, s := range strings.Split(ns, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			numbers = append(numbers, n)
		}
		games = append(games, game{cardID: cardID, winners: win, numbers: numbers, copies: 1})
	}
	var sum float64
	// part 1
	for i, g := range games {
		var count float64
		for _, n := range g.numbers {
			if !slices.Contains(g.winners, n) {
				continue
			}
			count++
		}
		games[i].matchCount = int(count)
		if count == 0 {
			continue
		}
		sum += math.Exp2(count - 1)
	}
	fmt.Println(sum)
	l, r := 10, 10
	img := myimage.NewRGBA(l, r, func(img *image.RGBA) {
		for i := 0; i < l*r; i++ {
			x, y := i/l, i%r
			switch {
			case slices.Contains(games[0].numbers, i) && slices.Contains(games[0].winners, i):
				img.Set(x, y, color.RGBA{B: 0, R: 0, G: 100, A: 255})
			case slices.Contains(games[0].winners, i):
				img.Set(x, y, color.RGBA{B: 0, R: 100, G: 0, A: 255})
			case slices.Contains(games[0].numbers, i):
				img.Set(x, y, color.RGBA{B: 0, R: 0, G: 0, A: 255})
			default:
				img.Set(x, y, color.RGBA{B: 150, R: 150, G: 150, A: 255})
			}
		}
	})
	myimage.Save(img, "card1.png", 50)
	var frames []*image.Paletted
	for i, g := range games {
		log.Print(i)
		l, r := 10, 10
		img := myimage.NewPaletted(l, r, func(img *image.Paletted) {
			for i := 0; i < l*r; i++ {
				x, y := i/l, i%r
				switch {
				case slices.Contains(g.numbers, i) && slices.Contains(g.winners, i):
					img.Set(x, y, color.RGBA{B: 0, R: 0, G: 100, A: 255})
				case slices.Contains(g.winners, i):
					img.Set(x, y, color.RGBA{B: 0, R: 100, G: 0, A: 255})
				case slices.Contains(g.numbers, i):
					img.Set(x, y, color.RGBA{B: 0, R: 0, G: 0, A: 255})
				default:
					img.Set(x, y, color.RGBA{B: 150, R: 150, G: 150, A: 255})
				}
			}
		}, 50)
		frames = append(frames, img)
	}
	myimage.SaveGIF(frames, "cards.gif", 50, 75)

	// part 2
	for i, g := range games {
		// fmt.Println(games[i])
		for j := 0; j < g.copies; j++ {
			l := min(i+1, len(games))
			r := min(l+g.matchCount, len(games))
			for k := range games[l:r] {
				games[l:r][k].copies++
			}
		}
	}
	var count uint64
	for i := range games {
		count += uint64(games[i].copies)
	}
	fmt.Println(count)
}

type game struct {
	cardID     int
	winners    []int
	numbers    []int
	matchCount int
	copies     int
}

func (g game) String() string {
	return fmt.Sprintf("<%d:%d-%d>", g.cardID, g.matchCount, g.copies)
}
