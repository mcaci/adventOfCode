package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var v [9][]byte
	var lineN int
	for scanner.Scan() {
		line := scanner.Text()
		lineN++
		if lineN < 9 {
			for i := range line {
				if (i-1)%4 == 0 && line[i] != ' ' {
					v[i/4] = append([]byte{line[i]}, v[i/4]...)
				}
			}
		}
		if lineN > 10 {
			cmdStr := strings.Fields(line)
			nCrates, from, to := cmdStr[1], cmdStr[3], cmdStr[5]
			cmd := newCommandP2(nCrates, from, to)
			v = cmd.apply(v)
		}
	}
	var s strings.Builder
	for i := range v {
		s.WriteByte(v[i][len(v[i])-1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.String())
}

type command struct{ n, from, to int }

func newCommand(n, from, to string) *command {
	nInt, _ := strconv.Atoi(n)
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)
	return &command{n: nInt, from: fromInt - 1, to: toInt - 1}
}
func (cmd *command) apply(v [9][]byte) [9][]byte {
	for i := 0; i < cmd.n; i++ {
		lastIdx := len(v[cmd.from]) - 1
		last := v[cmd.from][lastIdx]
		v[cmd.to] = append(v[cmd.to], last)
		v[cmd.from] = v[cmd.from][:lastIdx]
	}
	return v
}

type commandP2 command

func newCommandP2(n, from, to string) *commandP2 {
	cmd := newCommand(n, from, to)
	cmd2 := commandP2(*cmd)
	return &cmd2
}
func (cmd *commandP2) apply(v [9][]byte) [9][]byte {
	lastIdx := len(v[cmd.from]) - cmd.n
	lastBlock := v[cmd.from][lastIdx:]
	v[cmd.to] = append(v[cmd.to], lastBlock...)
	v[cmd.from] = v[cmd.from][:lastIdx]
	return v
}
