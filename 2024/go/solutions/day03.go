package solutions

import (
	"aoc24/helpers"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func Day03One() int {
	lines, _ := helpers.FileToLines("day03.txt")
	sum := 0
	for _, line := range lines {
		lineRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := lineRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += a * b
		}
	}
	return sum
}

func Day03Two() int {
	lines, _ := helpers.FileToLines("day03.txt")
	sum := 0
	eligible := true
	for _, line := range lines {
		lineRegex := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(don\'t)\(\)|(do)\(\)`)
		matches := lineRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "don't()":
				eligible = false
			case "do()":
				eligible = true
			default:
				if eligible {
					a, _ := strconv.Atoi(match[2])
					b, _ := strconv.Atoi(match[3])
					sum += a * b
				}
			}
		}
	}
	return sum
}

func Day03() {
	start := time.Now()
	one := Day03One()
	two := Day03Two()
	elapsed := time.Since(start)
	fmt.Printf("Day03, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day03", Day03)
}
