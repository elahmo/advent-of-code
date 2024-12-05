package solutions

import (
	"aoc24/helpers"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func day02One() int {
	lines, _ := helpers.FileToLines("day02.txt")
	var reports [][]int
	for _, line := range lines {
		lineRegex := regexp.MustCompile(`\d+`)
		matches := lineRegex.FindAllString(line, -1)
		var intMatches []int
		for _, match := range matches {
			intMatch, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			intMatches = append(intMatches, intMatch)
		}
		reports = append(reports, intMatches)
	}
	sum := 0
	for _, report := range reports {
		if isSafe(report) {
			sum += 1
		}
	}
	return sum
}

func isSafe(report []int) bool {
	valid := true
	var increasing bool
	if report[1] > report[0] {
		increasing = true
	} else if report[1] < report[0] {
		increasing = false
	} else {
		return false
	}
	for i := 1; i < len(report); i++ {
		if increasing {
			if report[i]-report[i-1] < 1 || report[i]-report[i-1] > 3 {
				return false
			}
			continue
		}
		if report[i-1]-report[i] < 1 || report[i-1]-report[i] > 3 {
			return false
		}
	}
	return valid
}

func day02Two() int {
	lines, _ := helpers.FileToLines("day02.txt")
	var reports [][]int
	for _, line := range lines {
		lineRegex := regexp.MustCompile(`\d+`)
		matches := lineRegex.FindAllString(line, -1)
		var intMatches []int
		for _, match := range matches {
			intMatch, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			intMatches = append(intMatches, intMatch)
		}
		reports = append(reports, intMatches)
	}
	sum := 0
	for _, report := range reports {
		if isSafeRelax(report) {
			sum += 1
		}
	}
	return sum
}

func isSafeRelax(report []int) bool {
	// Helper function to check if a slice satisfies the condition
	checkValid := func(slice []int) bool {
		increasing := slice[1] > slice[0]
		for i := 1; i < len(slice); i++ {
			if increasing {
				if slice[i]-slice[i-1] < 1 || slice[i]-slice[i-1] > 3 {
					return false
				}
			} else {
				if slice[i-1]-slice[i] < 1 || slice[i-1]-slice[i] > 3 {
					return false
				}
			}
		}
		return true
	}

	// Check the full slice first
	if len(report) > 1 && checkValid(report) {
		return true
	}

	// Iterate over all slices by removing one element at a time
	n := len(report)
	for i := 0; i < n; i++ {
		// Create a slice with the i-th element removed
		temp := make([]int, 0, n-1)
		temp = append(temp, report[:i]...)
		temp = append(temp, report[i+1:]...)

		// Check if the slice satisfies the condition
		if len(temp) > 1 && checkValid(temp) {
			return true
		}
	}

	// If neither the full slice nor any modified slice satisfies the condition, return false
	return false
}

func Day02() {
	start := time.Now()
	one := day02One()
	two := day02Two()
	elapsed := time.Since(start)
	fmt.Printf("Day02, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
