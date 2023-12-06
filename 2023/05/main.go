package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type category int

const (
	seed_to_soil = category(iota)
	soil_to_fertilizer
	fertilizer_to_water
	water_to_light
	light_to_temperature
	temperature_to_humidity
	humidity_to_location
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var infoID category
	var seedsList seeds
	var entries [7][]almalancMapEntry

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "seed"):
			sep := strings.Index(line, ":") + 1
			for _, n := range strings.Fields(line[sep:]) {
				ns, err := strconv.Atoi(n)
				if err != nil {
					continue
				}
				seedsList = append(seedsList, ns)
			}
		case strings.Contains(line, "seed-to-soil"):
			infoID = seed_to_soil
		case strings.Contains(line, "soil-to-fertilizer"):
			infoID = soil_to_fertilizer
		case strings.Contains(line, "fertilizer-to-water"):
			infoID = fertilizer_to_water
		case strings.Contains(line, "water-to-light"):
			infoID = water_to_light
		case strings.Contains(line, "light-to-temperature"):
			infoID = light_to_temperature
		case strings.Contains(line, "temperature-to-humidity"):
			infoID = temperature_to_humidity
		case strings.Contains(line, "humidity-to-location"):
			infoID = humidity_to_location
		case len(line) == 0:
		default:
			f := strings.Fields(line)
			d, _ := strconv.Atoi(f[0])
			s, _ := strconv.Atoi(f[1])
			r, _ := strconv.Atoi(f[2])
			entries[infoID] = append(entries[infoID], almalancMapEntry{d: d, s: s, r: r})
		}
	}

	// part 2
	// I thought to go backwards fron the almalanc mappings to see which values were available starting from the minimun range
	// To implement later if time permits
	// But the part 1 algorithm works for getting the star so I will subdivide the bigger set into subsets and work on them
	// Could be done in parallel again to try if time permits
	var i int = 16
	seedsList = update(seedsList[i : i+4])

	var seedLocations []int
	for i := range seedsList {
		locationID := seedsList[i]
	nextEntry:
		for j := range entries {
			for k := range entries[j] {
				var ok bool
				locationID, ok = toDest(entries[j][k], locationID)
				if !ok {
					continue
				}
				continue nextEntry
			}
		}
		seedLocations = append(seedLocations, locationID)
	}
	sort.Ints(seedLocations)
	log.Println(seedLocations[0])
}

type seeds []int

type almalancMapEntry struct{ d, s, r int }

func toDest(e almalancMapEntry, s int) (int, bool) {
	if s >= e.s && s < e.s+e.r {
		return e.d + (s - e.s), true
	}
	return s, false
}

func update(s seeds) seeds {
	var ns seeds
	for i := 0; i < len(s); i += 2 {
		for j := 0; j < s[i+1]; j++ {
			ns = append(ns, s[i]+j)
		}
	}
	return ns
}
