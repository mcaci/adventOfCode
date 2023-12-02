package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game [3]int

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var games []game
	for scanner.Scan() {
		var g game
		line := scanner.Text()
		colonPos := strings.Index(line, ":")
		line = line[colonPos+2:]
		line = strings.ReplaceAll(line, ";", "")
		line = strings.ReplaceAll(line, ",", "")
		lines := strings.Split(line, " ")
		for i := 0; i < len(lines); i += 2 {
			c, _ := strconv.Atoi(lines[i])
			switch lines[i+1] {
			case "blue":
				if g[0] < c {
					g[0] = c
				}
			case "green":
				if g[1] < c {
					g[1] = c
				}
			case "red":
				if g[2] < c {
					g[2] = c
				}
			}
		}
		games = append(games, g)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var sum int
	for _, g := range games {
		sum += g[0] * g[1] * g[2]
	}
	fmt.Println(sum)
}

func mainPart1() {
	file, err := os.Open("2023/02/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int
	var lineID int
	scanner := bufio.NewScanner(file)
nextGame:
	for scanner.Scan() {
		lineID++
		line := scanner.Text()
		colonPos := strings.Index(line, ":")
		line = line[colonPos+2:]
		line = strings.ReplaceAll(line, ";", "")
		line = strings.ReplaceAll(line, ",", "")
		lines := strings.Split(line, " ")
		for i := 0; i < len(lines); i += 2 {
			c, _ := strconv.Atoi(lines[i])
			switch lines[i+1] {
			case "blue":
				if c > 14 {
					continue nextGame
				}
			case "green":
				if c > 13 {
					continue nextGame
				}
			case "red":
				if c > 12 {
					continue nextGame
				}
			}
		}
		sum += lineID
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
