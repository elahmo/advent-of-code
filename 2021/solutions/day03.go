package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func LinesToInt(lines []string) []int {
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
	lines, _ := fileToLines("input/day03.txt")

	// create a structure
	countMap := map[int]int{
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
	}
	// go through each, append to structure
	for _, line := range lines {
		for idx, char := range line {
			if char == '1' {
				countMap[idx]++
			}
		}
	}
	// fmt.Println(countMap)

	gamma := ""
	epsilon := ""

	for i := 0; i < len(countMap); i++ {
		if countMap[i] > len(lines)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	// fmt.Println(gamma, epsilon)

	// convert to int
	num1, _ := strconv.ParseInt(gamma, 2, 32)
	num2, _ := strconv.ParseInt(epsilon, 2, 32)

	// multiply
	fmt.Println(num1 * num2)
}

func part2() {
	lines, _ := fileToLines("input/day03.txt")

	// create a structure
	countMap := map[int]int{
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
	}
	// go through each, append to structure
	for _, line := range lines {
		for idx, char := range line {
			if char == '1' {
				countMap[idx]++
			}
		}
	}

	gamma := ""

	remainingLines := lines
	for i := 0; i < len(countMap); i++ {
		criteria := ""

		if countMap[i] >= int(math.Ceil(float64(len(remainingLines))/2.0)) {
			criteria = "1"
		} else {
			criteria = "0"
		}
		// fmt.Println("doing index", i, criteria, countMap)
		var validLines []string
		for _, line := range remainingLines {
			lineRune := []rune(line)
			if string(lineRune[i]) == criteria {
				// fmt.Println("valid", line)
				validLines = append(validLines, line)
			}
		}
		remainingLines = validLines
		if len(validLines) == 1 {
			gamma = validLines[0]
			break
		}

		// update countmap
		for i := range countMap {
			countMap[i] = 0
		}
		for _, line := range remainingLines {
			for idx, char := range line {
				if char == '1' {
					countMap[idx]++
				}
			}
		}
	}

	// create a structure
	countMap = map[int]int{
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
	}
	// go through each, append to structure
	for _, line := range lines {
		for idx, char := range line {
			if char == '1' {
				countMap[idx]++
			}
		}
	}

	// fmt.Println("break")

	epsilon := ""

	remainingLines = lines
	for i := 0; i < len(countMap); i++ {
		criteria := ""

		if countMap[i] < int(math.Ceil(float64(len(remainingLines))/2.0)) {
			criteria = "1"
		} else {
			criteria = "0"
		}
		// fmt.Println("doing index", i, criteria, countMap)
		var validLines []string
		for _, line := range remainingLines {
			lineRune := []rune(line)
			if string(lineRune[i]) == criteria {
				// fmt.Println("valid", line)
				validLines = append(validLines, line)
			}
		}
		remainingLines = validLines
		if len(validLines) == 1 {
			epsilon = validLines[0]
			break
		}

		// update countmap
		for i := range countMap {
			countMap[i] = 0
		}
		for _, line := range remainingLines {
			for idx, char := range line {
				if char == '1' {
					countMap[idx]++
				}
			}
		}
	}

	// fmt.Println(gamma)
	// fmt.Println(epsilon)

	// convert to int
	num1, _ := strconv.ParseInt(gamma, 2, 32)
	num2, _ := strconv.ParseInt(epsilon, 2, 32)

	// multiply
	fmt.Println(num1 * num2)
}

func main() {
	start := time.Now()
	part1()
	part2()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
