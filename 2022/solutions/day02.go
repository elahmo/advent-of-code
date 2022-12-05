package solutions

import (
	"aoc22/helpers"
	"fmt"
	"time"
)

type Move struct {
	score int
	move  string
}

func playGame(opp string, me string) int {
	oppMove := letterToMove(opp)
	meMove := letterToMove(me)
	score := 0

	switch meMove.move {
	case "R":
		switch oppMove.move {
		case "R":
			score += 3
		case "P":
			score += 0
		case "S":
			score += 6
		}
	case "P":
		switch oppMove.move {
		case "R":
			score += 6
		case "P":
			score += 3
		case "S":
			score += 0
		}
	case "S":
		switch oppMove.move {
		case "R":
			score += 0
		case "P":
			score += 6
		case "S":
			score += 3
		}
	}
	score += meMove.score
	return score
}

func playGameForOutcome(opp string, me string) int {
	oppMove := letterToMove(opp)
	result := letterToResult(me)
	score := 0

	switch oppMove.move {
	case "R":
		switch result {
		case 0:
			score += 3
		case 3:
			score += 1
		case 6:
			score += 2
		}
	case "P":
		switch result {
		case 0:
			score += 1
		case 3:
			score += 2
		case 6:
			score += 3
		}
	case "S":
		switch result {
		case 0:
			score += 2
		case 3:
			score += 3
		case 6:
			score += 1
		}
	}
	score += result
	return score
}

func letterToMove(letter string) Move {
	scores := map[string]Move{
		"A": {1, "R"},
		"B": {2, "P"},
		"C": {3, "S"},
		"X": {1, "R"},
		"Y": {2, "P"},
		"Z": {3, "S"},
	}
	return scores[letter]
}

func letterToResult(letter string) int {
	result := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	return result[letter]
}

func day02One() int {
	lines, _ := helpers.FileToLines("day02.txt")
	score := 0
	for _, line := range lines {
		score += playGame(string(line[0]), string(line[2]))
	}
	return score
}

func day02Two() int {
	lines, _ := helpers.FileToLines("day02.txt")
	score := 0
	for _, line := range lines {
		score += playGameForOutcome(string(line[0]), string(line[2]))
	}
	return score
}

func Day02() {
	start := time.Now()
	one := day02One()
	two := day02Two()
	elapsed := time.Since(start)
	fmt.Printf("Day02, part 1: %d, part 2: %d (%s)\n", one, two, elapsed)
}
