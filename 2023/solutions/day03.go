package solutions

import (
	"aoc23/helpers"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func day03One() int {
	lines, _ := helpers.FileToLines("day03.txt")
	// collect all numbers and special characters indexes
	rexNum := regexp.MustCompile(`(\d+)`)
	rexSpecialChars := regexp.MustCompile(`([!@#$%^&*()_\-+={}\[\]:;<>,?~\\/\|\'"])`)
	var numbers [][][]int
	var specialChars [][][]int
	for _, line := range lines {
		matches := rexNum.FindAllStringIndex(line, -1)
		numbers = append(numbers, matches)
		matchesSpecial := rexSpecialChars.FindAllStringIndex(line, -1)
		specialChars = append(specialChars, matchesSpecial)
	}
	// check special characters next to numbers
	sum := 0
	for i := 0; i < len(numbers); i++ {
		currentLineNumbers := numbers[i]
		for j := 0; j < len(currentLineNumbers); j++ {
			touching := false
			currentNum := currentLineNumbers[j]
			numStart, numEnd := currentNum[0], currentNum[1]

			// check current row
			specialCharRow := specialChars[i]
			for k := 0; k < len(specialCharRow); k++ {
				currentSpecial := specialCharRow[k][0]
				if currentSpecial == numStart-1 || currentSpecial == numEnd {
					touching = true
					break
				}
			}

			// check row above
			if i > 0 && !touching {
				specialCharRow := specialChars[i-1]
				for k := 0; k < len(specialCharRow); k++ {
					currentSpecial := specialCharRow[k][0]
					if currentSpecial >= numStart-1 && currentSpecial <= numEnd {
						touching = true
						break
					}
				}
			}

			// check row below
			if i < len(numbers)-1 && !touching {
				specialCharRow := specialChars[i+1]
				for k := 0; k < len(specialCharRow); k++ {
					currentSpecial := specialCharRow[k][0]
					if currentSpecial >= numStart-1 && currentSpecial <= numEnd {
						touching = true
						break
					}
				}
			}

			if touching {
				num, _ := strconv.Atoi(lines[i][numStart:numEnd])
				sum += num
			}
		}

	}
	return sum
}

func day03Two() int {
	lines, _ := helpers.FileToLines("day03.txt")
	// collect all numbers and special characters indexes
	rexNum := regexp.MustCompile(`(\d+)`)
	rexSpecialChars := regexp.MustCompile(`(\*)`)
	var numbers [][][]int
	var specialChars [][][]int
	for _, line := range lines {
		matches := rexNum.FindAllStringIndex(line, -1)
		numbers = append(numbers, matches)
		matchesSpecial := rexSpecialChars.FindAllStringIndex(line, -1)
		specialChars = append(specialChars, matchesSpecial)
	}
	// check special characters next to numbers
	sum := 0
	for i := 0; i < len(specialChars); i++ {
		currentLineChars := specialChars[i]
		for j := 0; j < len(currentLineChars); j++ {
			var nums []int
			currentSpecial := currentLineChars[j][0]

			// check current row
			numRow := numbers[i]
			for k := 0; k < len(numRow); k++ {
				curNum := numRow[k]
				numStart, numEnd := curNum[0], curNum[1]
				if currentSpecial == numStart-1 || currentSpecial == numEnd {
					num, _ := strconv.Atoi(lines[i][numStart:numEnd])
					nums = append(nums, num)
				}
			}

			// check row above
			if i > 0 {
				numRow := numbers[i-1]
				for k := 0; k < len(numRow); k++ {
					curNum := numRow[k]
					numStart, numEnd := curNum[0], curNum[1]
					if currentSpecial >= numStart-1 && currentSpecial <= numEnd {
						num, _ := strconv.Atoi(lines[i-1][numStart:numEnd])
						nums = append(nums, num)
					}
				}
			}

			// check row below
			if i < len(numbers)-1 {
				numRow := numbers[i+1]
				for k := 0; k < len(numRow); k++ {
					curNum := numRow[k]
					numStart, numEnd := curNum[0], curNum[1]
					if currentSpecial >= numStart-1 && currentSpecial <= numEnd {
						num, _ := strconv.Atoi(lines[i+1][numStart:numEnd])
						nums = append(nums, num)
					}
				}
			}

			if len(nums) == 2 {
				sum += nums[0] * nums[1]
			}
		}

	}
	return sum
}

func Day03() {
	start := time.Now()
	one := day03One()
	two := day03Two()
	elapsed := time.Since(start)
	fmt.Printf("Day03, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
