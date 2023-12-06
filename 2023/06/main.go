package main

import "log"

type race struct {
	time           uint
	recordDistance uint
}

var sample = [3]race{
	{time: 7, recordDistance: 9},
	{time: 15, recordDistance: 40},
	{time: 30, recordDistance: 200},
}

var input = [4]race{
	{time: 40, recordDistance: 215},
	{time: 92, recordDistance: 1064},
	{time: 97, recordDistance: 1505},
	{time: 90, recordDistance: 1100},
}

var samplePart2 race = race{time: 71530, recordDistance: 940200}
var inputPart2 race = race{time: 40929790, recordDistance: 215106415051100}

func main() {
	var prod uint = 1
	for i := range input {
		var count uint
		for j := uint(0); j < input[i].time; j++ {
			d := distance(j, input[i].time)
			if d > input[i].recordDistance {
				count++
			}
		}
		prod *= count
	}
	log.Print(prod)

	var count uint
	for j := uint(0); j < inputPart2.time; j++ {
		d := distance(j, inputPart2.time)
		if d > inputPart2.recordDistance {
			count++
		}
	}
	log.Print(count)
}

func distance(holdTime, totalTime uint) uint {
	return holdTime * (totalTime - holdTime)
}
