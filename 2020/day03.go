package main

import (
	"bufio"
	"fmt"
	"os"
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

func parts1and2() {
	lines, _ := fileToLines("input/day03.txt")
	lineLength := len(lines[0])
	movements := [][]int{
		{1, 1},
		{3, 1}, //part 1
		{5, 1},
		{7, 1},
		{1, 2},
	}

	totalProduct := 1

	for _, movement := range movements {
		matches := 0
		currentx, currenty := 0, 0
		for i := 0; i < len(lines); i += movement[1] {
			line := lines[i]
			if currentx > lineLength-1 {
				currentx -= lineLength
			}
			if string(line[currentx]) == "#" {
				matches++
			}
			currentx += movement[0]
			currenty++
		}
		fmt.Println(matches)
		totalProduct *= matches
	}
	fmt.Println(totalProduct)
}

func main() {
	parts1and2()
}
