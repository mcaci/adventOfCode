package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	day16part2()
}

const (
	P2_EX1   = "C200B40A82"
	P2_EX1_R = 3
	P2_EX2   = "04005AC33890"
	P2_EX2_R = 54
	P2_EX3   = "880086C3E88112"
	P2_EX3_R = 7
	P2_EX4   = "CE00C43D881120"
	P2_EX4_R = 9
	P2_EX5   = "D8005AC2A8F0"
	P2_EX5_R = 1
	P2_EX6   = "F600BC2D8F"
	P2_EX6_R = 0
	P2_EX7   = "9C005AC2F8F0"
	P2_EX7_R = 0
	P2_EX8   = "9C0141080250320F1802104A08"
	P2_EX8_R = 1
)

var opMap = map[int]func(...int) int{
	0: func(ns ...int) int {
		var sum int
		for _, n := range ns {
			sum += n
		}
		return sum
	},
	1: func(ns ...int) int {
		prod := 1
		for _, n := range ns {
			prod *= n
		}
		return prod
	},
	2: func(ns ...int) int {
		min := math.MaxInt32
		for _, n := range ns {
			if n > min {
				continue
			}
			min = n
		}
		return min
	},
	3: func(ns ...int) int {
		var max int
		for _, n := range ns {
			if n < max {
				continue
			}
			max = n
		}
		return max
	},
	5: func(ns ...int) int {
		if ns[0] > ns[1] {
			return 1
		}
		return 0
	},
	6: func(ns ...int) int {
		if ns[0] < ns[1] {
			return 1
		}
		return 0
	},
	7: func(ns ...int) int {
		if ns[0] == ns[1] {
			return 1
		}
		return 0
	},
}

