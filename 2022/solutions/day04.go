package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func day04One() int {
	lines, _ := helpers.FileToLines("day04.txt")
	overlaps := 0
	for _, line := range lines {
		re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		oneA, _ := strconv.Atoi(m[0])
		oneB, _ := strconv.Atoi(m[1])
		twoA, _ := strconv.Atoi(m[2])
		twoB, _ := strconv.Atoi(m[3])
		if (oneA <= twoA && oneB >= twoB) || (oneA >= twoA && oneB <= twoB) {
			overlaps++
		}
	}
	return overlaps
}

func day04Two() int {
	lines, _ := helpers.FileToLines("day04.txt")
	overlaps := 0
	for _, line := range lines {
		re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		oneA, _ := strconv.Atoi(m[0])
		oneB, _ := strconv.Atoi(m[1])
		twoA, _ := strconv.Atoi(m[2])
		twoB, _ := strconv.Atoi(m[3])
		if (oneB >= twoA && oneB <= twoB) || (oneA <= twoB && oneA >= twoA) || (oneA <= twoA && oneB >= twoB) || (oneA >= twoA && oneB <= twoB) {
			overlaps++
		}
	}
	return overlaps
}

func Day04() {
	start := time.Now()
	one := day04One()
	two := day04Two()
	elapsed := time.Since(start)
	fmt.Printf("Day04, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
