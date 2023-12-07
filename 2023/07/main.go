package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand string

func cardCount(h hand) map[byte]uint {
	m := make(map[byte]uint)
	for i := range h {
		m[h[i]]++
	}
	return m
}

func FiveOfAKind(h hand) bool {
	c := cardCount(h)
	for _, v := range c {
		if v == 5 {
			return true
		}
	}
	return false
}
func FourOfAKind(h hand) bool {
	c := cardCount(h)
	for _, v := range c {
		if v == 4 {
			return true
		}
	}
	return false
}
func FullHouse(h hand) bool {
	c := cardCount(h)
	for _, v := range c {
		if (v == 2 || v == 3) && len(c) == 2 {
			return true
		}
	}
	return false
}
func ThreeOfAKind(h hand) bool {
	c := cardCount(h)
	for _, v := range c {
		if v == 3 && len(c) == 3 {
			return true
		}
	}
	return false
}
func TwoPair(h hand) bool {
	c := cardCount(h)
	var pairCount uint
	for _, v := range c {
		if v != 2 {
			continue
		}
		pairCount++
	}
	return pairCount == 2
}
func OnePair(h hand) bool {
	c := cardCount(h)
	var pairCount uint
	for _, v := range c {
		if v != 2 {
			continue
		}
		pairCount++
	}
	return pairCount == 1
}
func HighCard(h hand) bool {
	c := cardCount(h)
	return len(c) == 5
}
func FiveOfAKindWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0:
		return FiveOfAKind(h)
	case 1:
		return FourOfAKind(h)
	case 2, 3:
		return FullHouse(h)
	default:
		return true
	}
}
func FourOfAKindWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0:
		return FourOfAKind(h)
	case 1:
		return FourOfAKind(h) || ThreeOfAKind(h)
	case 2:
		return ThreeOfAKind(h) || TwoPair(h) || FullHouse(h)
	default:
		return true
	}
}
func FullHouseWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0, 3:
		return FullHouse(h)
	case 1:
		return ThreeOfAKind(h)
	case 2:
		return FullHouse(h) || TwoPair(h)
	default:
		return true
	}
}
func ThreeOfAKindWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0:
		return ThreeOfAKind(h)
	case 1:
		return OnePair(h) || TwoPair(h)
	default:
		return true
	}
}
func TwoPairWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0, 2:
		return TwoPair(h)
	case 1:
		return OnePair(h)
	default:
		return true
	}
}
func OnePairWithJoker(h hand) bool {
	c := cardCount(h)
	switch c['J'] {
	case 0:
		return OnePair(h)
	default:
		return true
	}
}

func Less(h1, h2 hand) bool {
	fs := []func(hand) bool{FiveOfAKindWithJoker, FourOfAKindWithJoker, FullHouseWithJoker, ThreeOfAKindWithJoker, TwoPairWithJoker, OnePairWithJoker, HighCard}
	// cardOrder := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	cardOrder := []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	var is, js []bool
	for _, f := range fs {
		is, js = append(is, f(h1)), append(js, f(h2))
	}
	switch {
	case slices.Index(is, true) < slices.Index(js, true):
		return false
	case slices.Index(is, true) > slices.Index(js, true):
		return true
	default:
		// slices.Index(is, true) == slices.Index(js, true):
		for i := 0; i < 5; i++ {
			if h1[i] == h2[i] {
				continue
			}
			return slices.Index(cardOrder, h1[i]) > slices.Index(cardOrder, h2[i])
		}
	}
	log.Print("cannot get to this false")
	return false
}

type game struct {
	hand hand
	bid  int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []game
	for scanner.Scan() {
		line := scanner.Text()
		f := strings.Fields(line)
		n, _ := strconv.Atoi(f[1])
		games = append(games, game{hand: hand(f[0]), bid: n})
	}
	sort.Slice(games, func(i, j int) bool { return Less(games[i].hand, games[j].hand) })
	// log.Print(games)
	var sum uint
	for i := range games {
		sum += uint(i+1) * uint(games[i].bid)
	}
	// Part 1:  241344943
	// Part 2:  243101568
	log.Print(sum)
}
