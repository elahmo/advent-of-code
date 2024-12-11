package solutions

import (
	"aoc24/helpers"
	"fmt"
	"slices"
	"time"
)

func day01One() int {
	lines, _ := helpers.FileToLines("day01.txt")
	var numsLeft, numsRight []int
	for _, line := range lines {
		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d   %d", &num1, &num2)
		if err != nil {
			panic(err)

		}
		numsLeft = append(numsLeft, num1)
		numsRight = append(numsRight, num2)
	}
	slices.Sort(numsLeft)
	slices.Sort(numsRight)
	sum := 0
	for i := range len(numsLeft) {
		sum += helpers.Abs(numsLeft[i] - numsRight[i])
	}
	return sum
}

func day01Two() int {
	lines, _ := helpers.FileToLines("day01.txt")
	var numsLeft []int
	numsRight := make(map[int]int)
	for _, line := range lines {
		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d   %d", &num1, &num2)
		if err != nil {
			panic(err)
		}
		numsLeft = append(numsLeft, num1)
		numsRight[num2]++
	}
	sum := 0
	for _, num := range numsLeft {
		sum += num * numsRight[num]
	}
	return sum
}

func Day01() {
	start := time.Now()
	one := day01One()
	two := day01Two()
	elapsed := time.Since(start)
	fmt.Printf("Day01, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}

func init() {
	RegisterDay("Day01", Day01)
}
