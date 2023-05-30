package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var common = p2(scanner)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	n := bytes.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return r - 97 + 1
		}
		if unicode.IsUpper(r) {
			return r - 65 + 27
		}
		return 0
	}, common)
	fmt.Println(len(n), n)
	var sum int
	for _, v := range n {
		sum += int(v)
	}
	fmt.Println(sum)
}

func p1(scanner *bufio.Scanner) []byte {
	var common []byte
nextLine:
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		fmt.Println(line[:mid], line[mid:])
		// nextChar:
		for _, c1 := range line[:mid] {
			for _, c2 := range line[mid:] {
				if c1 != c2 {
					continue
				}
				common = append(common, byte(c1))
				continue nextLine
			}
		}
		// fmt.Println(line)
	}
	return common
}

func p2(scanner *bufio.Scanner) []byte {
	var common []byte

	for {
		scanner.Scan()
		line1 := scanner.Text()
		m1 := make(map[byte]bool)
		for i := range line1 {
			m1[line1[i]] = true
		}
		scanner.Scan()
		line2 := scanner.Text()
		m2 := make(map[byte]bool)
		for i := range line2 {
			m2[line2[i]] = true
		}
		ok := scanner.Scan()
		line3 := scanner.Text()
		m3 := make(map[byte]bool)
		for i := range line3 {
			m3[line3[i]] = true
		}

		isBadge := func(m1, m2, m3 map[byte]bool, b byte) bool {
			return m1[b] && m2[b] && m3[b]
		}
		for i := byte(0); i < 150; i++ {
			if !isBadge(m1, m2, m3, i) {
				continue
			}
			common = append(common, i)
			break
		}
		if !ok {
			break
		}
	}
	return common
}
