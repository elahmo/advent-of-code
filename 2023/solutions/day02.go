package solutions

import (
	"aoc23/helpers"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func day02One() int {
	lines, _ := helpers.FileToLines("day02.txt")
	maxCounts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0
	rexGame := regexp.MustCompile(`^Game (\d+)`)
	rexBalls := regexp.MustCompile(`(\d+) (\w+)`)

	for _, line := range lines {
		validLine := true
		game, _ := strconv.Atoi(rexGame.FindStringSubmatch(line)[1])
		matches := rexBalls.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			colour := match[2]
			count, _ := strconv.Atoi(match[1])
			if count > maxCounts[colour] {
				validLine = false
			}
		}
		if validLine {
			sum += game
		}
	}
	return sum
}

func day02Two() int {
	lines, _ := helpers.FileToLines("day02.txt")
	sum := 0
	rexBalls := regexp.MustCompile(`(\d+) (\w+)`)
	for _, line := range lines {
		colourCounts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		matches := rexBalls.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			colour := match[2]
			count, _ := strconv.Atoi(match[1])
			if colourCounts[colour] < count {
				colourCounts[colour] = count
			}
		}
		sum += colourCounts["red"] * colourCounts["green"] * colourCounts["blue"]
	}
	return sum
}

func Day02() {
	start := time.Now()
	one := day02One()
	two := day02Two()
	elapsed := time.Since(start)
	fmt.Printf("Day02, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
