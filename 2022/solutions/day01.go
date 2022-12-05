package solutions

import (
	"aoc22/helpers"
	"fmt"
	"sort"
	"strconv"
)

func day01One() int {
	lines, _ := helpers.FileToLines("day01.txt")
	maxCalories := 0
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err == nil {
			currentCalories += calories
		}
	}
	return maxCalories
}

func day01Two() int {
	lines, _ := helpers.FileToLines("day01.txt")
	var calories []int
	currentCalories := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(line)
		if err == nil {
			currentCalories += calories
		}
	}
	sort.Ints(calories)
	return calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
}

func Day01() {
	fmt.Printf("Day01, part 1: %d, part 2: %d\n", day01One(), day01Two())
}
