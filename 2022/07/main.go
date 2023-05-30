package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	root := &file{name: "/"}
	var current = root
	var total int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		l := parse(line)
		switch cmd := l.(type) {
		case cd:
			current = find(cmd, current)
		case ls:
		case file:
			current.list = append(current.list, &file{name: cmd.name, size: cmd.size, pdir: current})
			total += cmd.size
		}
	}
	// part 1
	// filter := func(f *file) bool { return !(f.size == 0 && f.fullSize() <= 100000) }
	// part 2
	filter := func(f *file) bool { return !(f.size == 0) }
	dirSizes := visit(root, filter)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirSizes)
	var sum int
	for i := range dirSizes {
		sum += dirSizes[i]
	}
	fmt.Println("part 1", sum)
	const (
		diskSize = 70000000
		needed   = 30000000
	)
	occupied := diskSize - total
	toFree := needed - occupied
	sort.Ints(dirSizes)
	idx := sort.Search(len(dirSizes), func(i int) bool {
		return total-dirSizes[i] <= diskSize-needed
	})
	fmt.Println(dirSizes)
	fmt.Println(diskSize, total, occupied, toFree)
	fmt.Println("part 2", dirSizes[idx], total-dirSizes[idx])
}

func visit(f *file, filter func(*file) bool) []int {
	var sizes []int
	if !filter(f) {
		sizes = append(sizes, f.fullSize())
	}
	for i := range f.list {
		sizes = append(sizes, visit(f.list[i], filter)...)
	}
	return sizes
}

func parse(line string) any {
	f := strings.Fields(line)
	switch f[0] {
	case "$":
		switch f[1] {
		case "ls":
			return ls{}
		case "cd":
			return cd{path: f[2]}
		}
	case "dir":
		return file{name: f[1]}
	default:
		s, _ := strconv.Atoi(f[0])
		return file{name: f[1], size: s}
	}
	return line
}

type file struct {
	name string
	size int
	list []*file
	pdir *file
}

func (f file) fullSize() int {
	if f.size != 0 {
		return f.size
	}
	var totalSize int
	for _, subF := range f.list {
		if subF.size != 0 {
			totalSize += subF.size
			continue
		}
		totalSize += subF.fullSize()
	}
	return totalSize
}

type ls struct{}
type cd struct{ path string }

func find(cmd cd, root *file) *file {
	switch cmd.path {
	case "/":
		return root
	case "..":
		return root.pdir
	default:
		for _, f := range root.list {
			if f.name != cmd.path {
				continue
			}
			return f
		}
	}
	return nil
}
