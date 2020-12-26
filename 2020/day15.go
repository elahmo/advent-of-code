package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	ints := make([]int, 0)
	for _, w := range lines {
		nums := strings.Split(w, ",")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err == nil {
				ints = append(ints, i)
			}
		}
	}
	return ints
}

func part1and2() {
	lines, _ := fileToLines("input/day15.txt")
	numbers := linesToInt(lines)
	seen := map[int]bool{}
	for _, num := range numbers {
		seen[num] = true
	}
	lastNum := numbers[len(numbers)-1]
	seenLast := false
	currentNumber := 0
	for counter := len(numbers) + 1; counter <= 30000000; counter++ {
		if !seenLast {
			currentNumber = 0
		} else {
			// find difference
			var last, prior int
			for i := len(numbers) - 1; i >= 0; i-- {
				if numbers[i] == lastNum {
					last = i
					break
				}
			}
			for i := last - 1; i >= 0; i-- {
				if numbers[i] == lastNum {
					prior = i
					break
				}
			}
			currentNumber = last - prior
		}
		numbers = append(numbers, currentNumber)
		seenLast = seen[currentNumber]
		seen[currentNumber] = true
		lastNum = currentNumber
	}
	fmt.Println(lastNum)
}

func main() {
	part1and2() //2 is very slow... VERY, needs to be optimised
}