func day16part2() {
	ex1 := flag.Bool("ex1", false, "trigger example 1")
	ex2 := flag.Bool("ex2", false, "trigger example 2")
	ex3 := flag.Bool("ex3", false, "trigger example 3")
	ex4 := flag.Bool("ex4", false, "trigger example 4")
	ex5 := flag.Bool("ex5", false, "trigger example 5")
	ex6 := flag.Bool("ex6", false, "trigger example 6")
	ex7 := flag.Bool("ex7", false, "trigger example 7")
	ex8 := flag.Bool("ex8", false, "trigger example 8")
	flag.Parse()
	var expected int
	var r io.ByteReader
	var in string
	switch {
	case *ex1:
		in = P2_EX1
		r = strings.NewReader(in)
		expected = P2_EX1_R
	case *ex2:
		in = P2_EX2
		r = strings.NewReader(in)
		expected = P2_EX2_R
	case *ex3:
		in = P2_EX3
		r = strings.NewReader(in)
		expected = P2_EX3_R
	case *ex4:
		in = P2_EX4
		r = strings.NewReader(in)
		expected = P2_EX4_R
	case *ex5:
		in = P2_EX5
		r = strings.NewReader(in)
		expected = P2_EX5_R
	case *ex6:
		in = P2_EX6
		r = strings.NewReader(in)
		expected = P2_EX6_R
	case *ex7:
		in = P2_EX7
		r = strings.NewReader(in)
		expected = P2_EX7_R
	case *ex8:
		in = P2_EX8
		r = strings.NewReader(in)
		expected = P2_EX8_R
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
	res := parseMsg(bytes.NewBuffer(cpMsg))
	fmt.Println(res)
	t := opTree(res)
	t.solve()
	result := t.val
	switch {
	case *ex1, *ex2, *ex3, *ex4, *ex5, *ex6, *ex7, *ex8:
		fmt.Println(res)
		fmt.Println(result, result == expected)
	default:
		fmt.Println(res)
		fmt.Println(result)
	}
}

type opNode struct {
	op    func(...int) int
	val   int
	child []*opNode
}

func (n *opNode) solve() {
	if len(n.child) == 0 {
		return
	}
	for i := range n.child {
		n.child[i].solve()
	}
	var vals []int
	for i := range n.child {
		vals = append(vals, n.child[i].val)
	}
	n.val = n.op(vals...)
}

func (n opNode) String() string {
	s := fmt.Sprintf("[%d],", n.val)
	for i := range n.child {
		s += fmt.Sprintf("[%d],", n.child[i].String())
	}
	return s
}

func opTree(ops []packet) *opNode {
	if len(ops) == 0 {
		return nil
	}
	var node opNode
	switch ops[0].opType {
	case 0, 1, 2, 3:
		node.op = opMap[int(ops[0].opType)]
		for i := 1; i < len(ops); i++ {
			if ops[i].opType != 4 {
				break
			}
			node.child = append(node.child, opTree(ops[i:]))
			continue
		}
	case 4:
		node.op = opMap[0]
		node.val = ops[0].literal
		return &node
	case 5, 6, 7:
		node.op = opMap[int(ops[0].opType)]
		node.child = make([]*opNode, 2)
		node.child[0] = opTree(ops[1:])
		switch ops[1].opType {
		case 4:
			node.child[1] = opTree(ops[2:])
		default:
			var opID int
			for i := 2; i < len(ops); i++ {
				if ops[i].opType == 4 {
					continue
				}
				opID = i
				break
			}
			node.child[1] = opTree(ops[opID:])
		}
	default:
		return nil
	}
	return &node
}

type packet struct {
	version int64
	opType  int64
	literal int
}

func parseMsg(r io.Reader) []packet {
	var pkgs []packet
	var done bool
	for !done {
		// log.Print("---")
		v, err := parseBitList(r, 3)
		if err != nil {
			log.Print(err)
			if err == io.EOF {
				done = true
				continue
			}
		}
		log.Print("VERSION is ", v)
		t, err := parseBitList(r, 3)
		if err != nil {
			log.Print(err)
			if err == io.EOF {
				done = true
				continue
			}
		}
		pkg := packet{version: v, opType: t}
		switch t {
		case 4:
			log.Print("type is ", t, ": this is a LITERAL")
			var done bool
			var binN string
			for !done {
				firstBit, err := parseBitList(r, 1)
				log.Print("first bit is ", firstBit, ": if it's 0 it stops, if it's 1 it continues")
				if err != nil {
					log.Print(err)
					if err == io.EOF {
						done = true
						continue
					}
				}
				s, err := parseLiteral(r, 4)
				if err != nil {
					log.Print(err)
					if err == io.EOF {
						done = true
						continue
					}
				}
				binN += s
				if firstBit == 0 {
					done = true
				}
			}

			n, err := strconv.ParseInt(string(binN), 2, 64)
			if err != nil {
				log.Print(err)
				if err == io.EOF {
					done = true
					continue
				}
			}
			log.Print("literal is ", binN)
			pkg.literal = int(n)
		default:
			log.Print("type is ", t, ": this is an OPERATOR")
			modeBit, err := parseBitList(r, 1)
			if err != nil {
				log.Print(err)
				if err == io.EOF {
					done = true
					continue
				}
			}
			switch modeBit {
			case 0:
				log.Print("MODEBIT is ", modeBit, ": next 15 bits will give the max len of subpackets")
				l, err := parseBitList(r, 15)
				if err != nil {
					log.Print(err)
					if err == io.EOF {
						done = true
						continue
					}
				}
				log.Print("max len of subpackets is ", l)
			case 1:
				log.Print("MODEBIT is ", modeBit, ": next 11 bits will give the number of subpackets immediately contained by this packet")
				l, err := parseBitList(r, 11)
				if err != nil {
					log.Print(err)
					if err == io.EOF {
						done = true
						continue
					}
				}
				log.Print("number of subpackets immediately contained by this packet is ", l)
			default:
				log.Fatalf("unexpected modeBit %d", modeBit)
			}
		}
		pkgs = append(pkgs, pkg)
	}
	return pkgs
}

func parseLiteral(r io.Reader, nBits int64) (string, error) {
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

func parseBitList(r io.Reader, nBits int64) (int64, error) {
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
