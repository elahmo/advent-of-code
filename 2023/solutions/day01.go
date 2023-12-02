package solutions

import (
	"aoc23/helpers"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day01One() int {
	lines, _ := helpers.FileToLines("day01.txt")
	sum := 0
	for _, line := range lines {
		var firstDigit, lastDigit string
		for _, char := range line {
			charString := string(char)
			_, err := strconv.Atoi(charString)
			if err == nil {
				firstDigit = charString
				break
			}
		}
		for i := len(line) - 1; i > -1; i-- {
			charString := string(line[i])
			_, err := strconv.Atoi(charString)
			if err == nil {
				lastDigit = charString
				break
			}
		}
		num, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += num
	}
	return sum
}

func day01Two() int {
	lines, _ := helpers.FileToLines("day01.txt")
	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	sum := 0
	for _, line := range lines {
		firstDigit := 9999999
		lastDigit := -1
		var firstNum, lastNum string
		for i, char := range line {
			charString := string(char)
			_, err := strconv.Atoi(charString)
			if err == nil {
				firstDigit = i
				firstNum = charString
				break
			}
		}
		for i := len(line) - 1; i > -1; i-- {
			charString := string(line[i])
			_, err := strconv.Atoi(charString)
			if err == nil {
				lastDigit = i
				lastNum = charString
				break
			}
		}

		// try to find words in line
		for word, number := range numbersMap {
			wordIdx := strings.Index(line, word)
			if wordIdx > -1 {
				if wordIdx < firstDigit {
					firstDigit = wordIdx
					firstNum = number
				}
			}
			lastWordIdx := strings.LastIndex(line, word)
			if lastWordIdx > -1 {
				if lastWordIdx > lastDigit {
					lastDigit = lastWordIdx
					lastNum = number
				}
			}
		}
		num, _ := strconv.Atoi(firstNum + lastNum)
		sum += num
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
