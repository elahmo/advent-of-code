package solutions

import (
	"aoc22/helpers"
	"fmt"
	"regexp"
	"time"
)

func day05One() string {
	lines, _ := helpers.FileToLines("day05.txt")
	var stacks [9][]string
	movesIdx := 0
	for idx, line := range lines {
		for i := 1; i < len(line); i += 4 {
			char := string(line[i])
			if char != " " {
				stacks[(i-1)/4] = append(stacks[(i-1)/4], char)
			}
		}
		if string(line[1]) == "1" {
			movesIdx = idx + 2
			break
		}
	}
	result := ""
	for _, line := range lines[movesIdx:] {
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		qty, source, target := helpers.StrToInt(m[0]), helpers.StrToInt(m[1])-1, helpers.StrToInt(m[2])-1
		stacks[target] = helpers.PrependElements(stacks[target], stacks[source][:qty])
		stacks[source] = helpers.RemoveElements(stacks[source], qty)
	}
	for _, stack := range stacks {
		result += stack[0]
	}
	return result
}

func day05Two() string {
	lines, _ := helpers.FileToLines("day05.txt")
	var stacks [9][]string
	movesIdx := 0
	for idx, line := range lines {
		for i := 1; i < len(line); i += 4 {
			char := string(line[i])
			if char != " " {
				stacks[(i-1)/4] = append(stacks[(i-1)/4], char)
			}
		}
		if string(line[1]) == "1" {
			movesIdx = idx + 2
			break
		}
	}
	result := ""
	for _, line := range lines[movesIdx:] {
		re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		qty, source, target := helpers.StrToInt(m[0]), helpers.StrToInt(m[1])-1, helpers.StrToInt(m[2])-1
		stacks[target] = helpers.PrependElementsMultiple(stacks[target], stacks[source][:qty])
		stacks[source] = helpers.RemoveElements(stacks[source], qty)
	}
	for _, stack := range stacks {
		result += stack[0]
	}
	return result
}

func Day05() {
	start := time.Now()
	one := day05One()
	two := day05Two()
	elapsed := time.Since(start)
	fmt.Printf("Day05, part 1: %s, part 2: %s (%s)\n", one, two, elapsed)
}
