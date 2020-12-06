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
	ints := make([]int, 0, len(lines))
	for _, w := range lines {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func linesToGroups(lines []string) []int {
	var groups = make([]int, 500)
	currentGroup := 0
	commonAnswers := make([]string, 0, 26)
	firstWord := true
	for _, w := range lines {
		if w != "" {
			if firstWord {
				commonAnswers = strings.Split(w, "")
				firstWord = false
			} else {
				commonAnswers = Intersection(commonAnswers, strings.Split(w, ""))
			}
		} else {
			groups[currentGroup] = len(commonAnswers)
			commonAnswers = make([]string, 0, 26)
			firstWord = true
			currentGroup++
		}
	}
	return groups
}

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func part2() {
	lines, _ := fileToLines("input/day06.txt")
	groupData := linesToGroups(lines)

	total := 0
	for _, group := range groupData {
		total += group
	}
	fmt.Println(total)
	return
}

func main() {
	part2()
}
