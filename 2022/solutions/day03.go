package solutions

import (
	"aoc22/helpers"
	"fmt"
	"strings"
)

func charToPoint(c string) int {
	runes := []rune(c)
	asciiValue := int(runes[0])
	// uppercase
	if asciiValue < 96 {
		return asciiValue - 38
	}
	// lowercase
	return asciiValue - 96

}

func day03One() int {
	lines, _ := helpers.FileToLines("day03.txt")
	var common []string
	for _, line := range lines {
		firstHalf := line[0 : len(line)/2]
		secondHalf := line[len(line)/2:]

		common = append(common, helpers.Unique(helpers.IntersectionStrings(firstHalf, secondHalf))...)
	}
	score := 0
	for _, char := range common {
		score += charToPoint(char)
	}
	return score
}

func day03Two() int {
	lines, _ := helpers.FileToLines("day03.txt")
	var common []string

	for i := 0; i < len(lines)/3; i++ {
		lineOne := lines[i*3]
		lineTwo := lines[i*3+1]
		lineThree := lines[i*3+2]

		oneAndTwo := strings.Join(helpers.IntersectionStrings(lineOne, lineTwo), "")
		common = append(common, helpers.Unique(helpers.IntersectionStrings(oneAndTwo, lineThree))...)
	}
	score := 0
	for _, char := range common {
		score += charToPoint(char)
	}
	return score
}

func Day03() {
	fmt.Printf("Day03, part 1: %d, part 2: %d", day03One(), day03Two())
}
