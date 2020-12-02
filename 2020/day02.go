package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	lines, _ := fileToLines("input/day02.txt")
	matches := 0

	for _, line := range lines {
		re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		pwcount := strings.Count(m[3], m[2])
		min, _ := strconv.Atoi(m[0])
		max, _ := strconv.Atoi(m[1])
		if pwcount >= min && pwcount <= max {
			matches++
		}
	}
	fmt.Println(matches)
}

func part2() {
	lines, _ := fileToLines("input/day02.txt")
	matches := 0

	for _, line := range lines {
		re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
		m := re.FindAllStringSubmatch(line, -1)[0][1:]
		pw := m[3]
		char := m[2]
		idxa, _ := strconv.Atoi(m[0])
		idxb, _ := strconv.Atoi(m[1])

		switch {
		case string(pw[idxa-1]) == char && string(pw[idxb-1]) == char:
			break
		case string(pw[idxa-1]) == char || string(pw[idxb-1]) == char:
			matches++
		}
	}
	fmt.Println(matches)
}

func main() {
	part1()
	part2()
}
