package day01

import (
	"fmt"

	h "helpers/helpers.go"
)

func part1() {
	lines, _ := h.FileToLines("input/day01.txt")
	numbers := h.LinesToInt(lines)

	increaseCount := 0
	currentDepth := numbers[0]
	for _, num := range numbers {
		if num > currentDepth {
			increaseCount++
		}
		currentDepth = num
	}
	fmt.Println(increaseCount)
}

func part2() {
	lines, _ := fileToLines("input/day01.txt")
	numbers := linesToInt(lines)

	increaseCount := 0
	currentDepth := numbers[0] + numbers[1] + numbers[2]
	for idx, _ := range numbers[1 : len(numbers)-1] {
		sum := numbers[idx] + numbers[idx+1] + numbers[idx+2]
		if sum > currentDepth {
			increaseCount++
		}
		currentDepth = sum
	}
	fmt.Println(increaseCount)
}
