package solutions

import (
	"aoc23/helpers"
	"fmt"
	"math"
	"regexp"
	"time"
)

func day04One() int {
	lines, _ := helpers.FileToLines("day04.txt")
	rex := regexp.MustCompile(`(\d+)`)
	sum := 0
	for _, line := range lines {
		matches := rex.FindAllString(line, -1)
		winning := make(map[string]bool)
		for _, match := range matches[1:11] {
			winning[match] = true
		}
		var points int
		for _, match := range matches[11:] {
			if ok := winning[match]; ok {
				points++
			}
		}
		if points > 0 {
			sum += int(math.Pow(2, float64(points-1)))
		}
	}
	return sum
}

func day04Two() int {
	lines, _ := helpers.FileToLines("day04.txt")
	rex := regexp.MustCompile(`(\d+)`)
	sum := 0
	extraCardCounts := make([]int, len(lines))
	for linesIdx, line := range lines {
		matches := rex.FindAllString(line, -1)
		winning := make(map[string]bool)
		for _, match := range matches[1:11] {
			winning[match] = true
		}
		var points int
		for _, match := range matches[11:] {
			if ok := winning[match]; ok {
				points++
			}
		}
		// increase card counts for numbers of points below
		currentCardCount := 1 + extraCardCounts[linesIdx]
		for i := 0; i < points; i++ {
			extraCardCounts[linesIdx+i+1] += currentCardCount
		}
		sum += extraCardCounts[linesIdx] + 1
	}
	return sum
}

func Day04() {
	start := time.Now()
	one := day04One()
	two := day04Two()
	elapsed := time.Since(start)
	fmt.Printf("Day04, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
