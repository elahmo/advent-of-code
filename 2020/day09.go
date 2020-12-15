package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func checkSums(numbers []int, sum int) (addsUp bool) {
	for _, numOne := range numbers {
		for _, numTwo := range numbers {
			if numOne == numTwo {
				continue
			}
			if numOne+numTwo == sum {
				addsUp = true
				break
			}
		}
		if addsUp {
			break
		}
	}
	return addsUp
}

func findFirstBrokenNumber(numbers []int, preamble int) (result int) {
	for i := 0; i < len(numbers)-preamble; i++ {
		pastNumbers := numbers[0+i : preamble+i]
		nextNumber := numbers[i+preamble]
		addsUp := checkSums(pastNumbers, nextNumber)
		if !addsUp {
			result = nextNumber
			break
		}
	}
	return
}

func part1() {
	lines, _ := fileToLines("input/day09.txt")
	numbers := linesToInt(lines)
	preamble := 25
	targetNumber := findFirstBrokenNumber(numbers, preamble)
	fmt.Println(targetNumber)
}

func checkForContiguousNumbers(numbers []int, targetNumber int) (success bool) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
		if sum == targetNumber {
			success = true
			sortedInts := numbers[0 : i+1]
			sort.Ints(sortedInts)
			fmt.Println(sortedInts[0] + sortedInts[len(sortedInts)-1])
			break
		} else if sum > targetNumber {
			break
		}
	}
	return
}

func part2() {
	lines, _ := fileToLines("input/day09.txt")
	numbers := linesToInt(lines)
	preamble := 25
	targetNumber := findFirstBrokenNumber(numbers, preamble)
	for index := range numbers {
		match := checkForContiguousNumbers(numbers[index:], targetNumber)
		if match {
			break
		}
	}
	return
}

func main() {
	part1()
	part2()
}
