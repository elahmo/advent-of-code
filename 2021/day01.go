package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func fileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func linesToInt(lines []string) []int {
	ints := make([]int, 0, len(lines))
	for _, w := range lines {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func part1() {
	lines, _ := fileToLines("input/day01.txt")
	numbers := linesToInt(lines)

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

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	fmt.Println("Part 1 took ", elapsed)
	start = time.Now()
	part2()
	elapsed = time.Since(start)
	fmt.Println("Part 2 took ", elapsed)
}
