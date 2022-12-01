package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type node struct {
	bracket  string
	closed   bool
	children []node
	parent   node
}

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

func part1() {
	lines, _ := fileToLines("input/day10.txt")

	score := 0

	for _, line := range lines {
		points := 0
		currentLevel := ""
		for _, letter := range line {
			currentLevel, points = updateLevel(currentLevel, string(letter))
		}
		score += points
	}

	fmt.Println(score)

}

func updateLevel(current, bracket string) (string, int) {
	scores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	updated := current
	points := 0
	if current == "" {
		if validOpeningBracket(bracket) {
			updated += bracket
		} else {
			points += scores[bracket]
		}
		return updated, points
	}

	if validOpeningBracket(bracket) {
		updated += bracket
	} else {
		if !validBracketPair(updated[len(updated)-1:], bracket) {
			points += scores[bracket]
		}
	}

	return updated, points
}

func validOpeningBracket(bracket string) bool {
	if bracket == "[" || bracket == "{" || bracket == "(" || bracket == "<" {
		return true
	}
	return false
}

func validClosingBracket(bracket string) bool {
	if bracket == "]" || bracket == "}" || bracket == ")" || bracket == ">" {
		return true
	}
	return false
}

func validBracketPair(bracketa, bracketb string) bool {
	if (bracketa == "[" && bracketb == "]") || (bracketa == "(" && bracketb == ")") || (bracketa == "<" && bracketb == ">") || (bracketa == "{" && bracketb == "}") {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
