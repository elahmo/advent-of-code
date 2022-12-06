package solutions

import (
	"aoc22/helpers"
	"fmt"
	"time"
)

func day06One() int {
	lines, _ := helpers.FileToLines("day06.txt")
	line := lines[0]
	for i := 0; i < len(line)-4; i++ {
		currentMarker := []string{string(line[i]), string(line[i+1]), string(line[i+2]), string(line[i+3])}
		if len(helpers.Unique(currentMarker)) == 4 {
			return i + 4
		}
	}
	return 0
}

func day06Two() int {
	lines, _ := helpers.FileToLines("day06.txt")
	line := lines[0]
	for i := 0; i < len(line)-14; i++ {
		if helpers.UniqueChars(line[i:i+14]) == 14 {
			return i + 14
		}
	}
	return 0
}

func Day06() {
	start := time.Now()
	one := day06One()
	two := day06Two()
	elapsed := time.Since(start)
	fmt.Printf("Day06, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
