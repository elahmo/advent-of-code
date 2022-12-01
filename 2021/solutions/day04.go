package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

type Board struct {
	numbers [25]int
	visited [25]bool
	won     bool
}

func contains(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

func part1() {
	lines, _ := fileToLines("input/day04.txt")

	// get the boards
	var boards []Board
	var numbers [25]int
	var visited [25]bool
	boardRegEx := regexp.MustCompile(`\d+`)
	boardIdx := 0
	for _, line := range lines[2:] {
		if line == "" {
			boards = append(boards, Board{numbers: numbers, visited: visited, won: false})
			for i := range visited {
				visited[i] = false
			}
			boardIdx = 0
		} else {
			nums := boardRegEx.FindAllStringSubmatch(line, -1)
			for idx, num := range nums {
				bingo, _ := strconv.Atoi(num[0])
				curIdx := idx + boardIdx*5
				numbers[curIdx] = bingo
			}
			boardIdx++
		}
	}
	// fmt.Println("Boards", boards)

	// get the numbers
	numRegEx := regexp.MustCompile(`\d+`)
	nums := numRegEx.FindAllStringSubmatch(lines[0], -1)
	won := false
	var drawnNumbers []int
	for _, num := range nums {
		intNum, _ := strconv.Atoi(num[0])
		drawnNumbers = append(drawnNumbers, intNum)

		if len(drawnNumbers) >= 5 {
			boards, won = updateAndCheckAllBoards(boards, drawnNumbers)
			if won {
				return
			}
		}
	}

}

func updateAndCheckAllBoards(boards []Board, drawnNumbers []int) ([]Board, bool) {
	// update all boards
	var updatedBoards []Board
	for _, board := range boards {
		for i := 0; i < 5; i++ {
			for y := 0; y < 5; y++ {
				if contains(drawnNumbers, board.numbers[y+i*5]) {
					board.visited[y+i*5] = true
				}
			}
		}
		updatedBoards = append(updatedBoards, board)
	}

	// check for winner
	winningBoards := countWinningBoards(updatedBoards)
	won := false
	for idx, board := range updatedBoards {
		if board.won {
			continue
		}
		for i := 0; i < 5; i++ {
			if board.visited[5*i] && board.visited[5*i+1] && board.visited[5*i+2] && board.visited[5*i+3] && board.visited[5*i+4] {
				if winningBoards == len(updatedBoards)-1 {
					score := calculateScore(board)
					// fmt.Println("WINNER, row", i, board, score, drawnNumbers[len(drawnNumbers)-1])
					fmt.Println(score * drawnNumbers[len(drawnNumbers)-1])
					won = true
					break
				} else {
					updatedBoards[idx].won = true
					break
				}
			}
			if board.visited[i] && board.visited[i+5] && board.visited[i+10] && board.visited[i+15] && board.visited[i+20] {
				if winningBoards == len(updatedBoards)-1 {
					score := calculateScore(board)
					// fmt.Println("WINNER, col", i, board, score, drawnNumbers[len(drawnNumbers)-1])
					fmt.Println(score * drawnNumbers[len(drawnNumbers)-1])
					won = true
					break
				} else {
					updatedBoards[idx].won = true
					break
				}
			}
		}
	}

	// fmt.Println("Drawn number: ", drawnNumbers[len(drawnNumbers)-1])
	// fmt.Println("Winners: ", countWinningBoards(updatedBoards))

	return updatedBoards, won
}

func countWinningBoards(boards []Board) (count int) {
	for _, board := range boards {
		if board.won {
			count++
		}
	}
	return
}

func calculateScore(board Board) (score int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.visited[i+j*5] {
				score += board.numbers[i+j*5]
			}
		}
	}
	return
}

func main() {
	start := time.Now()
	part1()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
