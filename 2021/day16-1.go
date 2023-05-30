package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main_day16part1() {
	day16part1()
}

const (
	EX1    = "8A004A801A8002F478"
	EX1_VS = 16
	EX2    = "620080001611562C8802118E34"
	EX2_VS = 12
	EX3    = "C0015000016115A2E0802F182340"
	EX3_VS = 23
	EX4    = "A0016C880162017C3686B18A3D4780"
	EX4_VS = 31
)

func day16part1() {
	ex1 := flag.Bool("ex1", false, "trigger example 1")
	ex2 := flag.Bool("ex2", false, "trigger example 2")
	ex3 := flag.Bool("ex3", false, "trigger example 3")
	ex4 := flag.Bool("ex4", false, "trigger example 4")
	flag.Parse()
	var expected int
	var r io.ByteReader
	var in string
	switch {
	case *ex1:
		in = EX1
		r = strings.NewReader(EX1)
		expected = EX1_VS
	case *ex2:
		in = EX2
		r = strings.NewReader(EX2)
		expected = EX2_VS
	case *ex3:
		in = EX3
		r = strings.NewReader(EX3)
		expected = EX3_VS
	case *ex4:
		in = EX4
		r = strings.NewReader(EX4)
		expected = EX4_VS
	default:
		f, err := os.Open("day16")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = bufio.NewReader(f)
	}
	var binMsg []byte
inRead:
	for {
		b, err := r.ReadByte()
		switch err {
		case nil:
		case io.EOF:
			break inRead
		default:
			log.Fatal(err)
		}
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			binMsg = append(binMsg, []byte(fmt.Sprintf("%04b", b-'0'))...)
		// error: forgot +10 in the computation of the letters for hexadecimal
		case 'A', 'B', 'C', 'D', 'E', 'F':
			binMsg = append(binMsg, []byte(fmt.Sprintf("%04b", b-'A'+10))...)
		}
	}
	cpMsg := make([]byte, len(binMsg))
	copy(cpMsg, binMsg)
	fmt.Println(in)
	fmt.Println(string(cpMsg))
	v := parsePackets(bytes.NewBuffer(cpMsg))
	switch {
	case *ex1, *ex2, *ex3, *ex4:
		fmt.Println(v, v == int64(expected))
	default:
		fmt.Println(v)
	}
}

func parsePackets(r io.Reader) int64 {
	var v int64
	var done bool
	for !done {
		vTmp, err := parseBits(r, 3)
		if err != nil {
			log.Print(err)
			if err == io.EOF {
				done = true
			}
		}
		log.Print("VERSION is ", vTmp)
		v += vTmp

		t, err := parseBits(r, 3)
		if err != nil {
			log.Print(err)
		}
		switch t {
		case 4:
			log.Print("type is ", t, ": this is a LITERAL")
			var done bool
			var binN string
			for !done {
				firstBit, err := parseBits(r, 1)
				log.Print("first bit is ", firstBit, ": if it's 0 it stops, if it's 1 it continues")
				if err != nil {
					log.Print(err)
				}
				s, err := parseString(r, 4)
				if err != nil {
					log.Print(err)
				}
				binN += s
				if firstBit == 0 {
					done = true
				}
			}
			log.Print("literal is ", binN)
		default:
			log.Print("type is ", t, ": this is an OPERATOR")
			modeBit, err := parseBits(r, 1)
			if err != nil {
				log.Print(err)
			}
			switch modeBit {
			case 0:
				log.Print("MODEBIT is ", modeBit, ": next 15 bits will give the max len of subpackets")
				l, err := parseBits(r, 15)
				if err != nil {
					log.Print(err)
				}
				log.Print("max len of subpackets is ", l)
			case 1:
				log.Print("MODEBIT is ", modeBit, ": next 11 bits will give the number of subpackets immediately contained by this packet")
				l, err := parseBits(r, 11)
				if err != nil {
					log.Print(err)
				}
				log.Print("number of subpackets immediately contained by this packet is ", l)
			default:
				log.Fatalf("unexpected modeBit %d", modeBit)
			}
		}
	}
	return v
}

func parseBits(r io.Reader, nBits int64) (int64, error) {
	buf := make([]byte, nBits)
	nRead, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	if nRead > 0 && nRead < len(buf) {
		return 0, fmt.Errorf("not enough bits vs buffer len: %d < %d", nRead, len(buf))
	}
	log.Printf("reading %d bit(s), bits read %q", nRead, string(buf))
	n, err := strconv.ParseInt(string(buf), 2, 64)
	if err != nil {
		return 0, err
	}
	return n, err
}

func parseString(r io.Reader, nBits int64) (string, error) {
	buf := make([]byte, nBits)
	nRead, err := r.Read(buf)
	if nRead > 0 && nRead < len(buf) {
		return "", fmt.Errorf("not enough bits vs buffer len: %d < %d", nRead, len(buf))
	}
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
